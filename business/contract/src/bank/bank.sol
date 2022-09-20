// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // amountBet represents the amount of money that has been bet.
    mapping (address => uint256) private amountBet;

    // EventLog provides support for external logging.
    event EventLog(string value);

    // =========================================================================

    // constructor is called when the contract is deployed.
    constructor() {
        Owner = msg.sender;
    }

    // =========================================================================
    // Owner Only Calls

    // onlyOwner can be used to restrict access to a function for only the owner.
    modifier onlyOwner {
        if (msg.sender != Owner) revert();
        _;
    }

    // Drain will drain the full value of the contract and transfer it to another
    // contract/wallet address.
    function Drain(address target) onlyOwner payable public {
        uint balance = address(this).balance;
        address payable account = payable(target);
        account.transfer(balance);
        emit EventLog(string.concat("transfer[", Error.Itoa(balance), "]"));
    }

    // PlaceBet moves money from accountBalances to amountBet.
    function PlaceBet(address person, uint256 amount, uint256 fee) onlyOwner public {

        // Subtract the fee from the account balance immediately.
        if (accountBalances[person] < fee) {
            accountBalances[Owner] += accountBalances[person];
            accountBalances[person] = 0;
        } else {
            accountBalances[Owner] += fee;
            accountBalances[person] -= fee;
        }

        // Check if the balance is enough to accommodate the bet.
        if (accountBalances[person] < amount) {
            revert("account balance too low");
        }

        amountBet[person] = amount;
        emit EventLog(string.concat("bet[", Error.Itoa(amount), "] total[", Error.Itoa(amountBet[person]), "]"));
    }

    // Reconcile settles the accounting for a game that was played.
    function Reconcile(address winner, address loser, uint256 amount, uint256 fee) onlyOwner public {

        // Subtract the fee from the loser's account balance immediately.
        if (accountBalances[loser] < fee) {
            // TODO: Should this revert, return, or should reconcile be allowed
            // to continue, zeroing out the bet for both parties?
            emit EventLog("loser balance too low, taking remainder as fee");
            accountBalances[Owner] += accountBalances[loser];
            accountBalances[loser] = 0;
        } else {
            accountBalances[Owner] += fee;
            accountBalances[loser] -= fee;
        }

        // Reconcile the winnings.
        if (accountBalances[loser] < amount) {
            emit EventLog("loser balance too low, moving remainder to winner account");
            accountBalances[winner] += accountBalances[loser];
            accountBalances[loser] = 0;
        } else {
            accountBalances[winner] += amount;
            accountBalances[loser] -= amount;
        }

        // Remove amount from account bets, be cautious of uint overflow.
        if (amountBet[winner] < amount) {
            amountBet[winner] = 0;
        } else {
            amountBet[winner] -= amount;
        }

        // TODO: If the amount transferred above was less than the bet amount,
        // does the loser still retain some "owed" amount?
        if (amountBet[loser] < amount) {
            amountBet[loser] = 0;
        } else {
            amountBet[loser] -= amount;
        }

        emit EventLog(string.concat("winner balance[", Error.Itoa(accountBalances[winner]), "] loser balance[", Error.Itoa(accountBalances[loser]), "]"));
    }

    // AccountBalance returns the specified account's balance and amount bet.
    function AccountBalance(address account) onlyOwner view public returns (uint[2] memory) {
        return [accountBalances[account], amountBet[account]];
    }

    // =========================================================================
    // Account Only Calls

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accountBalances[msg.sender] - amountBet[msg.sender];
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        accountBalances[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accountBalances[msg.sender]), "]"));
    }

    // Withdraw all of the available balance from the account.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        // IF AMOUNT BET IS > BALANCE REVERT
        // WE DON'T WANT THE UNSIGNED INT TO ROLLOVER
        if (amountBet[msg.sender] > accountBalances[msg.sender]) {
            revert("not enough balance");
        }

        uint bal = accountBalances[msg.sender] - amountBet[msg.sender];
        if (bal == 0) {
            revert("not enough balance");
        }

        account.transfer(bal);
        accountBalances[msg.sender] -= bal;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(bal), "]"));
    }
}
