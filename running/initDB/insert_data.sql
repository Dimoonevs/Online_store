-- Insert data into table sellers
INSERT INTO sellers (name, phone) VALUES
('Sellers1', '0961865483'),
('Sellers2', '0961865483'),
('Sellers3', '0961865483'),
('Sellers4', '0961865483'),
('Sellers5', '0961865483'),
('Sellers6', '0961865483');

-- Insert data into table products
INSERT INTO products (name, price, sellers_id) VALUES
('Product1', 1000.00, 4),
('Product2', 3000.00, 6),
('Product3', 2000.00, 5),
('Product4', 6000.00, 1),
('Product5', 8000.00, 2),
('Product6', 5000.00, 2),
('Product7', 3000.00, 3),
('Product8', 9000.00, 3),
('Product9', 2000.00, 5),
('Product10', 4000.00, 5);

-- Insert data into table customer
INSERT INTO customer (name, phone) VALUES
('Customer1', '0961865483'),
('Customer2', '0961865483'),
('Customer3', '0961865483'),
('Customer4', '0961865483'),
('Customer5', '0961865483'),
('Customer6', '0961865483');

-- Insert data into table orders
INSERT INTO orders (customer_id) VALUES
(1),
(2),
(3),
(4),
(5),
(6),
(6),
(6),
(1),
(1);

-- Insert data into table customer orders_products
INSERT INTO orders_products (order_id, product_id) VALUES
(1, 3),
(2, 1),
(3, 2),
(1, 3),
(5, 3),
(6, 4),
(2, 6),
(4, 5),
(1, 5),
(1, 1),
(1, 6),
(3, 2),
(7, 2),
(8, 1),
(9, 6),
(10, 4),
(2, 3);
