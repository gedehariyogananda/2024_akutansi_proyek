-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE journal_entries
(
    id               VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    ammount          NUMERIC(20, 2),
    account_id       VARCHAR(50)              NOT NULL,
    type             VARCHAR(255)             NOT NULL,
    additional_data  JSONB,
    type             VARCHAR(255)             NOT NULL,
    note             VARCHAR(255),
    date             TIMESTAMP WITH TIME ZONE NOT NULL,
    transaction_code VARCHAR(255)
);

-- migrate:down
DROP TABLE journal_entries;
