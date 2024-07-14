CREATE TABLE IF NOT EXISTS order_history (
    client_name VARCHAR(255),
    exchange_name VARCHAR(255),
    label VARCHAR(255),
    pair VARCHAR(255),
    side VARCHAR(255),
    type VARCHAR(255),
    base_qty FLOAT,
    price FLOAT,
    algorithm_name_placed VARCHAR(255),
    lowest_sell_prc FLOAT,
    highest_buy_prc FLOAT,
    commission_quote_qty FLOAT,
    time_placed TIMESTAMP
);

CREATE TABLE IF NOT EXISTS order_book (
    id SERIAL PRIMARY KEY,
    exchange VARCHAR(255),
    pair VARCHAR(255),
    asks JSONB,
    bids JSONB
);