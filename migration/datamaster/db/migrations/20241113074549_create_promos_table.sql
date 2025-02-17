-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE promos
(
    id         VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name       VARCHAR(255)             NOT NULL,
    type       VARCHAR(50)              NOT NULL,
    start_date TIMESTAMP WITH TIME ZONE,
    end_date   TIMESTAMP WITH TIME ZONE,
    is_all     BOOLEAN                  NOT NULL DEFAULT FALSE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    company_id VARCHAR(255)             NOT NULL,
    amount     INT                      NOT NULL
);

-- migrate:down
DROP TABLE promos;
