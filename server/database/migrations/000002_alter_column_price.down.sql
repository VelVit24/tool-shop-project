ALTER TABLE products
ALTER COLUMN price TYPE decimal;
ALTER TABLE order_items
ALTER COLUMN price TYPE decimal;
ALTER TABLE orders
ALTER COLUMN total TYPE decimal;