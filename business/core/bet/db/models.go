package db

import "time"

// Account represents an individual account.
type Account struct {
	Address string `db:"address"`
	Nonce   int    `db:"nonce"`
}

// ============================================================================

// Bet represents an individual bet.
type Bet struct {
	ID               string    `db:"bet_id"`
	Status           string    `db:"status_id"`
	Description      string    `db:"description"`
	Terms            string    `db:"terms"`
	Amount           int       `db:"amount"`
	ModeratorAddress string    `db:"moderator_address"`
	DateExpired      time.Time `db:"date_expired"`
	DateCreated      time.Time `db:"date_created"`
	DateUpdated      time.Time `db:"date_updated"`
}

// BetPlayer represents the connection between a Bet and an Account that is in
// a player role.
type BetPlayer struct {
	BetID   string `db:"bet_id"`
	Address string `db:"address"`
	InFavor bool   `db:"in_favor"`
}

// BetSignature represents an individual signature by an Account on a Bet.
type BetSignature struct {
	BetID      string    `db:"bet_id"`
	Address    string    `db:"address"`
	Nonce      int       `db:"nonce"`
	Signature  string    `db:"signature"`
	DateSigned time.Time `db:"date_signed"`
}
