package book_test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"

	scbook "github.com/ardanlabs/bets/business/contract/go/book"
	"github.com/ardanlabs/bets/business/core/book"
	"github.com/ardanlabs/ethereum"
	"github.com/ardanlabs/ethereum/currency"
)

const (
	OwnerAddress    = "0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd"
	OwnerKeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-12T14-47-50.112225000Z--6327a38415c53ffb36c11db55ea74cc9cb4976fd"
	OwnerPassPhrase = "123"

	Player1Address    = "0x0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1KeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-13T16-59-42.277071000Z--0070742ff6003c3e809e78d524f0fe5dcc5ba7f7"
	Player1PassPhrase = "123"

	Player2Address    = "0x8e113078adf6888b7ba84967f299f29aece24c55"
	Player2KeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-05-13T16-57-20.203544000Z--8e113078adf6888b7ba84967f299f29aece24c55"
	Player2PassPhrase = "123"

	ModeratorAddress    = "0x40CFaB8ab694937d644764A3f58237be4c568458"
	ModeratorKeyPath    = "../../../zarf/ethereum/keystore/UTC--2022-09-29T16-18-17.064954000Z--40cfab8ab694937d644764a3f58237be4c568458"
	ModeratorPassPhrase = "123"
)

// These variables provide some static GWei to play with.
var (
	oneUSD    = big.NewFloat(662_833.00)
	tenUSD    = big.NewFloat(0).Mul(oneUSD, big.NewFloat(10))
	twentyUSD = big.NewFloat(0).Mul(oneUSD, big.NewFloat(20))
	fiftyUSD  = big.NewFloat(0).Mul(oneUSD, big.NewFloat(50))
)

// We need a string for the bet id.
var betID = "1234"

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ethereum, err := ethereum.New(ctx, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Adding money to player 1 account")

	// Add money to this account.
	if err := ethereum.SendTransaction(ctx, Player1Address, currency.GWei2Wei(fiftyUSD), 21000); err != nil {
		fmt.Println("Player1Address:", err)
		os.Exit(1)
	}

	fmt.Println("Adding money to player 2 account")

	// Add money to this account.
	if err := ethereum.SendTransaction(ctx, Player2Address, currency.GWei2Wei(fiftyUSD), 21000); err != nil {
		fmt.Println("Player2Address:", err)
		os.Exit(1)
	}

	m.Run()
}

func Test_DepositWithdraw(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect player 1 to the smart contract.
	playerClient, err := book.New(ctx, nil, ethereum.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
	}

	// =========================================================================
	// Deposit process

	// Get the starting balance.
	startingBalance, err := playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// Perform a deposit from the player's wallet.
	depositTx, depositReceipt, err := playerClient.Deposit(ctx, twentyUSD)
	if err != nil {
		t.Fatalf("error making deposit: %s", err)
	}

	// Calculate the expected balance by subtracting the amount deposited and the
	// gas fees for the transaction.
	gasCost := big.NewInt(0).Mul(depositTx.GasPrice(), big.NewInt(0).SetUint64(depositReceipt.GasUsed))
	depositWeiAmount := currency.GWei2Wei(twentyUSD)
	expectedBalance := big.NewInt(0).Sub(startingBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err := playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the deposit minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}

	// =========================================================================
	// Withdraw process

	// Perform a withdraw to the player's wallet.
	withdrawTx, withdrawReceipt, err := playerClient.Withdraw(ctx)
	if err != nil {
		t.Fatalf("error calling withdraw: %s", err)
	}

	// Calculate the expected balance by adding the amount withdrawn and the
	// gas fees for the transaction.
	gasCost = big.NewInt(0).Mul(withdrawTx.GasPrice(), big.NewInt(0).SetUint64(withdrawReceipt.GasUsed))
	expectedBalance = big.NewInt(0).Add(currentBalance, depositWeiAmount)
	expectedBalance.Sub(expectedBalance, gasCost)

	// Get the updated wallet balance.
	currentBalance, err = playerClient.EthereumBalance(ctx)
	if err != nil {
		t.Fatalf("error getting player's wallet balance: %s", err)
	}

	// The player's wallet balance should match the withdrawal minus the fees.
	if expectedBalance.Cmp(currentBalance) != 0 {
		t.Fatalf("expecting final balance to be %d; got %d", expectedBalance, currentBalance)
	}
}

func Test_WithdrawWithoutBalance(t *testing.T) {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect player 1 to the smart contract.
	playerClient, err := book.New(ctx, nil, ethereum.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
	}

	// Perform a withdraw to the player's wallet.
	if _, _, err := playerClient.Withdraw(ctx); err == nil {
		t.Fatal("expecting error when trying to withdraw from an empty balance")
	}
}

func Test_PlaceBet(t *testing.T) {
	placeBet(t)
}

