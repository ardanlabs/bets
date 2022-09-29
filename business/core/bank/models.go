package bank

import (
	"math/big"
	"time"
)

// NewBet represents the input required to place a bet.
type NewBet struct {
	AmountGWei    *big.Float
	FeeAmountGWei *big.Float
	Expiration    time.Time
	Moderator     string
	Participants  []string
	Nonces        []*big.Int
	Signatures    [][]byte
}

// Validate verifies the new bet value is properly initialized.
func (nb NewBet) Validate() error {
	return nil
}

// ReconcileBet represents the input required to reconcile a bet.
type ReconcileBet struct {
	Nonce     *big.Int
	Moderator string
	Signature []byte
	Winners   []string
}

// Validate verifies the reconcile bet value is properly initialized.
func (rb ReconcileBet) Validate() error {
	return nil
}

// CancelBetModerator represents the input required to cancel a bet by the moderator.
type CancelBetModerator struct {
	FeeAmountGWei *big.Float
	Nonce         *big.Int
	Moderator     string
	Signature     []byte
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbm CancelBetModerator) Validate() error {
	return nil
}

// CancelBetParticipants represents the input required to cancel a bet by the participants.
type CancelBetParticipants struct {
	FeeAmountGWei *big.Float
	Nonces        []*big.Int
	Signatures    [][]byte
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbp CancelBetParticipants) Validate() error {
	return nil
}

// CancelBetOwner represents the input required to cancel a bet by the contract owner.
type CancelBetOwner struct {
	FeeAmountGWei *big.Float
}

// Validate verifies the reconcile bet value is properly initialized.
func (cbo CancelBetOwner) Validate() error {
	return nil
}
