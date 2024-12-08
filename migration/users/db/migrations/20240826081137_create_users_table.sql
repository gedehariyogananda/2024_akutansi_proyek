-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE users (
    id VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE ,
    username VARCHAR(255) NOT NULL,
    phone VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    company_id VARCHAR(100) NOT NULL
);

-- migrate:down
DROP TABLE users;
