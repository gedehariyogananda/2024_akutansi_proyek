-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE sellable_stocks
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    sellable_product_id VARCHAR(255)             NOT NULL,
    product_type        VARCHAR(255)             NOT NULL,
    quantity            INT                      NOT NULL,
    current_quantity    INT                      NOT NULL,
    expired_date        TIMESTAMP WITH TIME ZONE NOT NULL,
    company_id          VARCHAR(255)             NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE sellable_stocks;
