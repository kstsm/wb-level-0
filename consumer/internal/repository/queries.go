package repository

const (
	SaveOrder = `
INSERT INTO orders (order_uid,
                    track_number,
                    entry,
                    locale,
                    internal_signature,
                    customer_id,
                    delivery_service,
                    shardkey,
                    sm_id,
                    date_created,
                    oof_shard)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id;
`
	SaveDelivery = `
INSERT INTO deliveries (order_id, 
                        name,
                        phone,
                        zip,
                        city,
                        address,
                        region,
                        email)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8);
`
	SavePayment = `
INSERT INTO payments (order_id,
                      transaction,
                      request_id,
                      currency,
                      provider,
                      amount,
                      payment_dt,
                      bank,
                      delivery_cost,
                      goods_total,
                      custom_fee)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11);
`
	SaveItem = `
INSERT INTO items (order_id,
                   chrt_id,
                   track_number,
                   price,
                   rid,
                   name,
                   sale,
                   size,
                   total_price,
                   nm_id,
                   brand,
                   status)
VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);
`

	GetOrderByID = `
SELECT json_build_object(
    'id', o.id,
    'order_uid', o.order_uid,
    'track_number', o.track_number,
    'entry', o.entry,
    'locale', o.locale,
    'internal_signature', o.internal_signature,
    'customer_id', o.customer_id,
    'delivery_service', o.delivery_service,
    'shardkey', o.shardkey,
    'sm_id', o.sm_id,
    'date_created', o.date_created AT TIME ZONE 'UTC',
    'oof_shard', o.oof_shard,
    'delivery', (
		SELECT row_to_json(d) 
		FROM deliveries d 
		WHERE d.order_id = o.id),
    'payment', (
		SELECT row_to_json(p) 
		FROM payments p 
		WHERE p.order_id = o.id),
    'items', 
COALESCE((SELECT json_agg(i) 
		FROM items i 
		WHERE i.order_id = o.id), '[]')
) AS order_json
FROM orders o
WHERE o.order_uid = $1
`
	GetAllOrders = `
SELECT json_agg(
    json_build_object(
        'id', o.id,
        'order_uid', o.order_uid,
        'track_number', o.track_number,
        'entry', o.entry,
        'locale', o.locale,
        'internal_signature', o.internal_signature,
        'customer_id', o.customer_id,
        'delivery_service', o.delivery_service,
        'shardkey', o.shardkey,
        'sm_id', o.sm_id,
        'date_created', o.date_created AT TIME ZONE 'UTC',
        'oof_shard', o.oof_shard,
        'delivery', (
            SELECT row_to_json(d) 
            FROM deliveries d 
            WHERE d.order_id = o.id
        ),
        'payment', (
            SELECT row_to_json(p) 
            FROM payments p 
            WHERE p.order_id = o.id
        ),
        'items', COALESCE((
            SELECT json_agg(i) 
            FROM items i 
            WHERE i.order_id = o.id
        ), '[]')
    )
) AS orders_json
FROM orders o;
`
)
