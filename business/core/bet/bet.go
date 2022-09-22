// Package bet provides an example of a core business API. Right now these
// calls are just wrapping the data/store layer. But at some point you will
// want to audit or something that isn't specific to the data/store layer.
package bet

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ardanlabs/bets/business/core/bet/db"
	"github.com/ardanlabs/bets/business/sys/database"
	"github.com/ardanlabs/bets/business/sys/validate"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound       = errors.New("account not found")
	ErrInvalidAddress = errors.New("address is not in its proper form")
	ErrInvalidNonce   = errors.New("nonce must be 1 greater than previous nonce")
)

// Core manages the set of APIs for product access.
type Core struct {
	store db.Store
}

// NewCore constructs a core for product api access.
func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, sqlxDB),
	}
}

// =========================================================================
// Account Support

// CreateAccount adds an Account to the database. It returns the created Account
// with fields populated.
func (c Core) CreateAccount(ctx context.Context, na NewAccount) (Account, error) {
	if err := validate.Check(na); err != nil {
		return Account{}, fmt.Errorf("validating data: %w", err)
	}

	dbAccount := db.Account{
		Address: na.Address,
		Nonce:   0,
	}

	if err := c.store.CreateAccount(ctx, dbAccount); err != nil {
		return Account{}, fmt.Errorf("create: %w", err)
	}

	return toAccount(dbAccount), nil
}

// UpdateAccount modifies data about an Account. It will error if the specified
// address is invalid or does not reference an existing Account.
func (c Core) UpdateAccount(ctx context.Context, address string, ua UpdateAccount) error {
	if !validate.CheckAddress(address) {
		return ErrInvalidAddress
	}

	if err := validate.Check(ua); err != nil {
		return fmt.Errorf("validating data: %w", err)
	}

	dbAccount, err := c.store.QueryAccountByAddress(ctx, address)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			return ErrNotFound
		}
		return fmt.Errorf("updating account address[%s]: %w", address, err)
	}

	if ua.Nonce != dbAccount.Nonce+1 {
		return ErrInvalidNonce
	}

	dbAccount.Nonce = ua.Nonce

	if err := c.store.UpdateAccount(ctx, dbAccount); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}

// =========================================================================
// Bet Support

// CreateBet adds a Bet to the database.
func (c Core) CreateBet(ctx context.Context, nb NewBet, now time.Time) (Bet, error) {
	if err := validate.Check(nb); err != nil {
		return Bet{}, fmt.Errorf("validating data: %w", err)
	}

	var dbBet db.Bet

	tran := func(tx sqlx.ExtContext) error {

		// If moderator is provided, make sure it does exist as an account.
		if nb.ModeratorAddress != "" {
			acc := db.Account{
				Address: nb.ModeratorAddress,
				Nonce:   0,
			}
			if err := c.store.Tran(tx).CreateAccount(ctx, acc); err != nil {
				return fmt.Errorf("create moderator account: %w", err)
			}
		}

		// Ensures that the players accounts exists.
		for _, player := range nb.Players {
			acc := db.Account{
				Address: player.Address,
				Nonce:   0,
			}
			if err := c.store.Tran(tx).CreateAccount(ctx, acc); err != nil {
				return fmt.Errorf("create player account: %w", err)
			}
		}

		// Create the bet.
		dbBet.ID = validate.GenerateID()
		dbBet.Status = "negotiation"
		dbBet.Description = nb.Description
		dbBet.Terms = nb.Terms
		dbBet.Amount = nb.Amount
		dbBet.ModeratorAddress = nb.ModeratorAddress
		dbBet.DateExpired = nb.DateExpired
		dbBet.DateCreated = now
		dbBet.DateUpdated = now

		if err := c.store.Tran(tx).CreateBet(ctx, dbBet); err != nil {
			return fmt.Errorf("create bet: %w", err)
		}

		// Add the players into the bet.
		for _, player := range nb.Players {
			acc := db.Player{
				BetID:   dbBet.ID,
				Address: player.Address,
				InFavor: player.InFavor,
			}
			if err := c.store.Tran(tx).AddPlayer(ctx, acc); err != nil {
				return fmt.Errorf("create player account: %w", err)
			}
		}

		return nil
	}

	if err := c.store.WithinTran(ctx, tran); err != nil {
		return Bet{}, fmt.Errorf("tran: %w", err)
	}

	// Build the bet response content.
	bet := Bet{
		ID:               dbBet.ID,
		Status:           dbBet.Status,
		Description:      dbBet.Description,
		Terms:            dbBet.Terms,
		Amount:           dbBet.Amount,
		ModeratorAddress: dbBet.ModeratorAddress,
		DateExpired:      dbBet.DateExpired,
		DateCreated:      dbBet.DateCreated,
		DateUpdated:      dbBet.DateUpdated,
	}

	var players []Player
	for _, player := range nb.Players {
		players = append(players, Player{
			BetID:   dbBet.ID,
			Address: player.Address,
			InFavor: player.InFavor,
		})
	}

	bet.Players = players

	return bet, nil
}
