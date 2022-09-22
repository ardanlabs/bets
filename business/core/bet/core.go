package bet

import (
	"errors"

	"github.com/ardanlabs/bets/business/core/bet/db"
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
