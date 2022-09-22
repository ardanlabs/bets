package bet

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
	"github.com/ardanlabs/bets/business/sys/validate"
)

// Create adds a Bet to the database.
func (c Core) Create(ctx context.Context, nb NewBet, now time.Time) (Bet, error) {
	if err := validate.Check(nb); err != nil {
		return Bet{}, fmt.Errorf("validating data: %w", err)
	}

	// If moderator is provided, make sure it does exist as an account.
	if nb.ModeratorAddress != "" {
		acc := db.Account{
			Address: nb.ModeratorAddress,
			Nonce:   0,
		}
		if err := c.store.CreateAccount(ctx, acc); err != nil {
			return Bet{}, fmt.Errorf("create moderator account: %w", err)
		}
	}

	// Ensures that the players accounts exists.
	for _, player := range nb.Players {
		acc := db.Account{
			Address: player.Address,
			Nonce:   0,
		}
		if err := c.store.CreateAccount(ctx, acc); err != nil {
			return Bet{}, fmt.Errorf("create player account: %w", err)
		}
	}

	// Create the bet.
	dbBet := db.Bet{
		ID:               validate.GenerateID(),
		Status:           "negotiation",
		Description:      nb.Description,
		Terms:            nb.Terms,
		Amount:           nb.Amount,
		ModeratorAddress: nb.ModeratorAddress,
		DateExpired:      nb.DateExpired,
		DateCreated:      now,
		DateUpdated:      now,
	}

	if err := c.store.CreateBet(ctx, dbBet); err != nil {
		return Bet{}, fmt.Errorf("create bet: %w", err)
	}

	// Add the players into the bet.
	for _, player := range nb.Players {
		acc := db.BetPlayer{
			BetID:   dbBet.ID,
			Address: player.Address,
			InFavor: player.InFavor,
		}
		if err := c.store.AddPlayer(ctx, acc); err != nil {
			return Bet{}, fmt.Errorf("create player account: %w", err)
		}
	}

	// Build the bet response content.
	bet := Bet{
		ID:               dbBet.ID,
		Status:           dbBet.Status,
		Description:      dbBet.Description,
		Terms:            dbBet.Terms,
		Amount:           dbBet.Amount,
		ModeratorAddress: dbBet.ModeratorAddress,
		DateExpired:      dbBet.DateExpired,
		DateCreated:      dbBet.DateCreated,
		DateUpdated:      dbBet.DateUpdated,
	}

	var players []BetPlayer
	for _, player := range nb.Players {
		players = append(players, BetPlayer{
			BetID:   dbBet.ID,
			Address: player.Address,
			InFavor: player.InFavor,
		})
	}

	bet.Players = players

	return bet, nil
}
