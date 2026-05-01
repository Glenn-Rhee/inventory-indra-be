-- +goose Up
CREATE TABLE users (
    id          varchar(255) NOT NULL PRIMARY KEY,
    username    varchar(255) NOT NULL,
    password    varchar(255) NOT NULL,
    img_url     varchar(255),
    created_at  TIMESTAMPTZ DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS users;

