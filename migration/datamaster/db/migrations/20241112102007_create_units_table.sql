-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE units
(
    id         VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name       VARCHAR(50)              NOT NULL,
    company_id VARCHAR(255)             NOT NULL,
    code       VARCHAR(20)              NOT NULL,
    status     BOOLEAN                  NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE units;
