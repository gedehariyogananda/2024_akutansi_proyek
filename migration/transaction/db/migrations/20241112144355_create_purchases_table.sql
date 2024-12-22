-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE purchases
(
    id                    VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    total_purchase_amount DECIMAL(15, 2)           NOT NULL,
    company_id            VARCHAR(255)             NOT NULL,
    discount             DECIMAL(15, 2)           NOT NULL,
    tax                   DECIMAL(15, 2)           NOT NULL,
    payment_type        VARCHAR(255)             NOT NULL,
    due_date              TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at            TIMESTAMP WITH TIME ZONE NOT NULL
);

-- migrate:down
DROP TABLE purchases;
