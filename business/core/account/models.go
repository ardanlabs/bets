package account

import (
	"github.com/ardanlabs/bets/business/core/account/db"
)

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
