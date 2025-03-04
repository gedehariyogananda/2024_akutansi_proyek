-- migrate:up
SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE log_activities (
    id VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    user_id VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    device VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);


-- migrate:down
DROP TABLE log_activities;

