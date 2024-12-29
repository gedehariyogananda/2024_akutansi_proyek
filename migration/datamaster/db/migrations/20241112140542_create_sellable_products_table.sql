-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE sellable_products
(
    id               VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name             VARCHAR(255)             NOT NULL,
    company_id       VARCHAR(255)             NOT NULL,
    smallest_unit_id VARCHAR(255)             NOT NULL,
    category_id      VARCHAR(255)             NOT NULL,
    image            VARCHAR(255)             NOT NULL,
    sku             VARCHAR(255)             NOT NULL,
    description      VARCHAR(255)             NOT NULL,
    status           BOOLEAN                  NOT NULL,
    has_receipt      BOOLEAN                  NOT NULL,
    current_quantity INT                      NOT NULL,
    deleted_at       TIMESTAMP WITH TIME ZONE,
    updated_at       TIMESTAMP WITH TIME ZONE,
    created_at       TIMESTAMP WITH TIME ZONE,
    price            DECIMAL(15, 2)           NOT NULL
);

-- migrate:down
DROP TABLE sellable_products;
