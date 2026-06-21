-- =====================
-- Categories
-- =====================

INSERT INTO categories (name, slug)
VALUES
('Дрели и шуруповерты', 'dreli-i-shurupoverty'),
('Перфораторы', 'perforatory'),
('Болгарки', 'bolgarki'),
('Пилы', 'pily'),
('Измерительный инструмент', 'izmeritelnyy-instrument'),
('Ручной инструмент', 'ruchnoy-instrument'),
('Расходные материалы', 'rashodnye-materialy');


-- =====================
-- Products
-- =====================

INSERT INTO products
(
    name,
    description,
    price,
    stock,
    image_url,
    id_category,
    slug
)
VALUES

(
    'Дрель Makita DF333DWYE',
    'Аккумуляторная дрель-шуруповерт 12V',
    7990,
    15,
    '/images/makita-df333.jpg',
    (SELECT id FROM categories WHERE slug='dreli-i-shurupoverty'),
    'drill-makita-df333'
),

(
    'Шуруповерт Bosch GSR 120-LI',
    'Компактный аккумуляторный шуруповерт',
    9990,
    20,
    '/images/bosch-gsr120.jpg',
    (SELECT id FROM categories WHERE slug='dreli-i-shurupoverty'),
    'shurupovert-bosch-gsr120'
),

(
    'Перфоратор DeWalt D25133K',
    'Профессиональный перфоратор SDS-plus',
    15990,
    7,
    '/images/dewalt-d25133.jpg',
    (SELECT id FROM categories WHERE slug='perforatory'),
    'perforator-dewalt-d25133'
),

(
    'Перфоратор Makita HR2470',
    'Ударная дрель-перфоратор 780W',
    12990,
    10,
    '/images/makita-hr2470.jpg',
    (SELECT id FROM categories WHERE slug='perforatory'),
    'perforator-makita-hr2470'
),

(
    'Угловая шлифмашина Bosch GWS 750',
    'Болгарка 750W для резки и шлифовки',
    6490,
    12,
    '/images/bosch-gws750.jpg',
    (SELECT id FROM categories WHERE slug='bolgarki'),
    'bolgarka-bosch-gws750'
),

(
    'Циркулярная пила Makita HS7601',
    'Ручная дисковая пила 1200W',
    18990,
    5,
    '/images/makita-hs7601.jpg',
    (SELECT id FROM categories WHERE slug='pily'),
    'pila-makita-hs7601'
),

(
    'Лазерный уровень ADA Cube',
    'Компактный лазерный уровень',
    4990,
    25,
    '/images/ada-cube.jpg',
    (SELECT id FROM categories WHERE slug='izmeritelnyy-instrument'),
    'lazernyy-uroven-ada-cube'
),

(
    'Молоток Stanley',
    'Стальной молоток с удобной ручкой',
    1490,
    50,
    '/images/stanley-hammer.jpg',
    (SELECT id FROM categories WHERE slug='ruchnoy-instrument'),
    'molotok-stanley'
),

(
    'Набор бит Bosch 32 шт',
    'Набор насадок для шуруповерта',
    1990,
    100,
    '/images/bosch-bits.jpg',
    (SELECT id FROM categories WHERE slug='rashodnye-materialy'),
    'nabor-bit-bosch-32'
);