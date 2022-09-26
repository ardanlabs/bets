// Package bank represents all the transactions necessary for the game.
package bank

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ardanlabs/bets/business/contract/go/bank"
	"github.com/ardanlabs/bets/foundation/web"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/zap"
)

// Bank represents a bank that allows for the reconciling of a game and
// information about account balances.
type Bank struct {
	logger     *zap.SugaredLogger
	contractID string
	ethereum   *ethereum.Ethereum
	contract   *bank.Bank
}

// New returns a new bank with the ability to manage the game money.
func New(ctx context.Context, logger *zap.SugaredLogger, network string, keyPath string, passPhrase string, contractID string) (*Bank, error) {
	ethereum, err := ethereum.New(ctx, network, keyPath, passPhrase)
	if err != nil {
		return nil, fmt.Errorf("network connect: %w", err)
	}

	contract, err := bank.NewBank(common.HexToAddress(contractID), ethereum.RawClient())
	if err != nil {
		return nil, fmt.Errorf("new contract: %w", err)
	}

	b := Bank{
		logger:     logger,
		contractID: contractID,
		ethereum:   ethereum,
		contract:   contract,
	}

	b.log(ctx, "new bank", "network", network, "contractid", contractID)

	return &b, nil
}

// ContractID returns contract id in use.
func (b *Bank) ContractID() string {
	return b.contractID
}

// Client returns the underlying contract client.
func (b *Bank) Client() *ethereum.Ethereum {
	return b.ethereum
}

// AccountBalance will return the balance for the specified account. Only the
// owner of the smart contract can make this call.
func (b *Bank) AccountBalance(ctx context.Context, accountID string) (BalanceGWei *big.Float, err error) {
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

// Balance will return the balance for the connected account.
func (b *Bank) Balance(ctx context.Context) (GWei *big.Float, err error) {
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

// PlaceBets moves money from the bettors' account balance into the amountLocked.
func (b *Bank) PlaceBets(
	ctx context.Context,
	betID string,
	bettors []string,
	moderator string,
	amount *big.Float,
	expiration time.Time,
	nonce []*big.Int,
	signatures [][]byte,
) (*types.Transaction, *types.Receipt, error) {
	converter := currency.NewDefaultConverter()

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	var bettorAddresses []common.Address
	for _, bettor := range bettors {
		bettorAddresses = append(bettorAddresses, common.HexToAddress(bettor))
	}

	moderatorAddress := common.HexToAddress(moderator)

	betAmount := currency.GWei2Wei(amount)

	expires := new(big.Int).SetInt64(expiration.Unix())

	// TODO: extract v, r, s from signatures
	v := []uint8{0}
	r := [][32]byte{}
	s := [][32]byte{}

	// Set the fee
	feeAmount := converter.USD2Wei(big.NewFloat(1))

	tx, err := b.contract.PlaceBetsSigned(
		tranOpts,
		betID,
		bettorAddresses,
		moderatorAddress,
		betAmount,
		expires,
		nonce,
		v, r, s,
		feeAmount,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("place bet: %w", err)
	}

	b.log(ctx, "place bet started", "accounts", bettors, "amount", betAmount)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log(ctx, "place bet completed")

	return tx, receipt, nil
}

// Reconcile will apply with ante to the winner and loser accounts, plus provide
// the house the game fee.
func (b *Bank) Reconcile(
	ctx context.Context,
	betID string,
	winners []string,
	moderator string,
	nonce *big.Int,
	signature []byte,
) (*types.Transaction, *types.Receipt, error) {
	converter := currency.NewDefaultConverter()

	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	var winnerAddresses []common.Address
	for _, winner := range winners {
		winnerAddresses = append(winnerAddresses, common.HexToAddress(winner))
	}

	moderatorAddress := common.HexToAddress(moderator)

	// TODO: extract v, r, s from signature
	v := uint8(0)
	r := [32]byte{}
	s := [32]byte{}

	// Set the fee
	feeAmount := converter.USD2Wei(big.NewFloat(1))

	tx, err := b.contract.ReconcileSigned(
		tranOpts,
		betID,
		winnerAddresses,
		moderatorAddress,
		nonce,
		v, r, s,
		feeAmount,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("reconcile: %w", err)
	}

	b.log(ctx, "reconcile started", "bet", betID, "winners", winners)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log(ctx, "reconcile completed")

	return tx, receipt, nil
}

// Drain will drain the full value of the smart contract and transfer it to
// another contract/wallet address.
func (b *Bank) Drain(ctx context.Context, target string) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	targetID := common.HexToAddress(target)

	tx, err := b.contract.Drain(tranOpts, targetID)
	if err != nil {
		return nil, nil, fmt.Errorf("drain: %w", err)
	}

	b.log(ctx, "drain started", "target", target)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log(ctx, "drain completed")

	return tx, receipt, nil
}

// Deposit will add the given amount to the account's contract balance.
func (b *Bank) Deposit(ctx context.Context, amountGWei *big.Float) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, amountGWei)
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Deposit(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("deposit: %w", err)
	}

	b.log(ctx, "deposit started", "accountid", b.ethereum.Address().String(), "amountGWei", amountGWei)

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log(ctx, "deposit completed")

	return tx, receipt, nil
}

// Withdraw will move all the account's balance in the contract, to the account's wallet.
func (b *Bank) Withdraw(ctx context.Context) (*types.Transaction, *types.Receipt, error) {
	tranOpts, err := b.ethereum.NewTransactOpts(ctx, 0, big.NewFloat(0))
	if err != nil {
		return nil, nil, fmt.Errorf("new trans opts: %w", err)
	}

	tx, err := b.contract.Withdraw(tranOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("withdraw: %w", err)
	}

	b.log(ctx, "withdraw started", "accountid", b.ethereum.Address().String())

	receipt, err := b.ethereum.WaitMined(ctx, tx)
	if err != nil {
		return nil, nil, fmt.Errorf("wait mined: %w", err)
	}

	b.log(ctx, "withdraw completed")

	return tx, receipt, nil
}

// OwnerBalance returns the current balance for the account used to
// create this bank.
func (b *Bank) OwnerBalance(ctx context.Context) (wei *big.Int, err error) {
	balance, err := b.ethereum.Balance(ctx)
	if err != nil {
		return nil, fmt.Errorf("current balance: %w", err)
	}

	return balance, nil
}

// =============================================================================

// log will write to the configured log if a traceid exists in the context.
func (b *Bank) log(ctx context.Context, msg string, keysAndvalues ...interface{}) {
	if b.logger == nil {
		return
	}

	keysAndvalues = append(keysAndvalues, "traceid", web.GetTraceID(ctx))
	b.logger.Infow(msg, keysAndvalues...)
}
