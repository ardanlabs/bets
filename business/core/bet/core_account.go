package bet

import (
	"context"
	"errors"
	"fmt"

	"github.com/ardanlabs/bets/business/core/bet/db"
	"github.com/ardanlabs/bets/business/sys/database"
	"github.com/ardanlabs/bets/business/sys/validate"
)

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
