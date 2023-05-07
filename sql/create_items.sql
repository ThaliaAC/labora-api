CREATE DATABASE "labora-project-1";
CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    customer_name VARCHAR(80) NOT NULL,
    order_date DATE NOT NULL,
    product VARCHAR(200) NOT NULL,
    quantity INTEGER NOT NULL CHECK(quantity > 0),
    price NUMERIC(10, 2) NOT NULL CHECK(price > 0)
);
