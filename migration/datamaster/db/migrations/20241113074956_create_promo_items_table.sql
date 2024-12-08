-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE promo_items
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    promo_id            VARCHAR(255)             NOT NULL,
    sellable_product_id VARCHAR(255)             NOT NULL
);

-- migrate:down
DROP TABLE promo_items;
