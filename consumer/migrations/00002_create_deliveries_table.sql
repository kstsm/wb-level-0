-- +goose Up
CREATE TABLE deliveries
(
    order_id INT PRIMARY KEY REFERENCES orders (id) ON DELETE CASCADE,
    name     TEXT NOT NULL,
    phone    TEXT NOT NULL,
    zip      TEXT NOT NULL,
    city     TEXT NOT NULL,
    address  TEXT NOT NULL,
    region   TEXT NOT NULL,
    email    TEXT NOT NULL
);

-- +goose Down
DROP TABLE deliveries;
