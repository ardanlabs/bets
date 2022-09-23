// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // betPool represents the amount of money stored in the pool for a given bet.
    mapping (string => uint256) private betPool;

    // betExpires represents the expiration date of a given bet.
    mapping (string => uint256) private betExpires;

    // betModerator represents the moderator for a given bet.
    mapping (string => address) private betModerator;

    // accountNonce represents the current nonce for a bettor or moderator.
    mapping (address => uint) private accountNonce;

    // amountBet represents the amount the account has bet thus far (deprecated)
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

    // PlaceBetsSigned will place bets for all participants.
    function PlaceBetsSigned(
        string    memory betId,
        address   moderator,
        uint256   amount,
        uint256   expiration,
        uint[]    memory nonce,
        uint8[]   memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) onlyOwner public {
        // Initialize bet information.
        betExpires[betId] = expiration;
        betModerator[betId] = moderator;

        // Loop through bettor information and signatures.
        for (uint bettor = 0; bettor < nonce.length; bettor++) {
            // Hash the bet information.
            bytes32 hash = hashPlaceBet(betId, moderator, amount, expiration, nonce[bettor]);

            // Retrieve the bettor's public address from the signed hash and the
            // bettor's signature.
            address bettorAddress = ecrecover(hash, v[bettor], r[bettor], s[bettor]);

            // Ensure the bettor has sufficient balance for the bet.
            if (accountBalances[bettorAddress] < amount) {
                revert("insufficient funds");
            }

            // Ensure the bettor's nonce is valid.
            if (nonce[bettor] != accountNonce[bettorAddress] + 1) {
                revert("invalid bettor nonce");
            }

            // Move the funds from the bettor's balance into the betting pool.
            betPool[betId] += amount;
            accountBalances[bettorAddress] -= amount;

            // Increment account nonce for later bets.
            accountNonce[bettorAddress]++;

            emit EventLog(string.concat("betId[", betId, "] bettor[", Error.Addrtoa(bettorAddress), "] bet[", Error.Itoa(amount), "]"));
        }
    }

    // ReconcileSigned allows a moderator to reconcile a bet.
    function ReconcileSigned(
        string memory betId,
        address[] memory winners,
        uint nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) onlyOwner public {
        // Bet can only be reconciled within 24 hours after expiration.
        if (betExpires[betId] < block.timestamp || block.timestamp > betExpires[betId] + 24 hours) {
            revert("bet cannot be reconciled at this time");
        }

        // Hash the reconciliation information.
        bytes32 hash = hashReconcile(betId, winners, nonce);

        // Retrieve the moderator from the signed hash and signature.
        address moderator = ecrecover(hash, v, r, s);

        // Ensure the address matches the bet's moderator on file.
        if (moderator != betModerator[betId]) {
            revert("not authorized to reconcile this bet");
        }

        // Ensure the moderator's nonce is valid.
        if (nonce != accountNonce[moderator] + 1) {
            revert("invalid moderator nonce");
        }

        // Set winnings amount per winner.
        uint256 winnings = betPool[betId] / winners.length;

        // Reconcile winnings.
        for (uint winner = 0; winner < winners.length; winner++) {
            accountBalances[winners[winner]] += winnings;
            emit EventLog(string.concat("betId[", betId, "] bettor[", Error.Addrtoa(winners[winner]), "] winnings[", Error.Itoa(winnings), "]"));
        }

        // Empty bet pool.
        betPool[betId] = 0;

        // Increment moderator nonce.
        accountNonce[moderator]++;
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

    // =========================================================================

    // hashPlaceBet is an internal function to create a hash for the given bet
    // placement information.
    function hashPlaceBet(string memory betId, address moderator, uint256 amount, uint256 expiration, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, moderator, amount, expiration, nonce)));
    }

    // hashReconcile is an internal function to create a hash for the reconciliation
    // of a given bet.
    function hashReconcile(string memory betId, address[] memory winners, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, winners, nonce)));
    }

    // ethSignedHash is an internal function which signs a hash with the
    // Ethereum prefix.
    function ethSignedHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
