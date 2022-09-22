package db

import (
	"context"
	"fmt"

	"github.com/ardanlabs/bets/business/sys/database"
)

// AddPlayer adds a new player to an existing bet.
func (s Store) AddPlayer(ctx context.Context, player BetPlayer) error {
	const q = `
	START TRANSACTION;

	-- Ensure the player exists in the accounts table.
	INSERT INTO accounts
			(address, nonce)
	VALUES
			(:address, 0)
	ON CONFLICT DO NOTHING;

	-- Add the player to the bet.
	INSERT INTO bets_players
			(bet_id, address, in_favor)
	VALUES
			(:bet_id, :address, :in_favor);
	COMMIT;`

	if err := database.NamedExecContext(ctx, s.log, s.db, q, player); err != nil {
		return fmt.Errorf("adding player to bet: %w", err)
	}

	return nil
}

// QueryBetPlayers queries bets_players by bet ID.
func (s Store) QueryBetPlayers(ctx context.Context, betID string, pageNumber int, rowsPerPage int) ([]BetPlayer, error) {
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
			bets_players
	WHERE
			bet_id = :bet_id
	ORDER BY
			address
	OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY`

	var players []BetPlayer
	if err := database.NamedQuerySlice(ctx, s.log, s.db, q, data, &players); err != nil {
		return []BetPlayer{}, fmt.Errorf("selecting players by bet: %w", err)
	}

	return players, nil
}
