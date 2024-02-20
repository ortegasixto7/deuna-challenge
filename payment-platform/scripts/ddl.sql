CREATE TABLE customers (
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(100) NOT NULL,
    national_id VARCHAR(100) NOT NULL,
    name VARCHAR(250) NOT NULL,
    created_at BIGINT NOT NULL
);

CREATE TABLE merchants (
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(100) NOT NULL,
    merchant_id VARCHAR(100) NOT NULL,
    name VARCHAR(250) NOT NULL,
    created_at BIGINT NOT NULL
);

CREATE TABLE payments (
    id SERIAL NOT NULL PRIMARY KEY,
    code VARCHAR(100) NOT NULL,
    customer_id VARCHAR(100) NOT NULL,
    customer_name VARCHAR(250) NOT NULL,
    merchant_id VARCHAR(100) NOT NULL,
    merchant_name VARCHAR(250) NOT NULL,
    amount REAL NOT NULL,
    is_refunded BOOLEAN NOT NULL,
    external_code VARCHAR(100) NOT NULL,
    created_at BIGINT NOT NULL
);