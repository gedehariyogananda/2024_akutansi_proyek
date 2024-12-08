-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE transactions
(
    id                      VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    title                   VARCHAR(50)              NOT NULL,
    name                    VARCHAR(255)             NOT NULL,
    additional_data         JSONB,
    type                    VARCHAR(255)             NOT NULL,
    date                    TIMESTAMP WITH TIME ZONE NOT NULL,
    due_date                TIMESTAMP WITH TIME ZONE,
    note                    VARCHAR(255),
    transaction_record_code VARCHAR(255),
    transaction_id          VARCHAR(100),
    activity_type           VARCHAR(255)             NOT NULL,
    payment_type            VARCHAR(255)             NOT NULL,
    amount                 NUMERIC(20, 2)           NOT NULL
);

-- migrate:down
DROP TABLE transactions;
