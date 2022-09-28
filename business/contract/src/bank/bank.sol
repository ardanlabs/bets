// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./error.sol";

contract Bank {

    // These constants define the different states a bet can exist in.
    uint8 private constant STATE_NOTEXISTS  = 0;
    uint8 private constant STATE_LIVE       = 1;
    uint8 private constant STATE_RECONCILED = 2;

    // =========================================================================

    // Bet represents an individual bet's structure.
    struct Bet {
        uint8                     State;
        mapping (address => bool) IsParticipant;
        address[]                 Participants;
        address                   Moderator;
        uint256                   Amount;
        uint256                   Expiration;
    }

    // Account represents account information for an account.
    struct Account {
        bool    Exists;
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
    // Owner Called API's

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

    // AccountBalance returns the specified account's balance and amount bet.
    function AccountBalance(address account) onlyOwner view public returns (uint) {
        return accounts[account].Balance;
    }

    // GetNonce will retrieve the current nonce for a given account.
    function GetNonce(address account) onlyOwner view public returns (uint) {
        return accounts[account].Nonce;
    }

    // PlaceBet will add a bet to the system that is considered a live bet.
    function PlaceBet(
        string    memory   betId,        // Unique Bet identifier
        uint256            amount,       // Amount each participant is betting
        uint256            feeAmount,    // Amount each participant pays in upfront fees
        uint256            expiration,   // Time the bet expires
        address            moderator,    // Address of the moderator
        address[] memory   participants, // List of participant addresses
        uint[]    memory   nonce,        // Nonce used per participant for signing
        bytes[]   calldata signatures    // List of participant signatures
    ) onlyOwner public {

        // Construct a bet from the provided details.
        Bet bet = Bet (
            {
                State:        STATE_LIVE,
                Participants: participants,
                Moderator:    moderator,
                Expiration:   expiration,
                Amount:       (amount - feeAmount)
            }
        );

        // The total cost to each participant.
        uint256 totalCost = (amount + feeAmount);

        // Validate the signatures, balances, nounces.
        for (uint i = 0; i < participants.length; i++) {

            // Reconstruct the data that was signed by this participant.
            bytes32 hashData = keccak256(abi.encode(betId, participants[i], nonce[i], moderator, amount, expiration));

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
            if (accounts[participant].Balance < totalCost) {
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
            accounts[participants[i]].Balance -= totalCost;
            accounts[participants[i]].Nonce++;
            accounts[Owner].Balance += feeAmount;
        }

        // Check if we need to add an account for the moderator.
        if (!accounts[moderator].Exists) {
            accounts[moderator] = Account({Exists: true});
        }

        emit EventLog(string.concat("betId [", betId, "] has been added to the system"));
    }

    // ReconcileBet allows a moderator to reconcile a bet.
    function ReconcileBet(
        string    memory   betId,     // Unique Bet identifier
        uint               nonce,     // Nonce used by moderator for signing
        bytes     calldata signature, // Moderator signature
        address[] memory   winners    // List of winner addresses
    ) onlyOwner public {

        // Capture the bet information.
        Bet bet = bets[betId];
        if (!bet.Exists) {
            revert("unknown bet id");
        }

        // Ensure the bet is live.
        if (bet.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure the bet has passed its expiration.
        if (block.timestamp < bet.Expiration) {
            revert("bet has not yet expired");
        }

        // Ensure the nonce used by the moderator is the expected nonce.
        if (accounts[bet.Moderator].Nonce != nonce) {
            revert(string.concat("mod [", Error.Addrtoa(bet.Moderator), "] invalid nonce [", Error.Itoa(nonce) + "]"));
        }

        // Reconstruct the data that was signed by the moderator.
        bytes32 hashData = keccak256(abi.encode(betId, bet.Moderator, nonce));

        // Retrieve the moderator's public address from the signature.
        (address mod, Error.Err memory err) = extractAddress(hashData, signature);
        if (err.isError) {
            revert(err.msg);
        }

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (mod != bet.Moderator) {
            revert("invalid moderator signature");
        }

        // Ensure the winners provided match the bet's participants.
        if (bet.Participants.length != winners.length) {
            return Error.New("invalid number of winners");
        }
        for (uint i = 0; i < winners.length; i++) {
            if (!bet.IsParticipant[winners[i]]) {
                return Error.New("invalid winner");
            }
        }

        // Give each of the winners the amount listed in the bet.
        for (uint i = 0; i < winners.length; i++) {
            accounts[winners[i]].Balance += bet.Amount;
        }

        // Change the state of the bet to reconciled.
        bet.State = STATE_RECONCILED;

        // Increment the moderator's nonce.
        accounts[bet.Moderator].Nonce++;
    }

    // CancelBetModerator allows the moderator to cancel a bet at any time.
    function CancelBetModerator(
        string    memory betId,
        uint256   feeAmount,
        uint      nonce,
        bytes     calldata signature
    ) onlyOwner public {

        // Capture the bet information.
        Bet bet = bets[betId];
        if (!bet.Exists) {
            revert("unknown bet id");
        }

        // Ensure the bet is live.
        if (bet.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure the nonce used by the moderator is the expected nonce.
        if (accounts[bet.Moderator].Nonce != nonce) {
            revert(string.concat("mod [", Error.Addrtoa(bet.Moderator), "] invalid nonce [", Error.Itoa(nonce) + "]"));
        }

        // Reconstruct the data that was signed by the moderator.
        bytes32 hashData = keccak256(abi.encode(betId, bet.Moderator, nonce));

        // Retrieve the moderator's public address from the signature.
        (address mod, Error.Err memory err) = extractAddress(hashData, signature);
        if (err.isError) {
            revert(err.msg);
        }

        // Ensure the moderator on file for the bet is the one that signed to
        // reconcile the bet.
        if (mod != bet.Moderator) {
            revert("invalid moderator signature");
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Amount - feeAmount;
        for (uint i = 0; i < bet.Participants.length; i++) {
            accounts[bet.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += feeAmount;
        }

        // Increment the moderator's nonce.
        accounts[bet.Moderator].Nonce++;
    }

    // CancelBetParticipants allows all the participants to cancel a bet.
    function CancelBetParticipants(
        string    memory betId,
        uint256   feeAmount,
        uint[]    memory nonces,
        bytes[]   calldata signatures
    ) onlyOwner public {

        // Capture the bet information.
        Bet bet = bets[betId];
        if (!bet.Exists) {
            revert("unknown bet id");
        }

        // Ensure the bet is live.
        if (bet.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Ensure we have a proper number of signatures and nonces.
        if ((bet.Participants.length != signatures.length) || (bet.Participants.length != nonces.length)) {
            return Error.New("invalid number of signatures or nonces");
        }

        // Ensure the we have proper signatures from all the participants.
        for (uint i = 0; i < bet.Participants.length; i++) {
            address participant = bet.Participants[i];
            uint    nonce       = nonces[i];
            bytes[] signature   = signatures[i];

            // Ensure the nonce used by the participant is the expected nonce.
            if (accounts[participant].Nonce != nonce) {
                revert(string.concat("mod [", Error.Addrtoa(participant), "] invalid nonce [", Error.Itoa(nonce) + "]"));
            }

            // Reconstruct the data that was signed by the participant.
            bytes32 hashData = keccak256(abi.encode(betId, participant, nonce, bet.Moderator));

            // Retrieve the participant's public address from the signature.
            (address addr, Error.Err memory err) = extractAddress(hashData, signature);
            if (err.isError) {
                revert(err.msg);
            }

            // Ensure the participant's signature matches the address of file.
            if (addr != participant) {
                revert("invalid participant signature");
            }

            // Increment the nonce value for this participant.
            accounts[participant].Nonce++;
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Amount - feeAmount;
        for (uint i = 0; i < bet.Participants.length; i++) {
            accounts[bet.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += feeAmount;
        }
    }

    // CancelBetOwner allows the owner to cancel a bet at any time.
    function CancelBetOwner(
        string  memory betId,
        uint256        feeAmount
    ) onlyOwner public {

        // Capture the bet information.
        Bet bet = bets[betId];
        if (!bet.Exists) {
            revert("unknown bet id");
        }

        // Ensure the bet is live.
        if (bet.State != STATE_LIVE) {
            revert("bet is not live");
        }

        // Return the money back to the participants minus the fee.
        uint256 totalAmount = bet.Amount - feeAmount;
        for (uint i = 0; i < bet.Participants.length; i++) {
            accounts[bet.Participants[i]].Balance += totalAmount;
            accounts[Owner].Balance += feeAmount;
        }
    }

    // =========================================================================
    // Account Called API's

    // ExpiredCancel allows individual participants to cancel a bet 30 days
    // after cancelation.
    function CancelBetExpired(string memory betId) public {

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
    // Private Functions

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
