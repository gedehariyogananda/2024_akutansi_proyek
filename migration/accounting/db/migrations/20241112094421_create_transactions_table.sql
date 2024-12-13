-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE transactions
(
    id                      VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    title                   VARCHAR(50)              NOT NULL,
    name                    VARCHAR(255)             NOT NULL,
    additional_data         JSONB,
    date                    TIMESTAMP WITH TIME ZONE NOT NULL,
    due_date                TIMESTAMP WITH TIME ZONE,
    note                    VARCHAR(255),
    transaction_record_code VARCHAR(255), -- 1_PAT_02042004
    transaction_id          VARCHAR(100),
    payment_method           VARCHAR(255)             NOT NULL,
    payment_type            VARCHAR(255)             NOT NULL,
    amount                 NUMERIC(20, 2)           NOT NULL,
    company_id          VARCHAR(100) NOT NULL
);

-- migrate:down
DROP TABLE transactions;
