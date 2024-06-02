DROP TABLE IF EXISTS orders_products;
DROP TABLE IF EXISTS sellers_products;
DROP TABLE IF EXISTS sellers;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS customer;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS admin;


CREATE TABLE sellers (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         phone VARCHAR(20) NOT NULL
);
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          price NUMERIC(10,2) NOT NULL
);

CREATE TABLE customer(
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(100) NOT NULL,
                         phone VARCHAR(20) NOT NULL
);
CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        customer_id INT REFERENCES customer(id)
);
CREATE TABLE admin (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) NOT NULL,
                       password VARCHAR(100) NOT NULL
);

CREATE TABLE sellers_products(
                                 id SERIAL PRIMARY KEY,
                                 sellers_id INT NOT NULL,
                                 products_id INT NOT NULL,
                                 FOREIGN KEY (sellers_id) REFERENCES sellers(id),
                                 FOREIGN KEY (products_id) REFERENCES products(id)
);
CREATE TABLE orders_products (
                                 id SERIAL PRIMARY KEY,
                                 order_id INT NOT NULL,
                                 product_id INT NOT NULL,
                                 FOREIGN KEY (order_id) REFERENCES orders(id),
                                 FOREIGN KEY (product_id) REFERENCES products(id)
);
INSERT INTO admin (username, password) VALUES ('admin', '$2a$05$7Uq6hscDZ5UfdI4PodVOb.GgPvAr1ayxWpW0jW0W0UG5Y22bKdymm');



