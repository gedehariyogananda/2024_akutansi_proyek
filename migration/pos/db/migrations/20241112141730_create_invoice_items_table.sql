-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE invoice_items
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    invoice_id          VARCHAR(255)             NOT NULL,
    sellable_product_id VARCHAR(255)             NOT NULL,
    quantity            INT                      NOT NULL,
    company_id          VARCHAR(255)             NOT NULL,
    price               DECIMAL(15, 2)           NOT NULL,
    promo_id            VARCHAR(255),
    promo_amount        INT DEFAULT 0,
);

-- migrate:down
DROP TABLE invoice_items;
