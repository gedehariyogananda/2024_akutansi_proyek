-- migrate:up

INSERT INTO users (id,username, phone, name, email, password, created_at, updated_at, company_id) VALUES
('048ba575-618d-5a13-9c9a-d483917e41c0','ari','083133737660', 'ari', 'ari@gmail.com', '$2a$10$PI9hV1gmxXWUVRDHwQSrT.VOkggfLRTGvYVduPDVXd8upkoqqyfkC', now(), now(), 'company-id');

-- migrate:down
