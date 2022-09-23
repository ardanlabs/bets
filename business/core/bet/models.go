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
	ModeratorAddress string    `json:"moderatorAddress"`
	Players          []Player  `json:"players"`
	DateExpired      time.Time `json:"dateExpired"`
	DateCreated      time.Time `json:"dateCreated"`
	DateUpdated      time.Time `json:"dateUpdated"`
}

// Player represents the connection between a Bet and an Account that is in
// a player role.
type Player struct {
	BetID   string `db:"betId"`
	Address string `db:"address"`
	InFavor bool   `db:"inFavor"`
}

// NewBet is what we require from clients when adding a Bet.
type NewBet struct {
	Description      string      `json:"description"`
	Terms            string      `json:"terms"`
	Amount           int         `json:"amount"`
	ModeratorAddress string      `json:"moderatorAddress"`
	Players          []NewPlayer `json:"players" validate:"required"`
	DateExpired      time.Time   `json:"dateExpired"`
}

// NewPlayer represents the connection between a new Bet and an Account that is in
// a player role.
type NewPlayer struct {
	Address string `db:"address"`
	InFavor bool   `db:"inFavor"`
}

// UpdateBet is what we require from clients when updating a Bet.
type UpdateBet struct {
	Description      *string   `json:"description"`
	Terms            *string   `json:"terms"`
	Amount           *int      `json:"amount"`
	ModeratorAddress *string   `json:"moderatorAddress"`
	DateExpired      time.Time `json:"dateExpired"`
}

// ============================================================================

// Account represents an individual account
type Account struct {
	Address string `json:"address"`
	Nonce   int    `db:"nonce"`
}

// NewAccount contains information needed to create a new Account.
type NewAccount struct {
	Address string `json:"address"`
}

// UpdateAccount contains information needed to update an Account.
type UpdateAccount struct {
	Address string `json:"address"`
	Nonce   int    `json:"nonce"`
}

// ============================================================================

func toAccount(dbAccount db.Account) Account {
	return Account(dbAccount)
}

func toAccountSlice(dbAccounts []db.Account) []Account {
	accounts := make([]Account, len(dbAccounts))
	for i, dbAccount := range dbAccounts {
		accounts[i] = toAccount(dbAccount)
	}
	return accounts
}
