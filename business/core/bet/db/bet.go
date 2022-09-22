package db

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/bets/business/sys/database"
)

// CreateBet inserts a new Bet into the database.
func (s Store) CreateBet(ctx context.Context, bet Bet) error {
	const q = `
	INSERT INTO bets
			(bet_id, status_id, description, terms, amount, moderator_address, date_expired, date_created, date_updated)
	VALUES
			(:bet_id, :status_id, :description, :terms, :amount, :moderator_address, :date_expired, :date_created, :date_updated);`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, bet); err != nil {
		return fmt.Errorf("inserting bet: %w", err)
	}

	return nil
}

// UpdateBet replaces a bet in the database.
func (s Store) UpdateBet(ctx context.Context, bet Bet) error {
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
			"bet_id" = :bet_id;`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, bet); err != nil {
		return fmt.Errorf("updating bet: %w", err)
	}

	return nil
}

// DeleteBet removes a bet from the database.
func (s Store) DeleteBet(ctx context.Context, betID string) error {
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

// QueryBet retrieves a list of existing bets from the database.
func (s Store) QueryBet(ctx context.Context, pageNumber int, rowsPerPage int) ([]Bet, error) {
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

// QueryBetByID gets the specified bet from the database.
func (s Store) QueryBetByID(ctx context.Context, betID string) (Bet, error) {
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

// QueryBetByModeratorAddress retrieves all bets with the specified address as
// a moderator for the bet.
func (s Store) QueryBetByModeratorAddress(ctx context.Context, moderatorAddress string, pageNumber int, rowsPerPage int) ([]Bet, error) {
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
		return []Bet{}, fmt.Errorf("selecting moderator bets: %w", err)
	}

	return bets, nil
}

// QueryBetByPlayerAddress retrieves all bets filtering by a player address.
func (s Store) QueryBetByPlayerAddress(ctx context.Context, playerAddress string, pageNumber int, rowsPerPage int) ([]Bet, error) {
	data := struct {
		Offset        int    `db:"offset"`
		RowsPerPage   int    `db:"rows_per_page"`
		PlayerAddress string `db:"player_address"`
	}{
		Offset:        (pageNumber - 1) * rowsPerPage,
		RowsPerPage:   rowsPerPage,
		PlayerAddress: playerAddress,
	}

	const q = `
	SELECT
			bets.*
	FROM
			bets
      LEFT JOIN bets_players ON bets_players.bet_id = bets.bet_id
	WHERE
			bets_players.address = :player_address
	ORDER BY
			bets.bet_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var bets []Bet
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &bets); err != nil {
		return []Bet{}, fmt.Errorf("selecting player bets: %w", err)
	}

	return bets, nil
}

// QueryBetByStatus queries bets by their status.
func (s Store) QueryBetByStatus(ctx context.Context, status string, pageNumber int, rowsPerPage int) ([]Bet, error) {
	data := struct {
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
		Status      string `db:"status"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
		Status:      status,
	}

	const q = `
	SELECT
			*
	FROM
			bets
	WHERE
			status = :status
	ORDER BY
			bet_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var bets []Bet
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &bets); err != nil {
		return []Bet{}, fmt.Errorf("selecting bets by status: %w", err)
	}

	return bets, nil
}

// QueryBetByExpiration queries bets by expiration date by providing a start and
// end time.
func (s Store) QueryBetByExpiration(ctx context.Context, start, end time.Time, pageNumber int, rowsPerPage int) ([]Bet, error) {
	data := struct {
		Offset      int       `db:"offset"`
		RowsPerPage int       `db:"rows_per_page"`
		Start       time.Time `db:"start"`
		End         time.Time `db:"end"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
		Start:       start,
		End:         end,
	}

	const q = `
	SELECT
			*
	FROM
			bets
	WHERE
			date_expired >= :start
			AND
			date_expired <= :end
	ORDER BY
			bet_id
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var bets []Bet
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &bets); err != nil {
		return []Bet{}, fmt.Errorf("selecting bets by expiration: %w", err)
	}

	return bets, nil
}
