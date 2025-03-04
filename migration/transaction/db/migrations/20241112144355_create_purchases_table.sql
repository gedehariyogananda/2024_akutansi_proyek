-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE purchases
(
    id                    VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    purchase_number      VARCHAR(255)             NOT NULL,
    total_purchase_amount DECIMAL(15, 2)           NOT NULL,
    company_id            VARCHAR(255)             NOT NULL,
    discount             DECIMAL(15, 2)           NOT NULL,
    is_discount_percent  BOOLEAN                  NOT NULL,
    tax                   DECIMAL(15, 2)           NOT NULL,
    payment        VARCHAR(255)             NOT NULL,
    note                 VARCHAR(255)             NULL,
    due_date              TIMESTAMP WITH TIME ZONE  NULL,
    created_at            TIMESTAMP WITH TIME ZONE NOT NULL
);

-- migrate:down
DROP TABLE purchases;
