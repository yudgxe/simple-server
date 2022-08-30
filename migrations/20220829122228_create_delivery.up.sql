CREATE TABLE delivery (
    order_uid varchar unique,
    name varchar,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar,

    FOREIGN KEY (order_uid) REFERENCES orders (order_uid)
);