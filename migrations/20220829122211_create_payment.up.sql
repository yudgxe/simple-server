CREATE TABLE payment (
    transaction varchar unique,
    request_id varchar,
    currency varchar,
    provider varchar,
    amount serial,
    payment_dt integer,
    bank varchar,
    delivery_cost serial,
    goods_total integer,
    custom_fee serial,

    FOREIGN KEY (transaction) REFERENCES orders (order_uid)
);