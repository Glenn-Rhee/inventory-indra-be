-- +goose Up
ALTER TABLE products
ADD COLUMN deleted_at TIMESTAMPTZ,
ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT TRUE;

-- +goose Down
ALTER TABLE products
DROP COLUMN deleted_at,
DROP COLUMN is_active;