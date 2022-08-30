CREATE TABLE items (
    chrt_id integer primary key,
    track_number varchar,
    price serial,
    rid varchar,
    name varchar,
    sale integer,
    size varchar,
    total_price serial,
    nm_id integer,
    brand varchar,
    status integer,

    FOREIGN KEY (track_number) REFERENCES orders (track_number)
);