// Package commands provides all the different command options and logic.
package commands

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	scbook "github.com/ardanlabs/bets/business/contract/go/book"
	"github.com/ardanlabs/bets/business/core/book"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

// Deposit will move money from the wallet into the game contract.
func Deposit(ctx context.Context, converter *currency.Converter, book *book.Book, amountUSD float64) error {
	fmt.Println("\nDeposit Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", book.Client().Address())
	fmt.Println("amount          :", amountUSD)

	amountGWei := converter.USD2GWei(big.NewFloat(amountUSD))
	tx, receipt, err := book.Deposit(ctx, amountGWei)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Withdraw will remove money from the game contract back into the wallet.
func Withdraw(ctx context.Context, converter *currency.Converter, book *book.Book) error {
	fmt.Println("\nWithdraw Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("address         :", book.Client().Address())

	tx, receipt, err := book.Withdraw(ctx)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransaction(tx))
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Balance returns the current balance of the specified address.
func Balance(ctx context.Context, converter *currency.Converter, book *book.Book, address string) error {
	fmt.Println("\nGame Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	balance, err := book.AccountBalance(ctx, address)
	if err != nil {
		return err
	}

	fmt.Println("balance         :", currency.GWei2Wei(balance))
	fmt.Println("gwei            :", balance)
	fmt.Println("usd             :", converter.GWei2USD(balance))

	return nil
}

// =============================================================================

// Deploy will deploy the smart contract to the configured network.
func Deploy(ctx context.Context, converter *currency.Converter, clt *ethereum.Client) (err error) {
	startingBalance, err := clt.Balance(ctx)
	if err != nil {
		return err
	}
	defer func() {
		endingBalance, dErr := clt.Balance(ctx)
		if dErr != nil {
			err = dErr
			return
		}
		fmt.Print(converter.FmtBalanceSheet(startingBalance, endingBalance))
	}()

	// =========================================================================

	const gasLimit = 1600000
	const valueGwei = 0.0
	tranOpts, err := clt.NewTransactOpts(ctx, gasLimit, big.NewFloat(valueGwei))
	if err != nil {
		return err
	}

	// =========================================================================

	address, tx, _, err := scbook.DeployBook(tranOpts, clt.Backend)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransaction(tx))

	fmt.Println("\nContract Details")
	fmt.Println("----------------------------------------------------")
	fmt.Println("contract id     :", address.Hex())
	fmt.Printf("export GAME_CONTRACT_ID=%s\n", address.Hex())

	// =========================================================================

	fmt.Println("\nWaiting Logs")
	fmt.Println("----------------------------------------------------")
	log.Root().SetHandler(log.StdoutHandler)

	receipt, err := clt.WaitMined(ctx, tx)
	if err != nil {
		return err
	}
	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}

// Wallet returns the current wallet balance for the specified address.
func Wallet(ctx context.Context, converter *currency.Converter, clt *ethereum.Client, address common.Address) error {
	fmt.Println("\nWallet Balance")
	fmt.Println("----------------------------------------------------")
	fmt.Println("account         :", address)

	wei, err := clt.BalanceAt(ctx, address, nil)
	if err != nil {
		return err
	}

	fmt.Println("wei             :", wei)
	fmt.Println("gwei            :", currency.Wei2GWei(wei))
	fmt.Println("usd             :", converter.Wei2USD(wei))

	return nil
}

// Transaction returns the transaction and receipt information for the specified
// transaction. The txHex is expected to be in a 0x format.
func Transaction(ctx context.Context, converter *currency.Converter, clt *ethereum.Client, tranID string) error {
	fmt.Println("\nTransaction ID")
	fmt.Println("----------------------------------------------------")
	fmt.Println("tran id         :", tranID)

	txHash := common.HexToHash(tranID)
	tx, pending, err := clt.TransactionByHash(ctx, txHash)
	if err != nil {
		return err
	}

	if pending {
		return errors.New("transaction pending")
	}

	fmt.Print(converter.FmtTransaction(tx))

	receipt, err := clt.TransactionReceipt(ctx, txHash)
	if err != nil {
		return err
	}

	fmt.Print(converter.FmtTransactionReceipt(receipt, tx.GasPrice()))

	return nil
}
