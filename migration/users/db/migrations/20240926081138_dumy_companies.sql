-- migrate:up

INSERT INTO companies (id,code,name, address, image,created_at, updated_at, geography_id) VALUES
('073cebf0-4c4a-5d4f-9950-5ab3feaf1894','PAT-3239','PT Ari Ganteng', 'Jl. Merdeka No. 1, Jakarta', '/company-file/1726593321.png', now(), now(), 'geography_id');

-- migrate:down
