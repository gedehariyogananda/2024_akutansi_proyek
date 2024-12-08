-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE stock_opnames
(
    id           VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    title        VARCHAR(255),
    company_id   VARCHAR(255)             NOT NULL,
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL,
    changer_name VARCHAR(255)
);

-- migrate:down
DROP TABLE stock_opnames;
