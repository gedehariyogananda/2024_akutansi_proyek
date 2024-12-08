-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE purchase_sellable_products
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    purchase_id         VARCHAR(255)             NOT NULL,
    sellable_product_id VARCHAR(255)             NOT NULL,
    unit_price          DECIMAL(15, 2)           NOT NULL,
    quantity            INT                      NOT NULL,
    company_id          VARCHAR(255)             NOT NULL
);

-- migrate:down
DROP TABLE purchase_sellable_products;
