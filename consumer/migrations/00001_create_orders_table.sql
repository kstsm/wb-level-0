-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE orders
(
    id                 SERIAL PRIMARY KEY,
    order_uid          UUID       NOT NULL UNIQUE,
    track_number       TEXT       NOT NULL,
    entry              TEXT       NOT NULL,
    locale             VARCHAR(2) NOT NULL,
    internal_signature TEXT,
    customer_id        TEXT       NOT NULL,
    delivery_service   TEXT       NOT NULL,
    shardkey           TEXT       NOT NULL,
    sm_id              INT        NOT NULL,
    date_created       TIMESTAMP  NOT NULL,
    oof_shard          TEXT       NOT NULL
);

-- +goose Down
DROP TABLE orders;
