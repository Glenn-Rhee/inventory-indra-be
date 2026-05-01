-- +goose Up
CREATE TABLE stocks (
    id                  varchar(255) NOT NULL PRIMARY KEY,
    stock_per_butir     int NOT NULL,
    last_update         TIMESTAMPTZ,
    product_id          VARCHAR(255),

    CONSTRAINT fk_product
        FOREIGN KEY (product_id)
        REFERENCES products(id)
        ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS stocks;
