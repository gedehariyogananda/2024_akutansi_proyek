-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE receipt
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    sellable_product_id VARCHAR(255)             NOT NULL,
    material_product_id VARCHAR(255)             NOT NULL,
    quantity            INT                      NOT NULL
);

-- migrate:down
DROP TABLE receipt;
