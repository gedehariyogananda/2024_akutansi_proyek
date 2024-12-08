-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE sub_users (
    id VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    employee_key VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    company_id VARCHAR(100) NOT NULL
);

-- migrate:down
DROP TABLE sub_users;
