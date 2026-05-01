-- +goose Up
CREATE TYPE category_type AS ENUM ('MEDICINE', 'ESSENTIALS');

CREATE TABLE products(
    id                  varchar(255) NOT NULL PRIMARY KEY,
    name                varchar(255) NOT NULL,
    category            category_type NOT NULL,
    price_per_butir     int,
    expired_date        TIMESTAMPTZ NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT NOW(),
    updated_at          TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS products;
DROP TYPE IF EXISTS category_type;