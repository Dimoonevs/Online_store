1. Все продавцы и кол-во товаров у каждого
SELECT s.id AS sellars_id, s.name AS sellers_name, s.phone AS sellers_phone, COUNT(p.id)
FROM sellers s
LEFT JOIN products p ON  p.sellers_id = s.id
GROUP BY s.id, s.name, s.phone;

2. ТОП 5 покупаетелй по кол-ву купленных товаров
SELECT c.id AS customers_id, c.name AS customers_name, c.phone AS customers_phone, COUNT(op.product_id) AS bought_product
FROM customer c
INNER JOIN orders o ON c.id = o.customer_id
INNER JOIN orders_products op ON o.id = op.order_id
GROUP BY c.id, c.name, c.phone
ORDER BY bought_product DESC LIMIT 5;

3. ТОП 5 продавцов по сумме продаж
SELECT s.id AS sellers_id, s.name AS sellers_name, s.phone AS sellers_phone, SUM(p.price) AS total_price
FROM sellers s
INNER JOIN products p ON p.sellers_id = s.id
INNER JOIN orders_products op ON op.product_id = p.id
INNER JOIN orders o ON op.order_id = o.id
GROUP BY s.id, s.name, s.phone
ORDER BY total_price DESC LIMIT 5;

4. Покупатели которые совершили больше 2х заказов
SELECT c.id AS customers_id, c.name AS customers_name, c.phone AS customers_phone, COUNT(o.id) AS order_counts
FROM customer c
INNER JOIN orders o ON o.customer_id = c.id
GROUP BY c.id, c.name, c.phone
HAVING COUNT(o.id) > 2;


