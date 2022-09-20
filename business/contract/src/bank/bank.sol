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
    }

    // AccountBalance returns the current account's balance.
    function AccountBalance(address account) onlyOwner view public returns (uint[2] memory) {
        return [accountBalances[account], amountBet[msg.sender]];
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

        uint bal = accountBalances[msg.sender] - amountBet[msg.sender];
        if (bal == 0) {
            revert("not enough balance");
        }

        account.transfer(bal);        
        accountBalances[msg.sender] -= bal;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(bal), "]"));
    }
}