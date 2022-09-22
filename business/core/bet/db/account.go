package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/bets/business/sys/database"
)

// CreateAccount inserts a new Account into the database.
func (s Store) CreateAccount(ctx context.Context, account Account) error {
	const q = `
	INSERT INTO accounts
			(address, nonce)
	VALUES
			(:address, :nonce)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, account); err != nil {
		return fmt.Errorf("inserting account: %w", err)
	}

	return nil
}

// UpdateAccount updates an existing account in the database.
func (s Store) UpdateAccount(ctx context.Context, account Account) error {
	const q = `
	UPDATE
			accounts
	SET
			nonce = :nonce
	WHERE
			address = :address`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, account); err != nil {
		return fmt.Errorf("updating account: %w", err)
	}

	return nil
}

// QueryAccounts retrieves a list of existing accounts from the database.
func (s Store) QueryAccounts(ctx context.Context, pageNumber, rowsPerPage int) ([]Account, error) {
	data := struct {
		Offset      int `db:"offset"`
		RowsPerPage int `db:"rows_per_page"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
	}

	const q = `
	SELECT
			*
	FROM
			accounts
	ORDER BY
			address
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var accounts []Account
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &accounts); err != nil {
		return []Account{}, fmt.Errorf("selecting accounts: %w", err)
	}

	return accounts, nil
}

// QueryAccountByAddress QueryAccountsByAddress retrieves an account by address.
func (s Store) QueryAccountByAddress(ctx context.Context, address string) (Account, error) {
	data := struct {
		Address string `db:"address"`
	}{
		Address: address,
	}

	const q = `
	SELECT
			*
	FROM
			accounts
	WHERE
			address = :address`

	var account Account
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &account); err != nil {
		return Account{}, fmt.Errorf("selecting account[%q]: %w", address, err)
	}

	return account, nil
}
