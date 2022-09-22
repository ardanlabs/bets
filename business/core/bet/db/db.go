package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/bets/business/sys/database"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Store manages the set of APIs for product access.
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

// Create inserts a new Bet into the database.
func (s Store) Create(ctx context.Context, bet Bet) error {
	const q = `
	INSERT INTO bets
			(bet_id, status_id, description, terms, amount, moderator_address, date_expired, date_created, date_updated)
	VALUES
			(:bet_id, :status_id, :description, :terms, :amount, :moderator_address, :date_expired, :date_created, :date_updated)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, bet); err != nil {
		return fmt.Errorf("inserting bet: %w", err)
	}

	return nil
}

// Update replaces a bet in the database.
func (s Store) Update(ctx context.Context, bet Bet) error {
	const q = `
	UPDATE
			bets
	SET
			"status"            = :status,
			"description"       = :description,
			"terms"             = :terms,
			"amount"            = :amount,
			"moderator_address" = :moderator_address,
			"date_expired"      = :date_expired,
			"date_created"      = :date_created,
			"date_updated"      = :date_updated
	WHERE
			"bet_id" = :bet_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, bet); err != nil {
		return fmt.Errorf("updating bet: %w", err)
	}

	return nil
}

// Delete removes a bet from the database.
func (s Store) Delete(ctx context.Context, betID string) error {
	data := struct {
		ID string `db:"bet_id"`
	}{
		ID: betID,
	}

	const q = `
	DELETE FROM
			bets
	WHERE
			bet_id = :bet_id`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, data); err != nil {
		return fmt.Errorf("deleting betID[%s]: %w", betID, err)
	}

	return nil
}

// Query retrieves a list of existing bets from the database.
func (s Store) Query(ctx context.Context, pageNumber int, rowsPerPage int) ([]Bet, error) {
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
			bets
	ORDER BY
			bet_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var bets []Bet
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &bets); err != nil {
		return bets, fmt.Errorf("selecting bets: %w", err)
	}

	return bets, nil
}

// QueryByID gets the specified bet from the database.
func (s Store) QueryByID(ctx context.Context, betID string) (Bet, error) {
	data := struct {
		BetID string `db:"bet_id"`
	}{
		BetID: betID,
	}

	const q = `
	SELECT
			*
	FROM
			bets
	WHERE
			bet_id = :bet_id`

	var bet Bet
	if err := database.NamedQueryStruct(ctx, s.log, s.db, q, data, &bet); err != nil {
		return Bet{}, fmt.Errorf("selecting betID[%q]: %w", betID, err)
	}

	return bet, nil
}

// QueryByModeratorAddress retrieves all bets with the specified address as a
// moderator for the bet.
func (s Store) QueryByModeratorAddress(ctx context.Context, moderatorAddress string, pageNumber int, rowsPerPage int) ([]Bet, error) {
	data := struct {
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
		Moderator   string `db:"moderator"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
		Moderator:   moderatorAddress,
	}

	const q = `
	SELECT
			*
	FROM
			bets
	WHERE
			moderator_address = :moderator
	ORDER BY
			bet_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var bets []Bet
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &bets); err != nil {
		return bets, fmt.Errorf("selecting moderator bets: %w", err)
	}

	return bets, nil
}

// QueryByPlayerAddress
func (s Store) QueryByPlayerAddress(ctx context.Context, playerAddress string, pageNumber int, rowsPerPage int) ([]Bet, error) {

}
