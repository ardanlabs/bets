// Package gamegrp provides the handlers for game play.
package gamegrp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ardanlabs/bets/business/core/bank"
	"github.com/ardanlabs/bets/business/core/bet"
	"github.com/ardanlabs/bets/business/web/auth"
	v1Web "github.com/ardanlabs/bets/business/web/v1"
	"github.com/ardanlabs/bets/foundation/events"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Bet            bet.Core
	Converter      *currency.Converter
	Bank           *bank.Bank
	Log            *zap.SugaredLogger
	WS             websocket.Upgrader
	Evts           *events.Events
	Auth           *auth.Auth
	BankTimeout    time.Duration
	ConnectTimeout time.Duration

	mu sync.RWMutex
}

// Connect is used to return a game token for API usage.
func (h *Handlers) Connect(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	address, err := validateSignature(r, h.ConnectTimeout)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	token, err := generateToken(h.Auth, address)
	if err != nil {
		return v1Web.NewRequestError(err, http.StatusBadRequest)
	}

	data := struct {
		Token   string `json:"token"`
		Address string `json:"address"`
	}{
		Token:   token,
		Address: address,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// Events handles a web socket to provide events to a client.
func (h *Handlers) Events(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return v1Web.NewRequestError(errors.New("web value missing from context"), http.StatusBadRequest)
	}

	// Need this to handle CORS on the websocket.
	h.WS.CheckOrigin = func(r *http.Request) bool { return true }

	// This upgrades the HTTP connection to a websocket connection.
	c, err := h.WS.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	h.Log.Infow("websocket open", "path", "/v1/game/events", "traceid", v.TraceID)

	// Set the timeouts for the ping to identify if a web socket
	// connection is broken.
	pongWait := 15 * time.Second
	pingPeriod := (pongWait * 9) / 10

	c.SetReadDeadline(time.Now().Add(pongWait))

	// Setup the pong handler to log the receiving of a pong.
	f := func(appData string) error {
		c.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	}
	c.SetPongHandler(f)

	// This provides a channel for receiving events from the blockchain.
	ch := h.Evts.Acquire(v.TraceID)
	defer h.Evts.Release(v.TraceID)

	// Starting a ticker to send a ping message over the websocket.
	pingSend := time.NewTicker(pingPeriod)

	// Set up the ability to receive chat messages.
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// This supports the ability to add a chat system and receive a client
		// message.
		for {
			message, p, err := c.ReadMessage()
			if err != nil {
				return
			}
			h.Log.Infow("*********> socket read", "path", "/v1/game/events", "message", message, "p", string(p))
		}
	}()

	defer func() {
		wg.Wait()
		h.Log.Infow("websocket closed", "path", "/v1/game/events", "traceid", v.TraceID)
	}()
	defer c.Close()

	// Send game engine events back to the connected client.
	for {
		select {
		case msg, wd := <-ch:

			// If the channel is closed, release the websocket.
			if !wd {
				return nil
			}

			if err := c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				return err
			}

		case <-pingSend.C:
			if err := c.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
				return nil
			}
		}
	}
}

// Configuration returns the basic configuration the front end needs to use.
func (h *Handlers) Configuration(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	info := struct {
		Network    string `json:"network"`
		ChainID    int    `json:"chainId"`
		ContractID string `json:"contractId"`
	}{
		Network:    h.Bank.Client().Network(),
		ChainID:    h.Bank.Client().ChainID(),
		ContractID: h.Bank.ContractID(),
	}

	return web.Respond(ctx, w, info, http.StatusOK)
}

