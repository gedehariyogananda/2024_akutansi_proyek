-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE geographies
(
    id                VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    city_code         VARCHAR(20)              NOT NULL,
    province_code     VARCHAR(20)              NOT NULL,
    district_code     VARCHAR(20)              NOT NULL,
    sub_district_code VARCHAR(20)              NOT NULL,
    province_name     VARCHAR(20)              NOT NULL,
    district_name     VARCHAR(20)              NOT NULL,
    city_name         VARCHAR(20)              NOT NULL,
    sub_district_name VARCHAR(20)              NOT NULL
);

-- migrate:down
DROP TABLE geographies;
