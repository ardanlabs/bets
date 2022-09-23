// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accountBalances represents the amount of money an account has available.
    mapping (address => uint256) private accountBalances;

    // amountLocked represents the amount of money locked in for live bets.
    mapping (string => uint256) private amountLocked;

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
        address[] memory bettors,
        address   moderator,
        uint256   amount,
        uint256   expiration,
        uint[]    memory nonce,
        uint8[]   memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) onlyOwner public {

        // Loop through bettor information and signatures.
        for (uint bettor = 0; bettor < bettors.length; bettor++) {

            // Hash the bet information.
            bytes32 hash = hashPlaceBet(betId, bettors[bettor], moderator, amount, expiration, nonce[bettor]);

            // Retrieve the bettor's public address from the signed hash and the
            // bettor's signature.
            address bettorAddress = ecrecover(hash, v[bettor], r[bettor], s[bettor]);

            // Ensure the address retrieved from the signature matches the bettor.
            if (bettorAddress != bettors[bettor]) {
                revert("invalid bettor");
            }

            // Ensure the bettor has sufficient balance for the bet.
            if (accountBalances[bettorAddress] < amount) {
                revert("insufficient funds");
            }

            // Move the funds from the bettor's balance into the betting pool.
            amountLocked[betId] += amount;
            accountBalances[bettorAddress] -= amount;

            emit EventLog(string.concat("betId[", betId, "] bettor[", Error.Addrtoa(bettorAddress), "] bet[", Error.Itoa(amount), "]"));
        }
    }

    // ReconcileSigned allows a moderator to reconcile a bet.
    function ReconcileSigned(
        string memory betId,
        address[] memory winners,
        address moderator,
        uint nonce,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) onlyOwner public {

        // Hash the reconciliation information.
        bytes32 hash = hashReconcile(betId, winners, moderator, nonce);

        // Retrieve the moderator from the signed hash and signature.
        address validateModerator = ecrecover(hash, v, r, s);

        if (moderator != validateModerator) {
            revert("invalid moderator signature");
        }

        // Set winnings amount per winner.
        uint256 winnings = amountLocked[betId] / winners.length;

        // Reconcile winnings.
        for (uint winner = 0; winner < winners.length; winner++) {
            accountBalances[winners[winner]] += winnings;
            emit EventLog(string.concat("betId[", betId, "] bettor[", Error.Addrtoa(winners[winner]), "] winnings[", Error.Itoa(winnings), "]"));
        }

        // Empty bet pool.
        amountLocked[betId] = 0;
    }

    // AccountBalance returns the specified account's balance and amount bet.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accountBalances[account];
    }

    // =========================================================================
    // Account Only Calls

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accountBalances[msg.sender];
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        accountBalances[msg.sender] += msg.value;
        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accountBalances[msg.sender]), "]"));
    }

    // Withdraw all of the available balance from the account.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        uint bal = accountBalances[msg.sender];
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
    function hashPlaceBet(string memory betId, address bettor, address moderator, uint256 amount, uint256 expiration, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, bettor, moderator, amount, expiration, nonce)));
    }

    // hashReconcile is an internal function to create a hash for the reconciliation
    // of a given bet.
    function hashReconcile(string memory betId, address[] memory winners, address moderator, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, winners, moderator, nonce)));
    }

    // ethSignedHash is an internal function which signs a hash with the
    // Ethereum prefix.
    function ethSignedHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
