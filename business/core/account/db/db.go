package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/bets/business/sys/database"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for bet access.
type Store struct {
	log          *zap.SugaredLogger
	tr           database.Transactor
	db           sqlx.ExtContext
	isWithinTran bool
}

// NewStore constructs a data for api access.
func NewStore(log *zap.SugaredLogger, db *sqlx.DB) Store {
	return Store{
		log: log,
		tr:  db,
		db:  db,
	}
}

// WithinTran runs passed function and do commit/rollback at the end.
func (s Store) WithinTran(ctx context.Context, fn func(sqlx.ExtContext) error) error {
	if s.isWithinTran {
		fn(s.db)
	}
	return database.WithinTran(ctx, s.log, s.tr, fn)
}

// Tran return new Store with transaction in it.
func (s Store) Tran(tx sqlx.ExtContext) Store {
	return Store{
		log:          s.log,
		tr:           s.tr,
		db:           tx,
		isWithinTran: true,
	}
}

// Create inserts a new Account into the database.
func (s Store) Create(ctx context.Context, account Account) error {
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

// Update updates an existing account in the database.
func (s Store) Update(ctx context.Context, account Account) error {
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

// Query retrieves a list of existing accounts from the database.
func (s Store) Query(ctx context.Context, pageNumber, rowsPerPage int) ([]Account, error) {
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

// QueryByAddress retrieves an account by address.
func (s Store) QueryByAddress(ctx context.Context, address string) (Account, error) {
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
