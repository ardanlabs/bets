package db

// Account represents an individual account.
type Account struct {
	Address string `db:"address"`
	Nonce   int    `db:"nonce"`
}
