-- +goose Up
ALTER TABLE transactions
ADD COLUMN user_id varchar(255),
ADD CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES users(id);

-- +goose Down
ALTER TABLE transactions
DROP COLUMN user_id;