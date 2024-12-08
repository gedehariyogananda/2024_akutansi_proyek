-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE tax
(
    id         VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name       VARCHAR(255)             NOT NULL,
    precentage INT                      NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE tax;
