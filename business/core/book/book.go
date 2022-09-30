// Package book represents all the transactions necessary for the game.
package book

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ardanlabs/bets/business/contract/go/book"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// These constants define the different bet states.
const (
	StateNotExists  = 0
	StateLive       = 1
	StateReconciled = 2
	StateCancelled  = 3
)

// These constants set a defined gas limit for the different calls.
const (
	gasLimitDrain    = 3_000_000
	gasLimitPlaceBet = 3_000_000
)

// Sign is used to sign data for the different smart contract calls.
func Sign(privateKey *ecdsa.PrivateKey, betID string, address string, nonce uint) ([]byte, error) {
	if !common.IsHexAddress(address) {
		return nil, errors.New("invalid address")
	}

	String, _ := abi.NewType("string", "", nil)
	Address, _ := abi.NewType("address", "", nil)
	Uint, _ := abi.NewType("uint", "", nil)

	arguments := abi.Arguments{
		{
			Type: String,
		},
		{
			Type: Address,
		},
		{
			Type: Uint,
		},
	}

	bytes, err := arguments.Pack(betID, common.HexToAddress(address), big.NewInt(int64(nonce)))
	if err != nil {
		return nil, fmt.Errorf("arguments pack: %w", err)
	}

	signature, err := ethereum.SignBytes(bytes, privateKey)
	if err != nil {
		return nil, fmt.Errorf("signing message: %w", err)
	}

	sig, err := hex.DecodeString(signature[2:])
	if err != nil {
		return nil, fmt.Errorf("decoding signature string: %w", err)
	}

	return sig, nil
}

// =============================================================================

// Book represents a bank that allows for the reconciling of a game and
// information about account balances.
type Book struct {
	logger     *zap.SugaredLogger
	contractID string
	ethereum   *ethereum.Ethereum
	contract   *book.Book
}

// New returns a new bank with the ability to manage the game money.
func New(ctx context.Context, logger *zap.SugaredLogger, network string, keyPath string, passPhrase string, contractID string) (*Book, error) {
	ethereum, err := ethereum.New(ctx, network, keyPath, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("network connect: %w", err)
	}

	contract, err := book.NewBook(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return nil, fmt.Errorf("new contract: %w", err)
	}

	b := Book{
		logger:     logger,
		contractID: contractID,
		ethereum:   ethereum,
		contract:   contract,
	}

	b.log(ctx, "new bank", "network", network, "contractid", contractID)

	return &b, nil
}

// ContractID returns contract id in use.
func (b *Book) ContractID() string {
	return b.contractID
}

// Client returns the underlying contract client.
func (b *Book) Client() *ethereum.Ethereum {
	return b.ethereum
}

// =============================================================================
// Owner Only Calls

// Drain will drain the full value of the smart contract and transfer it to
// the owner address.
func (b *Book) Drain(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, gasLimitDrain, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Drain(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("drain: %w", err)
	}

	b.log(ctx, "drain started")
	defer b.log(ctx, "drain completed")

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// AccountBalance will return the balance for the specified account. Only the
// owner of the smart contract can make this call.
func (b *Book) AccountBalance(ctx context.Context, accountID string) (BalanceGWei *big.Float, err error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new call opts: %w", err)
	}

	balance, err := b.contract.AccountBalance(tranOpts, common.HexToAddress(accountID))
	if err != nil {
		return nil, fmt.Errorf("account balance: %w", err)
	}

	b.log(ctx, "account balance", "accountid", accountID, "balance", balance)

	return currency.Wei2GWei(balance), nil
}

// Nonce will return the current nonce for the specified account.
func (b *Book) Nonce(ctx context.Context, accountID string) (*big.Int, error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new call opts: %w", err)
	}

	nonce, err := b.contract.Nonce(tranOpts, common.HexToAddress(accountID))
	if err != nil {
		return nil, fmt.Errorf("account balance: %w", err)
	}

	b.log(ctx, "get nonce", "accountid", b.ethereum.Address().String(), "nonce", nonce)

	return nonce, nil
}

// Nonce will return the current nonce for the specified account.
func (b *Book) BetDetails(ctx context.Context, betID string) (BetInfo, error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return BetInfo{}, fmt.Errorf("new call opts: %w", err)
	}

	bbi, err := b.contract.BetDetails(tranOpts, betID)
	if err != nil {
		return BetInfo{}, fmt.Errorf("account balance: %w", err)
	}

	participants := make([]string, len(bbi.Participants))
	for i, participant := range bbi.Participants {
		participants[i] = participant.Hex()
	}

	betInfo := BetInfo{
		State:         int(bbi.State),
		Participants:  participants,
		Moderator:     bbi.Moderator.Hex(),
		AmountBetGWei: currency.Wei2GWei(bbi.AmountBetWei),
		Expiration:    time.Unix(bbi.Expiration.Int64(), 0),
	}

	b.log(ctx, "bet details", "betid", betID, "details", betInfo)

	return betInfo, nil
}

