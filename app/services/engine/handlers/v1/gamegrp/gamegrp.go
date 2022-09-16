// Package gamegrp provides the handlers for game play.
package gamegrp

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ardanlabs/bets/business/web/auth"
	v1Web "github.com/ardanlabs/bets/business/web/v1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"

	"github.com/ardanlabs/bets/business/core/bank"
	"github.com/ardanlabs/bets/foundation/events"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
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

// Query returns a list of bets with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	type Bet struct {
		ID int `json:"id"`
		Description string `json:"description"`
		Terms string `json:"terms"`
		Name string `json:"name"`
		ChallengerAddress string `json:"challengerAddress"`
		ExpirationDate string `json:"expirationDate"`
		Amount int `json:"amount"`
	}

	bets := []Bet{
		Bet{
			ID: 1,
			Description: "In 2022 there will be 2000 electric cars accidents",
			Terms: "Has to be in the us.",
			Name: "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate: "20221231000000",
			Amount: 30,
		},
		Bet{
			ID: 2,
			Description: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Autem, maxime!",
			Terms: "Has to be in the us.",
			Name: "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate: "20221231000000",
			Amount: 30,
		},
		Bet{
			ID: 3,
			Description: "Temporibus ratione doloremque dolorum atque? Incidunt dolore ipsa cum nobis quo enim?",
			Terms: "Has to be in the us.",
			Name: "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate: "20221231000000",
			Amount: 30,
		},
	}
	return web.Respond(ctx, w, bets, http.StatusOK)
}

// QueryByID returns a bet by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	bet := struct {
		ID int `json:"id"`
		Description string `json:"description"`
		Terms string `json:"terms"`
		Name string `json:"name"`
		ChallengerAddress string `json:"challengerAddress"`
		ExpirationDate string `json:"expirationDate"`
		Amount int `json:"amount"`
	}{
    ID: 1,
    Description: "In 2022 there will be 2000 electric cars accidents",
    Terms: "Has to be in the us.",
    Name: "Bruno",
    ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
    ExpirationDate: "20221231000000",
    Amount: 30,
	}

	return web.Respond(ctx, w, bet, http.StatusOK)
}

// Create creates a bet and returns it's ID.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	betID := struct {
		BetID int `json:"betId"`
	}{
		BetID: rand.Intn(100),
	}

	return web.Respond(ctx, w, betID, http.StatusOK)
}

// SignBet handles all bet signing. Returns the httpStatusCode
func (h *Handlers) SignBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return http.StatusOK
}

// AcceptMod handles the user accepting an address as moderator. Returns the httpStatusCode
func (h *Handlers) AcceptMod(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return http.StatusOK
}
