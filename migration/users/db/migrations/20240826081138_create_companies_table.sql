-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE companies (
    id VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255),
    image VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    geography_id VARCHAR(100)
);

-- migrate:down
DROP TABLE companies;
