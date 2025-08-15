-- +goose Up
CREATE TABLE items
(
    id           SERIAL PRIMARY KEY,
    order_id     INT REFERENCES orders (id) ON DELETE CASCADE,
    chrt_id      INT      NOT NULL,
    track_number TEXT     NOT NULL,
    price        INT      NOT NULL,
    rid          TEXT     NOT NULL,
    name         TEXT     NOT NULL,
    sale         SMALLINT NOT NULL,
    size         TEXT     NOT NULL,
    total_price  INT      NOT NULL,
    nm_id        INT      NOT NULL,
    brand        TEXT     NOT NULL,
    status       INT      NOT NULL
);

-- +goose Down
DROP TABLE items;