func Test_Reconcile(t *testing.T) {
	bet := placeBet(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// =========================================================================
	// Get pre-reconcile account balances

	// Capture player 1 balance in the smart contract.
	player1Bal, err := bet.player1Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 1: %s", err)
	}

	// Capture player 2 balance in the smart contract.
	player2Bal, err := bet.player2Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 2: %s", err)
	}

	// =========================================================================
	// Generate a signature for the call

	privateKey, err := ethereum.PrivateKeyByKeyFile(ModeratorKeyPath, ModeratorPassPhrase)
	if err != nil {
		t.Fatalf("extract private key moderator: %s", err)
	}
	signature, err := book.Sign(privateKey, betID, ModeratorAddress, 0)
	if err != nil {
		t.Fatalf("signing moderator: %s", err)
	}

	// =========================================================================
	// Reconcile the bet

	rec := book.ReconcileBet{
		Nonce:     big.NewInt(0),
		Moderator: ModeratorAddress,
		Signature: signature,
		Winners:   []string{Player1Address},
	}

	// Reconcile with player 1 as the winner and player 2 as the loser.
	if _, _, err := bet.ownerClient.ReconcileBet(ctx, betID, rec); err != nil {
		t.Fatalf("error calling Reconcile: %s", err)
	}

	// =========================================================================
	// Check balances

	// Capture and check player 1 balance has the correct winnings.
	gotBal, err := bet.player1Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 1: %s", err)
	}
	totalWinnings := big.NewFloat(0).Mul(bet.placeBet.AmountBetGWei, big.NewFloat(2))
	expBal := big.NewFloat(0).Add(player1Bal, totalWinnings)
	if gotBal.Cmp(expBal) != 0 {
		t.Fatalf("wrong player 1 balance, got %v  exp %v", currency.GWei2Wei(gotBal), currency.GWei2Wei(expBal))
	}

	// Capture and check player 2 balance hasn't changed.
	gotBal, err = bet.player2Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 2: %s", err)
	}
	if gotBal.Cmp(player2Bal) != 0 {
		t.Fatalf("wrong player 2 balance, got %v  exp %v", currency.GWei2Wei(gotBal), currency.GWei2Wei(player2Bal))
	}

	// =========================================================================
	// Check bet state

	// Capture the bet details and make sure they have been saved correctly.
	betInfo, err := bet.ownerClient.BetDetails(ctx, betID)
	if err != nil {
		t.Fatalf("error getting bet details: %s", err)
	}

	if betInfo.State != book.StateReconciled {
		t.Errorf("invalid bet state, got %d  exp %d", betInfo.State, book.StateLive)
	}

	exp := big.NewFloat(0)
	if betInfo.AmountBetGWei.Cmp(exp) != 0 {
		t.Errorf("wrong amount, got %v  exp %v", currency.GWei2Wei(betInfo.AmountBetGWei), currency.GWei2Wei(exp))
	}
}

// =============================================================================

func deployContract() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("*** Deploying Contract ***")

	contractID, err := smartContract(ctx)
	if err != nil {
		fmt.Println("error deploying a new contract:", err)
		return "", err
	}

	return contractID, nil
}

func smartContract(ctx context.Context) (string, error) {
	ethereum, err := ethereum.New(ctx, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase)
	if err != nil {
		return "", err
	}

	tranOpts, err := ethereum.NewTransactOpts(ctx, 5_000_000, big.NewFloat(0))
	if err != nil {
		return "", err
	}

	address, tx, _, err := scbook.DeployBook(tranOpts, ethereum.RawClient())
	if err != nil {
		return "", err
	}

	if _, err := ethereum.WaitMined(ctx, tx); err != nil {
		return "", err
	}

	return address.String(), nil
}

// =============================================================================

type bet struct {
	contractID    string
	player1Client *book.Book
	player2Client *book.Book
	ownerClient   *book.Book
	placeBet      book.PlaceBet
}

