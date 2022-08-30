CREATE TABLE orders (
    order_uid varchar primary key,
    track_number varchar unique,
    entry varchar,
    locale varchar,
    internal_signature varchar,
    customer_id varchar,
    delivery_service varchar,
    shardkey varchar,
    sm_id integer,
    date_created timestamp,
    oof_shard varchar
);