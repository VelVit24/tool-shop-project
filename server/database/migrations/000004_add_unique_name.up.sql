alter table products add constraint unique_product_name unique (name);
alter table categories add constraint unique_category_name unique (name);