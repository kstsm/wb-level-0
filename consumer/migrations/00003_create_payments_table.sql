-- +goose Up
CREATE TABLE payments
(
    order_id      INT PRIMARY KEY REFERENCES orders (id) ON DELETE CASCADE,
    transaction   TEXT   NOT NULL,
    request_id    TEXT,
    currency      TEXT   NOT NULL,
    provider      TEXT   NOT NULL,
    amount        INT    NOT NULL,
    payment_dt    BIGINT NOT NULL,
    bank          TEXT   NOT NULL,
    delivery_cost INT    NOT NULL,
    goods_total   INT    NOT NULL,
    custom_fee    INT
);

-- +goose Down
DROP TABLE payments;