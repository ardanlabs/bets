-- Description: Create table bets
CREATE TABLE bets (
    bet_id        UUID,
    status        TEXT,
    description   TEXT,
    terms         TEXT,
    amount        INT,
    expiration    TIMESTAMP,
    moderator_id  UUID,
    created       TIMESTAMP,
    updated       TIMESTAMP,

    PRIMARY KEY (bet_id),
    FOREIGN KEY (moderator_id) REFERENCES accounts(account_id) ON DELETE CASCADE
);

-- Description: Create table bets_players
CREATE TABLE bets_players (
    bet_id    UUID,
    account_id UUID,

    PRIMARY KEY (bet_id, account_id),
    FOREIGN KEY (bet_id) REFERENCES bets(bet_id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id) ON DELETE CASCADE
);

-- Description: Create table bets_signatures
CREATE TABLE bets_signatures (
    bet_id      UUID,
    account_id  UUID,
    signature   TEXT,
    nonce       INT,
    date        TIMESTAMP,

    PRIMARY KEY (bet_id, player_id),
    FOREIGN KEY (bet_id) REFERENCES bets(bet_id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES accounts(account_id) ON DELETE CASCADE
);

-- Description: Create table accounts
CREATE TABLE accounts (
  account_id UUID,
  address    TEXT,

  PRIMARY KEY (account_id)
);
