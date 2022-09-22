package bet

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
	"github.com/ardanlabs/bets/business/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Core manages the set of APIs for product access.
type Core struct {
	store db.Store
}

// NewCore constructs a core for product api access.
func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, sqlxDB),
	}
}

// Create adds a Bet to the database.
func (c Core) Create(ctx context.Context, nb NewBet, now time.Time) (Bet, error) {
	if err := validate.Check(nb); err != nil {
		return Bet{}, fmt.Errorf("validating data: %w", err)
	}
	dbBet := db.Bet{
		ID: nb.BetID,
	}
	return betFromDB(dbBet), nil
}
