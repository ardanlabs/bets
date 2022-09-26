// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // BetChannel represents an individual bet's structure.
    struct BetChannel {
        mapping (address => bool) IsParticipant;
        address[] Participants;
        address Moderator;
        uint256 Pool;
        uint256 Expiration;
    }

    // Account represents account information for an account.
    struct Account {
        uint256 Balance;
        uint Nonce;
    }

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accounts represents the acount information for all participants,
    // moderators, and the Owner.
    mapping (address => Account) private accounts;

    // betsMap represents current bets, organized by Bet ID.
    mapping (string => BetChannel) private betsMap;

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
    function PlaceBets(
        string    memory betId,
        address[] memory participants,
        address   moderator,
        uint256   amount,
        uint256   expiration,
        uint[]    memory nonce,
        uint8[]   memory v,
        bytes32[] memory r,
        bytes32[] memory s,
        uint256   feeAmount
    ) onlyOwner public {

        // Initialize the new bet's information.
        betsMap[betId].Participants = participants;
        betsMap[betId].Moderator = moderator;
        betsMap[betId].Expiration = expiration;

        // Calculate per-participant fee.
        uint256 fee = feeAmount / participants.length;

        // Loop through participant information and signatures.
        for (uint i = 0; i < participants.length; i++) {

            // Hash the bet information.
            bytes32 hash = hashPlaceBet(betId, participants[i], moderator, amount, expiration, nonce[i]);

            // Retrieve the participant's public address from the signed hash
            // and the participant's signature.
            address partAddress = ecrecover(hash, v[i], r[i], s[i]);

            // Ensure the address retrieved from the signature matches the participant.
            if (partAddress != participants[i]) {
                revert("invalid participant");
            }

            // Ensure the participant has sufficient balance for the bet.
            if (accounts[partAddress].Balance < amount + fee) {
                revert("insufficient funds");
            }

            // Ensure the nonce is valid.
            if (!validNonce(partAddress, nonce[i])) {
                revert("invalid nonce for participant");
            }

            // Store the participant's address in the bet's Participants map.
            betsMap[betId].IsParticipant[partAddress] = true;

            // Move the funds from the participant's balance into the betting pool.
            betsMap[betId].Pool += amount;
            accounts[partAddress].Balance -= amount + fee;
            accounts[Owner].Balance += fee;

            emit EventLog(string.concat("betId[", betId, "] part[", Error.Addrtoa(partAddress), "] bet[", Error.Itoa(amount), "]"));
        }

        incrementNonces(participants);
    }

    // ReconcileSigned allows a moderator to reconcile a bet.
    function Reconcile(
        string memory betId,
        address[] memory winners,
        address moderator,
        uint nonce,
        uint8 v,
        bytes32 r,
        bytes32 s,
        uint256 feeAmount
    ) onlyOwner public {

        // Take the fee from the bet pool.
        takeFee(betId, feeAmount);

        // Ensure the bet has passed its expiration.
        if (block.timestamp < betsMap[betId].Expiration) {
            revert("bet has not yet expired");
        }

        // Ensure the bet has not already been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bet is already reconciled");
        }

        // Ensure the moderator's nonce is valid.
        if (validNonce(moderator, nonce)) {
            revert("moderator nonce is invalid");
        }

        // Hash the reconciliation information.
        bytes32 hash = hashReconcile(betId, winners, moderator, nonce);

        // Retrieve the moderator from the signed hash and signature.
        address validateModerator = ecrecover(hash, v, r, s);

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (moderator != validateModerator) {
            revert("invalid moderator signature");
        }

        // Distribute remaining pool to the winners.
        distributePool(betId, winners);

        // Increment the moderator's nonce.
        accounts[moderator].Nonce++;
    }

    // ModeratorCancel allows the moderator to cancel a bet at any time.
    function ModeratorCancel(
        string    memory betId,
        uint      nonce,
        uint8     v,
        bytes32   r,
        bytes32   s,
        uint256   feeAmount
    ) onlyOwner public {

        // Take the fee from the bet pool.
        takeFee(betId, feeAmount);

        // Ensure the bet has not already been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bets may only be canceled if unreconciled");
        }

        // Retrieve the signer.
        bytes32 hash = hashCancel(betId, betsMap[betId].Participants, nonce);
        address signer = ecrecover(hash, v, r, s);

        // Ensure the signer is the moderator on file for the bet.
        if (signer != betsMap[betId].Moderator) {
            revert("signer does not have the authority to cancel the bet");
        }

        // Validate the moderator's nonce.
        if (!validNonce(signer, nonce)) {
            revert("invalid nonce for moderator");
        }

        // Ensure the participants match the bet's participants.
        ensureParticipants(betId, betsMap[betId].Participants);

        // Perform the refund.
        distributePool(betId, betsMap[betId].Participants);

        // Increment the moderator's nonce.
        accounts[signer].Nonce++;
    }

    // ParticipantCancel allows all participants to sign to cancel a bet before
    // it has expired.
    function ParticipantCancel(
        string    memory betId,
        uint[]    memory nonce,
        uint8[]   memory v,
        bytes32[] memory r,
        bytes32[] memory s,
        uint256   feeAmount
    ) onlyOwner public {

        // Take the fee from the bet pool.
        takeFee(betId, feeAmount);

        // Ensure the bet has not already been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bets may only be canceled if unreconciled");
        }

        // Ensure the participants provided match the bet's participants.
        ensureParticipants(betId, betsMap[betId].Participants);

        // Ensure all participants have signed to abort the bet.
        if (betsMap[betId].Participants.length != nonce.length) {
            revert("all participants must sign to abort");
        }

        // Ensure all signatories are participants in the bet.
        for (uint i = 0; i < nonce.length; i++) {
            bytes32 hash = hashCancel(betId, betsMap[betId].Participants, nonce[i]);
            address signer = ecrecover(hash, v[i], r[i], s[i]);
            if (!betsMap[betId].IsParticipant[signer]) {
                revert("invalid signer");
            }
        }

        // Perform the refund.
        distributePool(betId, betsMap[betId].Participants);

        // Increment nonces for all participants.
        incrementNonces(betsMap[betId].Participants);
    }

    // OwnerCancel allows the owner to cancel a bet at any time.
    function OwnerCancel(string memory betId, uint256 feeAmount) onlyOwner public {

        // Take the fee from the bet pool.
        takeFee(betId, feeAmount);

        // If the pool is zero it's already reconciled or couldn't handle the fee.
        if (betsMap[betId].Pool == 0) {
            revert("bet pool empty");
        }

        // Perform the refund.
        distributePool(betId, betsMap[betId].Participants);
    }

    // AccountBalance returns the specified account's balance and amount bet.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accounts[account].Balance;
    }

    // GetNonce will retrieve the current nonce for a given account.
    function GetNonce(address account) onlyOwner view public returns (uint) {
        return accounts[account].Nonce;
    }

    // =========================================================================
    // Account Only Calls

    // ExpiredCancel allows individual participants to cancel a bet 30 days
    // after cancelation.
    function ExpiredCancel(string memory betId) public {

        // Ensure the bet has expired.
        if (block.timestamp < betsMap[betId].Expiration + 30 days) {
            revert("bets may only be canceled 30+ days after expiration");
        }

        // Ensure the bet has not been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bets may only be canceled if unreconciled");
        }

        // Get the address of the canceler.
        address canceler = msg.sender;

        // Ensure the canceler is one of the participants.
        if (!betsMap[betId].IsParticipant[canceler]) {
            revert("canceler is not authorized to cancel this bet");
        }

        // Refund the pool to all participants.
        distributePool(betId, betsMap[betId].Participants);
    }

    // Balance returns the balance of the caller.
    function Balance() view public returns (uint) {
        return accounts[msg.sender].Balance;
    }

    // Deposit the given amount to the account balance.
    function Deposit() payable public {
        accounts[msg.sender].Balance += msg.value;
        emit EventLog(string.concat("deposit[", Error.Addrtoa(msg.sender), "] balance[", Error.Itoa(accounts[msg.sender].Balance), "]"));
    }

    // Withdraw all of the available balance from the account.
    function Withdraw() payable public {
        address payable account = payable(msg.sender);

        uint bal = accounts[msg.sender].Balance;
        if (bal == 0) {
            revert("not enough balance");
        }

        account.transfer(bal);
        accounts[msg.sender].Balance -= bal;

        emit EventLog(string.concat("withdraw[", Error.Addrtoa(msg.sender), "] amount[", Error.Itoa(bal), "]"));
    }

    // =========================================================================

    // ensureParticipants will ensure the provided addresses are a complete
    // match for a given bet's participants.
    function ensureParticipants(string memory betId, address[] memory addresses) internal view {

        // Ensure the participants provided match the bet's participants.
        if (betsMap[betId].Participants.length != addresses.length) {
            revert("invalid participants");
        }
        for (uint i = 0; i < addresses.length; i++) {
            if (!betsMap[betId].IsParticipant[addresses[i]]) {
                revert("invalid participant");
            }
        }
    }

    // takeFee will take the fee from the bet's pool.
    function takeFee(string memory betId, uint256 feeAmount) internal {

        // Ensure the pool is large enough for the fee.
        if (betsMap[betId].Pool < feeAmount) {
            accounts[Owner].Balance += betsMap[betId].Pool;
            betsMap[betId].Pool = 0;

            // Do not continue transaction, nothing left in pool.
            revert("bet pool too low for fee");
        }

        // Subtract the fee from the pool.
        betsMap[betId].Pool -= feeAmount;
        accounts[Owner].Balance += feeAmount;
    }

    // distributePool will distribute a bet's pool to the provided participants.
    function distributePool(string memory betId, address[] memory participants) internal {

        // Distribute the remaining pool to all participants. If this is a
        // fractional value then it is floored by default. The remainder will
        // later be added to the Owner account.
        uint256 amount = betsMap[betId].Pool / participants.length;
        for (uint i = 0; i < participants.length; i++) {
            accounts[participants[i]].Balance += amount;
            betsMap[betId].Pool -= amount;
            emit EventLog(string.concat("betId[", betId, "] participant[", Error.Addrtoa(participants[i]), "] amount[", Error.Itoa(amount), "]"));
        }

        // If there is a remainder, add it to the Owner's account.
        accounts[Owner].Balance += betsMap[betId].Pool;

        // Clear the bet pool.
        betsMap[betId].Pool = 0;
    }

    // validNonce ensures the provided nonce matches the nonce on file for the
    // given account address.
    function validNonce(address account, uint nonce) internal view returns (bool) {
        return accounts[account].Nonce == nonce;
    }

    // incrementNonces will increment the nonces on file for all provided
    // account addresses.
    function incrementNonces(address[] memory addresses) internal {
        for(uint i = 0; i < addresses.length; i++) {
            accounts[addresses[i]].Nonce++;
        }
    }

    // hashPlaceBet is an internal function to create a hash for the given bet
    // placement information.
    function hashPlaceBet(string memory betId, address participant, address moderator, uint256 amount, uint256 expiration, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, participant, moderator, amount, expiration, nonce)));
    }

    // hashReconcile is an internal function to create a hash for the reconciliation
    // of a given bet.
    function hashReconcile(string memory betId, address[] memory winners, address moderator, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, winners, moderator, nonce)));
    }

    // hashCancel is an internal function to create a hash for the cancelation
    // of a given bet.
    function hashCancel(string memory betId, address[] memory participants, uint nonce) internal pure returns (bytes32) {
        return ethSignedHash(keccak256(abi.encodePacked(betId, participants, nonce)));
    }

    // ethSignedHash is an internal function which signs a hash with the
    // Ethereum prefix.
    function ethSignedHash(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
