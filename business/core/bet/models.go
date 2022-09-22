package bet

import (
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
)

// Bet represents an individual bet.
type Bet struct {
	ID               string    `json:"id"`
	Status           string    `json:"status"`
	Description      string    `json:"description"`
	Terms            string    `json:"terms"`
	Amount           int       `json:"amount"`
	ModeratorAddress string    `json:"moderator_address"`
	DateExpired      time.Time `json:"date_expired"`
	DateCreated      time.Time `json:"date_created"`
	DateUpdated      time.Time `json:"date_updated"`
}

// NewBet is what we require from clients when adding a Bet.
type NewBet struct {
	BetID string `json:"betId" validate:"required"`
}

func betFromDB(bet db.Bet) Bet {
	return Bet(bet)
}
