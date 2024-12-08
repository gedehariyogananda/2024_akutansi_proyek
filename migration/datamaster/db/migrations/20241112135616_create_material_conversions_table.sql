-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE material_conversions
(
    id                  VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    material_product_id VARCHAR(255)             NOT NULL,
    name                VARCHAR(255)             NOT NULL,
    quantity            INT                      NOT NULL,
    deleted_at          TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE material_conversions;
