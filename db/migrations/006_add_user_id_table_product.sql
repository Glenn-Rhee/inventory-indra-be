-- +goose Up
ALTER TABLE products
ADD COLUMN user_id varchar(255),
ADD CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id);

-- +goose Down
DROP TABLE products
DROP COLUMN user_id;