package bet

import (
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
)

// Bet represents an individual bet.
type Bet struct {
	ID               string      `json:"id"`
	Status           string      `json:"status"`
	Description      string      `json:"description"`
	Terms            string      `json:"terms"`
	Amount           int         `json:"amount"`
	ModeratorAddress string      `json:"moderator_address"`
	Players          []BetPlayer `json:"players"`
	DateExpired      time.Time   `json:"date_expired"`
	DateCreated      time.Time   `json:"date_created"`
	DateUpdated      time.Time   `json:"date_updated"`
}

// BetPlayer represents the connection between a Bet and an Account that is in
// a player role.
type BetPlayer struct {
	BetID   string `db:"bet_id"`
	Address string `db:"address"`
	InFavor bool   `db:"in_favor"`
}

// NewBet is what we require from clients when adding a Bet.
type NewBet struct {
	Description      string         `json:"description"`
	Terms            string         `json:"terms"`
	Amount           int            `json:"amount"`
	ModeratorAddress string         `json:"moderator_address"`
	Players          []NewBetPlayer `json:"players" validate:"required"`
	DateExpired      time.Time      `json:"date_expired"`
}

// NewBetPlayer represents the connection between a new Bet and an Account that is in
// a player role.
type NewBetPlayer struct {
	Address string `db:"address"`
	InFavor bool   `db:"in_favor"`
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
