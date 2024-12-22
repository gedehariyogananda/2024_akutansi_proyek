-- migrate:up
CREATE TYPE journal_entries_type AS ENUM ('DEBIT', 'CREDIT');

CREATE TABLE journal_entries
(
    id               VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    amount           NUMERIC(20, 2),
    account_id       VARCHAR(50)              NOT NULL,
    company_id       VARCHAR(50)              NOT NULL,
    type             journal_entries_type     NOT NULL,
    additional_data  JSONB,
    note             VARCHAR(255),
    date             TIMESTAMP WITH TIME ZONE NOT NULL,
    transaction_code VARCHAR(255),
    company_id       VARCHAR(50)              NOT NULL
);

-- migrate:down
DROP TABLE journal_entries;
DROP TYPE journal_entries_type;
