package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/bets/business/sys/database"
)

// AddSignature adds a new player signature to an existing bet.
func (s Store) AddSignature(ctx context.Context, signature BetSignature) error {
	const q = `
	INSERT INTO bets_signatures
			(bet_id, address, nonce, signature, date_signed)
	VALUES
			(:bet_id, :address, :nonce, :signature, :date_signed)`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, signature); err != nil {
		return fmt.Errorf("adding signature to bet: %w", err)
	}

	return nil
}

// QueryBetSignatures queries bets_signatures by bet ID.
func (s Store) QueryBetSignatures(ctx context.Context, betID string, pageNumber int, rowsPerPage int) ([]BetSignature, error) {
	data := struct {
		Offset      int    `db:"offset"`
		RowsPerPage int    `db:"rows_per_page"`
		BetID       string `db:"bet_id"`
	}{
		Offset:      (pageNumber - 1) * rowsPerPage,
		RowsPerPage: rowsPerPage,
		BetID:       betID,
	}

	const q = `
	SELECT
			*
	FROM
			bets_signatures
	WHERE
			bet_id = :bet_id
	ORDER BY
			address
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var signatures []BetSignature
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &signatures); err != nil {
		return []BetSignature{}, fmt.Errorf("selecting signatures by bet: %w", err)
	}

	return signatures, nil
}
