-- migrate:up

SET TIME ZONE 'Asia/Jakarta';
CREATE TABLE invoices
(
    id             VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE,
    customer_name  VARCHAR(255)             NOT NULL,
    note           VARCHAR(255),
    tax_id         VARCHAR(255)             NOT NULL,
    payment_method VARCHAR(255)             NOT NULL,
    invoice_number VARCHAR(255)             NOT NULL,
    company_id     VARCHAR(255)             NOT NULL,
    status         INT                      NOT NULL,
    tax            INT                      NOT NULL,
    sub_total      DECIMAL(15, 2)           NOT NULL,
    deleted_at     TIMESTAMP WITH TIME ZONE,
    created_at     TIMESTAMP WITH TIME ZONE,
    updated_at     TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE invoices;
