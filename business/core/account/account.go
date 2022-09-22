package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardanlabs/bets/business/core/account/db"
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

// Core manages the set of APIs for account access.
type Core struct {
	store db.Store
}

// NewCore constructs a core for account api access.
func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		store: db.NewStore(log, sqlxDB),
	}
}

// Create adds an Account to the database. It returns the created Account with
// fields populated.
func (c Core) Create(ctx context.Context, na NewAccount) (Account, error) {
	if err := validate.Check(na); err != nil {
		return Account{}, fmt.Errorf("validating data: %w", err)
	}

	dbAccount := db.Account{
		Address: na.Address,
		Nonce:   0,
	}

	if err := c.store.Create(ctx, dbAccount); err != nil {
		return Account{}, fmt.Errorf("create: %w", err)
	}

	return toAccount(dbAccount), nil
}

// Update modifies data about an Account. It will error if the specified address
// is invalid or does not reference an existing Account.
func (c Core) Update(ctx context.Context, address string, ua UpdateAccount) error {
	if err := validate.CheckAddress(address); err != nil {
		return ErrInvalidAddress
	}

	if err := validate.Check(ua); err != nil {
		return fmt.Errorf("validating data: %w", err)
	}

	dbAccount, err := c.store.QueryByAddress(ctx, address)
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

	if err := c.store.Update(ctx, dbAccount); err != nil {
		return fmt.Errorf("update: %w", err)
	}

	return nil
}
