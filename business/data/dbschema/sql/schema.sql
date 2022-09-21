-- Version: 1.1
-- Description: Create status enums.
CREATE TYPE STATUSES AS ENUM ('negotiation', 'signatures', 'moderation', 'live', 'cancel');

-- Version: 1.2
-- Description: Create table accounts
CREATE TABLE accounts (
  address VARCHAR(42),
  nonce   INT,

  PRIMARY KEY (address)
);

-- Version: 1.3
-- Description: Create table bets
CREATE TABLE bets (
    bet_id            UUID,
    status_id         STATUSES,
    description       VARCHAR(240),
    terms             VARCHAR(240),
    amount            INT,
    moderator_address VARCHAR(42)   NULL,
    date_expired      TIMESTAMP,
    date_created      TIMESTAMP,
    date_updated      TIMESTAMP,

    PRIMARY KEY (bet_id),
    FOREIGN KEY (moderator_address) REFERENCES accounts(address) ON DELETE CASCADE
);

-- Version: 1.4
-- Description: Create table bets_players
CREATE TABLE bets_players (
    bet_id  UUID,
    address VARCHAR(42),

    PRIMARY KEY (bet_id, address),
    FOREIGN KEY (bet_id) REFERENCES bets(bet_id) ON DELETE CASCADE,
    FOREIGN KEY (address) REFERENCES accounts(address) ON DELETE CASCADE
);

-- Version: 1.5
-- Description: Create table bets_signatures
CREATE TABLE bets_signatures (
    bet_id      UUID,
    address     VARCHAR(42),
    nonce       INT,
    signature   VARCHAR(66),
    date_signed TIMESTAMP,

    PRIMARY KEY (bet_id, address, nonce),
    FOREIGN KEY (bet_id) REFERENCES bets(bet_id) ON DELETE CASCADE,
    FOREIGN KEY (address) REFERENCES accounts(address) ON DELETE CASCADE
);
