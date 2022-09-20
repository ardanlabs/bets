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

// Bet struct type
type Bet struct {
	ID             int    `json:"id"`
	Status         string `json:"status"`
	Placer         string `json:"placer"`
	Challenger     string `json:"challenger"`
	Moderator      string `json:"moderator"`
	Description    string `json:"description"`
	Terms          string `json:"terms"`
	ExpirationDate string `json:"expirationDate"`
	Amount         int    `json:"amount"`
}

// Query returns a list of bets with paging.
func (h *Handlers) Query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// Mockup bets slice
	bets := []Bet{
		{
			ID:             1,
			Status:         "signing",
			Placer:         "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC",
			Challenger:     "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
			Description:    "In 2022 there will be 2000 electric cars accidents",
			Terms:          "Has to be in the us.",
			ExpirationDate: "Fri Sep 16 2022",
			Amount:         30,
		},
		{
			ID:             2,
			Status:         "moderate",
			Placer:         "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC",
			Challenger:     "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
			Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
			Description:    "In 2022 there will be 2000 electric cars accidents",
			Terms:          "Has to be in the us.",
			ExpirationDate: "Fri Sep 16 2022",
			Amount:         30,
		},
		{
			ID:             3,
			Status:         "open",
			Placer:         "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC",
			Challenger:     "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
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
	bet := Bet{
		ID:             1,
		Status:         "signing",
		Placer:         "0x3c11fDf93a2Ec67E455C67DaaAdA0550C4bDA4FC",
		Challenger:     "0x0070742FF6003c3E809E78D524F0Fe5dcc5BA7F7",
		Moderator:      "0x39249126d90671284cd06495d19C04DD0e54d371",
		Description:    "In 2022 there will be 2000 electric cars accidents",
		Terms:          "Has to be in the us.",
		ExpirationDate: "Fri Sep 16 2022",
		Amount:         30,
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

// PersonSignBet handles the users signing. Returns the httpStatusCode
func (h *Handlers) PersonSignBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// ModSignBet handles the mod signing. Returns the httpStatusCode
func (h *Handlers) ModSignBet(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}

// setWinner handles a moderator request to set a winner. Takes the betId,
// the winner address, and the signature of the mod. Returns the httpStatusCode
func (h *Handlers) SetWinner(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
