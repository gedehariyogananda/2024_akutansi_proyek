-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE stock_opname_items
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    stock_opname_id     VARCHAR(255)             NOT NULL,
    quantity            INT                      NOT NULL,
    stock_id            VARCHAR(255)             NOT NULL,
    product_type        VARCHAR(255)             NOT NULL,
    difference_quantity INT                      NOT NULL,
    initial_quantity    INT                      NOT NULL,
    expired_date        TIMESTAMP WITH TIME ZONE NOT NULL,
    name                VARCHAR(255)             NOT NULL
);

-- migrate:down
DROP TABLE stock_opname_items;
