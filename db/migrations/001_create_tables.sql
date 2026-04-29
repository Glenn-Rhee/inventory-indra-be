-- +goose Up
CREATE TABLE USER (
    id          varchar(255) NOT NULL PRIMARY KEY,
    username    varchar(255) NOT NULL,
    password    varchar(255) NOT NULL,
    img_url     varchar(255),
    created_at  TIMESTAMPTZ DEFAULT NOW()
)

-- +goose Down
DROP TABLE IF EXISTS USER;

-- +goose Up
CREATE TYPE category_type AS ENUM ('MEDICINE', 'ESSENTIALS');

-- +goose Up
CREATE TABLE PRODUCT(
    id                  varchar(255) NOT NULL PRIMARY KEY,
    name                varchar(255) NOT NULL,
    category            category_type NOT NULL,
    price_per_butir     int,
    expired_date        TIMESTAMPTZ NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    updated_at          TIMESTAMPTZ
)

-- +goose Down
DROP TABLE IF EXISTS PRODUCT;

-- +goose Up
CREATE TABLE STOCK (
    id                  varchar(255) NOT NULL PRIMARY KEY,
    stock_per_butir     int NOT NULL,
    last_update         TIMESTAMPTZ,
    product_id          VARCHAR(255),

    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES PRODUCT(id)
        ON DELETE CASCADE
)

-- +goose Down
DROP TABLE IF EXISTS STOCK;

-- +goose UP
CREATE TYPE transaction_type AS ENUM ("IN", "OUT")

-- +goose UP
CREATE TABLE TRANSACTION (
    id                  varchar(255) NOT NULL PRIMARY KEY,
    product_id          varchar(255),
    transaction_type    transaction_type,
    quantity            int,
    created_at          TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES PRODUCT(id)
)