func placeBet(t *testing.T) bet {
	contractID, err := deployContract()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// =========================================================================
	// Establish books for each of the entities involved

	// Connect player 1 to the smart contract.
	player1Client, err := book.New(ctx, nil, ethereum.NetworkHTTPLocalhost, Player1KeyPath, Player1PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	// Connect player 2 to the smart contract.
	player2Client, err := book.New(ctx, nil, ethereum.NetworkHTTPLocalhost, Player2KeyPath, Player2PassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for player 1: %s", err)
	}

	// Connect owner to the smart contract.
	ownerClient, err := book.New(ctx, nil, ethereum.NetworkHTTPLocalhost, OwnerKeyPath, OwnerPassPhrase, contractID)
	if err != nil {
		t.Fatalf("error creating new book for owner: %s", err)
	}

	// =========================================================================
	// Give player accounts money

	// Deposit ~$20 USD into the player's account.
	if _, _, err := player2Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	// Deposit ~$20 USD into the player's account.
	if _, _, err := player1Client.Deposit(ctx, twentyUSD); err != nil {
		t.Fatalf("error making deposit player 1: %s", err)
	}

	// =========================================================================
	// Place a bet

	// Create the slice of signatures for the two players.
	signatures := make([][]byte, 2)
	privateKey, err := ethereum.PrivateKeyByKeyFile(Player1KeyPath, Player1PassPhrase)
	if err != nil {
		t.Fatalf("extract private key 1: %s", err)
	}
	signatures[0], err = book.Sign(privateKey, betID, Player1Address, 0)
	if err != nil {
		t.Fatalf("signing 1: %s", err)
	}
	privateKey, err = ethereum.PrivateKeyByKeyFile(Player2KeyPath, Player2PassPhrase)
	if err != nil {
		t.Fatalf("extract private key 2: %s", err)
	}
	signatures[1], err = book.Sign(privateKey, betID, Player2Address, 0)
	if err != nil {
		t.Fatalf("signing 2: %s", err)
	}

	// Set the bet amounts and the time to expire in an hour.
	expiration := time.Date(2022, time.September, 1, 1, 1, 1, 0, time.UTC)

	// Construct a PlaceBet to make the PlaceBet call.
	placeBet := book.PlaceBet{
		AmountBetGWei: tenUSD,
		AmountFeeGWei: oneUSD,
		Expiration:    expiration,
		Moderator:     ModeratorAddress,
		Participants:  []string{Player1Address, Player2Address},
		Nonces:        []*big.Int{big.NewInt(0), big.NewInt(0)},
		Signatures:    signatures,
	}

	// Place the bet inside the smart contract.
	if _, _, err := ownerClient.PlaceBet(ctx, betID, placeBet); err != nil {
		t.Fatalf("error calling PlaceBet: %s", err)
	}

	// =========================================================================
	// Check balances

	expPlayerBal := big.NewFloat(0).Sub(twentyUSD, big.NewFloat(0).Add(placeBet.AmountBetGWei, placeBet.AmountFeeGWei))
	expOwnerBal := big.NewFloat(0).Mul(placeBet.AmountFeeGWei, big.NewFloat(2))

	// Capture and check player 1 balance in the smart contract.
	gotBal, err := player1Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 1: %s", err)
	}
	if gotBal.Cmp(expPlayerBal) != 0 {
		t.Fatalf("wrong player 1 balance, got %v  exp %v", currency.GWei2Wei(gotBal), currency.GWei2Wei(expPlayerBal))
	}

	// Capture and check player 2 balance in the smart contract.
	gotBal, err = player2Client.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for player 2: %s", err)
	}
	if gotBal.Cmp(expPlayerBal) != 0 {
		t.Fatalf("wrong player 2 balance, got %v  exp %v", currency.GWei2Wei(gotBal), currency.GWei2Wei(expPlayerBal))
	}

	// Capture and check owner balance in the smart contract.
	gotBal, err = ownerClient.Balance(ctx)
	if err != nil {
		t.Fatalf("error getting balance for owner: %s", err)
	}
	if gotBal.Cmp(expOwnerBal) != 0 {
		t.Fatalf("wrong owner balance, got %v  exp %v", currency.GWei2Wei(gotBal), currency.GWei2Wei(expOwnerBal))
	}

	// =========================================================================
	// Check bet state

	// Capture the bet details and make sure they have been saved correctly.
	betInfo, err := ownerClient.BetDetails(ctx, betID)
	if err != nil {
		t.Fatalf("error getting bet details: %s", err)
	}

	if betInfo.State != book.StateLive {
		t.Errorf("invalid bet state, got %d  exp %d", betInfo.State, book.StateLive)
	}

	if len(betInfo.Participants) != 2 {
		t.Errorf("number of participants wrong, got %d  exp %d", len(betInfo.Participants), 2)
	}

	for i, part := range placeBet.Participants {
		if !strings.EqualFold(part, betInfo.Participants[i]) {
			t.Errorf("wrong participant address, got %s  exp %s", betInfo.Participants[i], part)
		}
	}

	if !strings.EqualFold(betInfo.Moderator, placeBet.Moderator) {
		t.Errorf("wrong moderator address, got %s  exp %s", betInfo.Moderator, placeBet.Moderator)
	}

	if betInfo.AmountBetGWei.Cmp(placeBet.AmountBetGWei) != 0 {
		t.Errorf("wrong amount, got %v  exp %v", currency.GWei2Wei(betInfo.AmountBetGWei), currency.GWei2Wei(placeBet.AmountBetGWei))
	}

	if betInfo.Expiration.UTC() != expiration {
		t.Errorf("wrong expiration, got %v  exp %v", betInfo.Expiration.UTC(), expiration)
	}

	return bet{
		contractID:    contractID,
		player1Client: player1Client,
		player2Client: player2Client,
		ownerClient:   ownerClient,
		placeBet:      placeBet,
	}
}
