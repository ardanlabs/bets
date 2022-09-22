package brunogrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ardanlabs/bets/business/core/bet"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/google/uuid"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
	Bet bet.Core
}

type Player struct {
	Address string `json:"address"`
	Signed  bool   `json:"signed"`
}

// Bet struct type
type Bet struct {
	ID             string   `json:"id"`
	Status         string   `json:"status"`
	Players        []Player `json:"players"`
	Moderator      string   `json:"moderator"`
	Description    string   `json:"description"`
	Terms          string   `json:"terms"`
	ExpirationDate string   `json:"expirationDate"`
	Amount         int      `json:"amount"`
}

// Query returns a list of bets with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Mockup bets slice

	players := []Player{
		{Address: "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC", Signed: false},
		{Address: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", Signed: false},
	}
	bets := []Bet{
		{
			ID:             uuid.NewString(),
			Status:         "signing",
			Players:        players,
			Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
			Description:    "In 2022 there will be 2000 electric cars accidents",
			Terms:          "Has to be in the us.",
			ExpirationDate: "Fri Sep 16 2022",
			Amount:         30,
		},
		{
			ID:             uuid.NewString(),
			Status:         "moderate",
			Players:        players,
			Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
			Description:    "In 2022 there will be 2000 electric cars accidents",
			Terms:          "Has to be in the us.",
			ExpirationDate: "Fri Sep 16 2022",
			Amount:         30,
		},
		{
			ID:             uuid.NewString(),
			Status:         "open",
			Players:        players,
			Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
			Description:    "In 2022 there will be 2000 electric cars accidents",
			Terms:          "Has to be in the us.",
			ExpirationDate: "Fri Sep 16 2022",
			Amount:         30,
		},
	}

	return web.Respond(ctx, w, bets, http.StatusOK)
}

// QueryByID returns a bet by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	players := []Player{
		{Address: "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC", Signed: false},
		{Address: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7", Signed: false},
	}

	bet := Bet{
		ID:             uuid.NewString(),
		Status:         "signing",
		Players:        players,
		Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
		Description:    "In 2022 there will be 2000 electric cars accidents",
		Terms:          "Has to be in the us.",
		ExpirationDate: "Fri Sep 16 2022",
		Amount:         30,
	}
	return web.Respond(ctx, w, bet, http.StatusOK)
}

// Create creates a bet and returns its ID.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, err := web.GetValues(ctx)
	if err != nil {
		return web.NewShutdownError("web value missing from context")
	}

	nb := bet.NewBet{
		BetID: uuid.NewString(),
	}

	bet, err := h.Bet.Create(ctx, nb, v.Now)
	if err != nil {
		return fmt.Errorf("creating new bet, newBet[%+v]: %w", nb, err)
	}

	return web.Respond(ctx, w, bet, http.StatusCreated)
}

// PersonSignBet handles the users signing. Returns the httpStatusCode
func (h *Handlers) PersonSignBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// ModSignBet handles the mod signing. Returns the httpStatusCode
func (h *Handlers) ModSignBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// SetWinner handles a moderator request to set a winner. Takes the betId,
// the winner address, and the signature of the mod. Returns the httpStatusCode
func (h *Handlers) SetWinner(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
