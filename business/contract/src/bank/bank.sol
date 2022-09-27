// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // Bet represents an individual bet's structure.
    struct Bet {
        bool                      Exists;
        mapping (address => bool) IsParticipant;
        address[]                 Participants;
        address                   Moderator;
        uint256                   Pool;
        uint256                   Expiration;
    }

    // Account represents account information for an account.
    struct Account {
        uint256 Balance;
        uint    Nonce;
    }

    // =========================================================================

    // Owner represents the address who deployed the contract.
    address public Owner;

    // accounts represents the acount information for all participants,
    // moderators, and the Owner.
    mapping (address => Account) private accounts;

    // bets represents current bets, organized by Bet ID.
    mapping (string => Bet) private bets;

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

    // Drain will drain the full value of the contract to the contract owner.
    function Drain() onlyOwner payable public {
        address payable account = payable(msg.sender);
        uint256 bal = address(this).balance;

        account.transfer(bal);
        emit EventLog(string.concat("drain[", Error.Addrtoa(account), "] amount[", Error.Itoa(bal), "]"));
    }

    // PlaceBet will add a bet to the system that is considered a live bet.
    function PlaceBet(
        string    memory   betId,        // Unique Bet identifier
        uint256            amount,       // Amount per participant is betting
        uint256            feeAmount,    // Amount per participant fee
        uint256            expiration,   // Time the bet expires
        address            moderator,    // Address of the moderator
        address[] memory   participants, // List of participant addresses
        uint[]    memory   nonce,        // Nonce used per participant for signing
        bytes[]   calldata signatures    // List of participant signatures
    ) onlyOwner public {

        // Construct a bet from the provided details.
        Bet bet = Bet (
            {
                Exists:       true,
                Participants: participants,
                Moderator:    moderator,
                Expiration:   expiration,
                Pool:         (participants.length * amount)
            }
        );

        // Validate the signatures, balances, nounces.
        for (uint i = 0; i < participants.length; i++) {

            // Reconstruct the data that was signed by this participant.
            bytes32 hashData = keccak256(abi.encode(betId, participants[i], moderator, amount, expiration, nonce[i]));

            // Retrieve the participant's public address from the signature.
            (address participant, Error.Err memory err) = extractAddress(hashData, signatures[i]);
            if (err.isError) {
                revert(err.msg);
            }

            // Ensure the address retrieved from the signature matches the participant.
            if (participant != participants[i]) {
                revert(string.concat("part [", Error.Addrtoa(participants[i]), "] address doesn't match signature"));
            }

            // Ensure the participant has a sufficient balance for the bet.
            if (accounts[participant].Balance < amount + feeAmount) {
                revert(string.concat("part [", Error.Addrtoa(participants[i]), "] has an insufficient balance"));
            }

            // Ensure the nonce is the expected nonce.
            if (accounts[participant].Nonce != nonce) {
                revert(string.concat("part [", Error.Addrtoa(participants[i]), "] invalid nonce [", Error.Itoa(nonce[i]) + "]"));
            }

            // Store the participant's address in the bet's Participants map.
            bet.IsParticipant[participant] = true;
        }

        // Add the bet to the bets maps.
        bets[betId] = bet;

        // Move the funds from the participant's balance into the betting pool.
        for (uint i = 0; i < participants.length; i++) {
            accounts[participants[i]].Balance -= amount + feeAmount;
            accounts[participants[i]].Nonce++;
            accounts[Owner].Balance += feeAmount;
        }

        emit EventLog(string.concat("betId [", betId, "] has been added to the system"));
    }

    // ReconcileSigned allows a moderator to reconcile a bet.
    function Reconcile(
        string    memory   betId,     // Unique Bet identifier
        uint256            feeAmount, // Amount per participant fee
        address[] memory   winners,   // List of winner addresses
        address            moderator, // Address of the moderator 
        uint               nonce,     // Nonce used by moderator for signing
        bytes     calldata signature  // Moderator signature
    ) onlyOwner public {

        // Capture the bet information.
        Bet bet = bets[betId];
        if (!bet.Exists) {
            revert("unknown bet id");
        }

        // Ensure the pool is large enough for the fee.
        uint256 totalFee = feeAmount * bet.Participants.length;
        if (bets[betId].Pool < totalFee) {
            revert("the total fee is larger than the pool");
        }

        // Ensure the bet has passed its expiration.
        if (block.timestamp < bet.Expiration) {
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
        (address validateModerator, Error.Err memory addrErr) = extractAddress(hash, signature);
        if (addrErr.isError) {
            revert(addrErr.msg);
        }

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (moderator != validateModerator) {
            revert("invalid moderator signature");
        }

        // Take the fee from the bet pool.
        Error.Err memory feeErr = takeFee(betId, feeAmount);
        if (feeErr.isError) {
            revert(feeErr.msg);
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
        bytes     calldata signature,
        uint256   feeAmount
    ) onlyOwner public {

        // Take the fee from the bet pool.
        Error.Err memory feeErr = takeFee(betId, feeAmount);
        if (feeErr.isError) {
            revert(feeErr.msg);
        }

        // Ensure the bet has not already been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bets may only be canceled if unreconciled");
        }

        // Retrieve the signer.
        bytes32 hash = hashCancel(betId, betsMap[betId].Participants, nonce);
        (address signer, Error.Err memory err) = extractAddress(hash, signature);
        if (err.isError) {
            revert(err.msg);
        }

        // Ensure the signer is the moderator on file for the bet.
        if (signer != betsMap[betId].Moderator) {
            revert("signer does not have the authority to cancel the bet");
        }

        // Validate the moderator's nonce.
        if (!validNonce(signer, nonce)) {
            revert("invalid nonce for moderator");
        }

        // Ensure the participants match the bet's participants.
        (Error.Err memory partErr) = ensureParticipants(betId, betsMap[betId].Participants);
        if (partErr.isError) {
            revert(partErr.msg);
        }

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
        bytes[]   calldata signatures,
        uint256   feeAmount
    ) onlyOwner public {

        // Take the fee from the bet pool.
        Error.Err memory feeErr = takeFee(betId, feeAmount);
        if (feeErr.isError) {
            revert(feeErr.msg);
        }

        // Ensure the bet has not already been reconciled.
        if (betsMap[betId].Pool == 0) {
            revert("bets may only be canceled if unreconciled");
        }

        // Ensure the participants provided match the bet's participants.
        (Error.Err memory partErr) = ensureParticipants(betId, betsMap[betId].Participants);
        if (partErr.isError) {
            revert(partErr.msg);
        }

        // Ensure all participants have signed to abort the bet.
        if (betsMap[betId].Participants.length != nonce.length) {
            revert("all participants must sign to abort");
        }

        // Ensure all signatories are participants in the bet.
        for (uint i = 0; i < nonce.length; i++) {
            bytes32 hash = hashCancel(betId, betsMap[betId].Participants, nonce[i]);
            (address signer, Error.Err memory err) = extractAddress(hash, signatures[i]);
            if (err.isError){
                revert(err.msg);
            }
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
        Error.Err memory feeErr = takeFee(betId, feeAmount);
        if (feeErr.isError) {
            revert(feeErr.msg);
        }

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
    function ensureParticipants(string memory betId, address[] memory addresses) internal view returns (Error.Err memory) {

        // Ensure the participants provided match the bet's participants.
        if (betsMap[betId].Participants.length != addresses.length) {
            return Error.New("invalid participants");
        }
        for (uint i = 0; i < addresses.length; i++) {
            if (!betsMap[betId].IsParticipant[addresses[i]]) {
                return Error.New("invalid participant");
            }
        }
        return Error.None();
    }

    // takeFee will take the fee from the bet's pool.
    function takeFee(string memory betId, uint256 feeAmount) private returns (Error.Err memory) {

        // Ensure the pool is large enough for the fee.
        if (bets[betId].Pool < feeAmount) {
            accounts[Owner].Balance += bets[betId].Pool;
            bets[betId].Pool = 0;

            // Do not continue transaction, nothing left in pool.
            return Error.New("bet pool too low for fee");
        }

        // Subtract the fee from the pool.
        bets[betId].Pool -= feeAmount;
        accounts[Owner].Balance += feeAmount;

        return Error.None();
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

    // extractAddress expects the raw data that was signed and will apply the Ethereum
    // salt value manually. This hides the underlying implementation of the salt.
    function extractAddress(bytes32 hashData, bytes calldata sig) private pure returns (address, Error.Err memory) {
        if (sig.length != 65) {
            return (address(0), Error.New("invalid signature length"));
        }

        bytes memory prefix = "\x19Ethereum Signed Message:\n32";
        bytes32 saltedData = keccak256(abi.encodePacked(prefix, hashData));

        bytes32 r = bytes32(sig[:32]);
        bytes32 s = bytes32(sig[32:64]);
        uint8 v = uint8(sig[64]);

        return (ecrecover(saltedData, v, r, s), Error.None());
    }
}
