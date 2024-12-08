-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE waiting_lists
(
    id            VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name          VARCHAR(255)             NOT NULL,
    company       VARCHAR(255)             NOT NULL,
    phone         VARCHAR(255)             NOT NULL,
    email         VARCHAR(255)             NOT NULL,
    business_type VARCHAR(255)             NOT NULL,
    created_at    TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE waiting_lists;