// USD2Wei converts the us dollar amount to wei based on the game engine's
// conversion rate.
func (h *Handlers) USD2Wei(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	usd, err := strconv.ParseFloat(web.Param(r, "usd"), 64)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("converting usd: %s", err), http.StatusBadRequest)
	}

	wei := h.Converter.USD2Wei(big.NewFloat(usd))

	data := struct {
		USD float64  `json:"usd"`
		WEI *big.Int `json:"wei"`
	}{
		USD: usd,
		WEI: wei,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// Test will validate things are working.
func (h *Handlers) Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	claims, err := auth.GetClaims(ctx)
	if err != nil {
		return v1Web.NewRequestError(auth.ErrForbidden, http.StatusForbidden)
	}
	address := claims.Subject

	resp := Status{
		Status:  "OK",
		Address: address,
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

// =========================================================================
// Bet Support

// CreateBet creates a new bet and returns its content.
func (h *Handlers) CreateBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	var newBetPayload NewBet
	if err := web.Decode(r, &newBetPayload); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	newBet := bet.NewBet{
		Description:      newBetPayload.Description,
		Terms:            newBetPayload.Terms,
		Amount:           newBetPayload.Amount,
		ModeratorAddress: newBetPayload.ModeratorAddress,
		DateExpired:      time.Unix(int64(newBetPayload.DateExpired), 0),
	}

	var newPlayers []bet.NewPlayer
	for _, player := range newBetPayload.Players {
		newPlayers = append(newPlayers, bet.NewPlayer(player))
	}
	newBet.Players = newPlayers

	bet, err := h.Bet.CreateBet(ctx, newBet, v.Now)
	if err != nil {
		return fmt.Errorf("creating new bet, newBet[%+v]: %w", newBet, err)
	}

	return web.Respond(ctx, w, bet, http.StatusCreated)
}

// UpdateBet updates a bet in the system.
func (h *Handlers) UpdateBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	var upBetPayload UpdateBet
	if err := web.Decode(r, &upBetPayload); err != nil {
		return fmt.Errorf("unable to decode payload: %w", err)
	}

	id := web.Param(r, "id")

	upBet := bet.UpdateBet{
		Description:      upBetPayload.Description,
		Terms:            upBetPayload.Terms,
		Amount:           upBetPayload.Amount,
		ModeratorAddress: upBetPayload.ModeratorAddress,
	}

	if upBetPayload.DateExpired != nil {
		upBet.DateExpired = time.Unix(int64(*upBetPayload.DateExpired), 0)
	}

	if err := h.Bet.UpdateBet(ctx, id, upBet, v.Now); err != nil {
		switch {
		case errors.Is(err, bet.ErrInvalidID):
			return v1Web.NewRequestError(err, http.StatusBadRequest)
		case errors.Is(err, bet.ErrNotFound):
			return v1Web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s] Bet[%+v]: %w", id, &upBet, err)
		}
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// QueryBet returns a list of bets with paging.
func (h *Handlers) QueryBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	page := web.Param(r, "page")
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("invalid page format, page[%s]", page), http.StatusBadRequest)
	}
	rows := web.Param(r, "rows")
	rowsPerPage, err := strconv.Atoi(rows)
	if err != nil {
		return v1Web.NewRequestError(fmt.Errorf("invalid rows format, rows[%s]", rows), http.StatusBadRequest)
	}

	bets, err := h.Bet.QueryBet(ctx, pageNumber, rowsPerPage)
	if err != nil {
		return fmt.Errorf("unable to query for bets: %w", err)
	}

	return web.Respond(ctx, w, bets, http.StatusOK)
}

// QueryBetByID returns a bet by its ID.
func (h *Handlers) QueryBetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := web.Param(r, "id")
	bt, err := h.Bet.QueryBetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, bet.ErrInvalidID):
			return v1Web.NewRequestError(err, http.StatusBadRequest)
		case errors.Is(err, bet.ErrNotFound):
			return v1Web.NewRequestError(err, http.StatusNotFound)
		default:
			return fmt.Errorf("ID[%s]: %w", id, err)
		}
	}

	return web.Respond(ctx, w, bt, http.StatusOK)
}

// =============================================================================

func validateSignature(r *http.Request, timeout time.Duration) (string, error) {
	var dt struct {
		Address   string `json:"address"`
		DateTime  string `json:"dateTime"` // YYYYMMDDHHMMSS
		Signature string `json:"sig"`
	}

	if err := web.Decode(r, &dt); err != nil {
		return "", fmt.Errorf("unable to decode payload: %w", err)
	}

	t, err := time.Parse("20060102150405", dt.DateTime)
	if err != nil {
		return "", fmt.Errorf("parse time: %w", err)
	}

	if d := time.Since(t); d > timeout {
		return "", fmt.Errorf("data is too old, %v", d.Seconds())
	}

	data := struct {
		Address  string `json:"address"`
		DateTime string `json:"dateTime"`
	}{
		Address:  dt.Address,
		DateTime: dt.DateTime,
	}

	address, err := ethereum.FromAddress(data, dt.Signature)
	if err != nil {
		return "", fmt.Errorf("unable to extract address: %w", err)
	}

	if !strings.EqualFold(strings.ToLower(address), strings.ToLower(data.Address)) {
		return "", fmt.Errorf("invalid address match, got[%s] exp[%s]", address, data.Address)
	}

	return address, nil
}

func generateToken(a *auth.Auth, address string) (string, error) {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   address,
			Issuer:    "bets project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
	}

	token, err := a.GenerateToken(claims)
	if err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}

	return token, nil
}
