-- +goose Up
CREATE TYPE transaction_type AS ENUM (
    'IN',
    'OUT'
);

CREATE TABLE transactions (
    id                  varchar(255) NOT NULL PRIMARY KEY,
    product_id          varchar(255),
    transaction_type    transaction_type,
    quantity            int,
    created_at          TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_product
        FOREIGN KEY(product_id)
        REFERENCES products(id)
);

-- +goose Down
DROP TABLE IF EXISTS transactions;
DROP TYPE IF EXISTS transaction_type;