// PlaceBet adds a new bet to the smart contract. This bet will be live if accepted
// by the smart contract and all participants will be bound to the bet.
func (b *Book) PlaceBet(ctx context.Context, betID string, pb PlaceBet) (*types.Transaction, *types.Receipt, error) {
	if err := pb.Validate(); err != nil {
		return nil, nil, fmt.Errorf("validate input: %w", err)
	}

	var participants []common.Address
	for _, participant := range pb.Participants {
		participants = append(participants, common.HexToAddress(participant))
	}

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, gasLimitPlaceBet, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.PlaceBet(
		tranOpts,
		betID,
		currency.GWei2Wei(pb.AmountBetGWei),
		currency.GWei2Wei(pb.AmountFeeGWei),
		new(big.Int).SetInt64(pb.Expiration.Unix()),
		common.HexToAddress(pb.Moderator),
		participants,
		pb.Nonces,
		pb.Signatures,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("place bet: %w", err)
	}

	b.log(ctx, "place bet started", "betID", betID)
	defer b.log(ctx, "place bet completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// ReconcileBet allows the moderator to sign off on the live bet and provide
// the winning accounts.
func (b *Book) ReconcileBet(ctx context.Context, betID string, rb ReconcileBet) (*types.Transaction, *types.Receipt, error) {
	if err := rb.Validate(); err != nil {
		return nil, nil, fmt.Errorf("validate input: %w", err)
	}

	var winners []common.Address
	for _, winner := range rb.Winners {
		winners = append(winners, common.HexToAddress(winner))
	}

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.ReconcileBet(
		tranOpts,
		betID,
		rb.Nonce,
		rb.Signature,
		winners,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("reconcile bet: %w", err)
	}

	b.log(ctx, "reconcile bet started", "betID", betID)
	defer b.log(ctx, "reconcile bet completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// CancelBetModerator allows the moderator to sign off on cancelling the bet.
func (b *Book) CancelBetModerator(ctx context.Context, betID string, cbm CancelBetModerator) (*types.Transaction, *types.Receipt, error) {
	if err := cbm.Validate(); err != nil {
		return nil, nil, fmt.Errorf("validate input: %w", err)
	}

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.CancelBetModerator(
		tranOpts,
		betID,
		currency.GWei2Wei(cbm.FeeAmountGWei),
		cbm.Nonce,
		cbm.Signature,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("cancel bet: %w", err)
	}

	b.log(ctx, "cancel bet moderator started", "betID", betID)
	defer b.log(ctx, "cancel bet moderator completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// CancelBetParticipants allows the participants to sign off on cancelling the bet.
func (b *Book) CancelBetParticipants(ctx context.Context, betID string, cbp CancelBetParticipants) (*types.Transaction, *types.Receipt, error) {
	if err := cbp.Validate(); err != nil {
		return nil, nil, fmt.Errorf("validate input: %w", err)
	}

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.CancelBetParticipants(
		tranOpts,
		betID,
		currency.GWei2Wei(cbp.FeeAmountGWei),
		cbp.Nonces,
		cbp.Signatures,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("cancel bet: %w", err)
	}

	b.log(ctx, "cancel bet participants started", "betID", betID)
	defer b.log(ctx, "cancel bet participants completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// CancelBetOwner allows the owner to sign off on cancelling the bet.
func (b *Book) CancelBetOwner(ctx context.Context, betID string, cbo CancelBetOwner) (*types.Transaction, *types.Receipt, error) {
	if err := cbo.Validate(); err != nil {
		return nil, nil, fmt.Errorf("validate input: %w", err)
	}

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.CancelBetOwner(
		tranOpts,
		betID,
		currency.GWei2Wei(cbo.FeeAmountGWei),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("cancel bet: %w", err)
	}

	b.log(ctx, "cancel bet owner started", "betID", betID)
	defer b.log(ctx, "cancel bet owner completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// =============================================================================
// Account Calls

// CancelBetExpired allows any participant to cancel the bet after the bet as
// expired for 30 days and it isn't reconciled.
func (b *Book) CancelBetExpired(ctx context.Context, betID string) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.CancelBetExpired(
		tranOpts,
		betID,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("cancel bet: %w", err)
	}

	b.log(ctx, "cancel bet expired started", "betID", betID)
	defer b.log(ctx, "cancel bet expired completed", "betID", betID)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// Balance will return the balance for the connected account.
func (b *Book) Balance(ctx context.Context) (GWei *big.Float, err error) {
	tranOpts, err := b.ethereum.NewCallOpts(ctx)
	if err != nil {
		return nil, fmt.Errorf("new call opts: %w", err)
	}

	wei, err := b.contract.Balance(tranOpts)
	if err != nil {
		return nil, fmt.Errorf("account balance: %w", err)
	}

	b.log(ctx, "balance", "accountid", b.ethereum.Address().String(), "wei", wei)

	return currency.Wei2GWei(wei), nil
}

// Deposit will add the given amount to the account's contract balance.
func (b *Book) Deposit(ctx context.Context, amountGWei *big.Float) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, amountGWei)
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Deposit(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("deposit: %w", err)
	}

	b.log(ctx, "deposit started", "accountid", b.ethereum.Address().String(), "amountGWei", amountGWei)
	defer b.log(ctx, "deposit completed")

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// Withdraw will move all the account's balance in the contract, to the account's wallet.
func (b *Book) Withdraw(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 1600000, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Withdraw(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("withdraw: %w", err)
	}

	b.log(ctx, "withdraw started", "accountid", b.ethereum.Address().String())
	defer b.log(ctx, "withdraw completed")

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	return tx, receipt, nil
}

// EthereumBalance returns the current balance for the account connecting this book.
func (b *Book) EthereumBalance(ctx context.Context) (wei *big.Int, err error) {
	balance, err := b.ethereum.Balance(ctx)
	if err != nil {
		return nil, fmt.Errorf("current balance: %w", err)
	}

	return balance, nil
}

// =============================================================================

// log will write to the configured log if a traceid exists in the context.
func (b *Book) log(ctx context.Context, msg string, keysAndvalues ...interface{}) {
	if b.logger == nil {
		return
	}

	keysAndvalues = append(keysAndvalues, "traceid", web.GetTraceID(ctx))
	b.logger.Infow(msg, keysAndvalues...)
}
