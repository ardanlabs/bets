package bet

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
	"github.com/ardanlabs/bets/business/sys/validate"
)

// Create adds a Bet to the database.
func (c Core) CreateBet(ctx context.Context, nb NewBet, now time.Time) (Bet, error) {
	if err := validate.Check(nb); err != nil {
		return Bet{}, fmt.Errorf("validating data: %w", err)
	}
	dbBet := db.Bet{
		ID: nb.BetID,
	}
	return toBet(dbBet), nil
}
