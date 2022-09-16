package brunogrp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/bets/foundation/web"
)

// Handlers manages the set of user endpoints.
type Handlers struct {
}

// Query returns a list of bets with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	type Bet struct {
		ID                int    `json:"id"`
		Description       string `json:"description"`
		Status						string `json:"status"`
		Terms             string `json:"terms"`
		Name              string `json:"name"`
		ChallengerAddress string `json:"challengerAddress"`
		ExpirationDate    string `json:"expirationDate"`
		Amount            int    `json:"amount"`
	}

	bets := []Bet{
		{
			ID:                1,
			Description:       "In 2022 there will be 2000 electric cars accidents",
			Terms:             "Has to be in the us.",
			Name:              "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate:    "20221231000000",
			Amount:            30,
		},
		{
			ID:                2,
			Description:       "Lorem ipsum dolor sit amet consectetur adipisicing elit. Autem, maxime!",
			Terms:             "Has to be in the us.",
			Name:              "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate:    "20221231000000",
			Amount:            30,
		},
		{
			ID:                3,
			Description:       "Temporibus ratione doloremque dolorum atque? Incidunt dolore ipsa cum nobis quo enim?",
			Terms:             "Has to be in the us.",
			Name:              "Bruno",
			ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			ExpirationDate:    "20221231000000",
			Amount:            30,
		},
	}
	return web.Respond(ctx, w, bets, http.StatusOK)
}

// QueryByID returns a bet by its ID.
func (h *Handlers) QueryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	bet := struct {
		ID                int    `json:"id"`
		Description       string `json:"description"`
		Terms             string `json:"terms"`
		Name              string `json:"name"`
		ChallengerAddress string `json:"challengerAddress"`
		ExpirationDate    string `json:"expirationDate"`
		Amount            int    `json:"amount"`
	}{
		ID:                1,
		Description:       "In 2022 there will be 2000 electric cars accidents",
		Terms:             "Has to be in the us.",
		Name:              "Bruno",
		ChallengerAddress: "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
		ExpirationDate:    "20221231000000",
		Amount:            30,
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
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// AcceptMod handles the user accepting an address as moderator. Returns the httpStatusCode
func (h *Handlers) AcceptMod(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
