CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "role" text DEFAULT 'user',
  "createdAt" timestamp DEFAULT (now())
);

CREATE TABLE "categories" (
  "id" serial PRIMARY KEY,
  "name" text NOT NULL
);

CREATE TABLE "products" (
  "id" serial PRIMARY KEY,
  "name" text NOT NULL,
  "description" text,
  "price" decimal NOT NULL,
  "stock" int NOT NULL CHECK (stock >= 0),
  "image_url" text,
  "id_category" int
);

CREATE TABLE "cart_items" (
  "id_user" int,
  "id_product" int,
  "amount" int NOT NULL,
  PRIMARY KEY ("id_user", "id_product")
);

CREATE TABLE "orders" (
  "id" serial PRIMARY KEY,
  "id_user" int NOT NULL,
  "status" text DEFAULT 'created',
  "total" decimal,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "order_items" (
  "id_order" int,
  "id_product" int,
  "amount" int NOT NULL,
  "price" decimal NOT NULL,
  PRIMARY KEY ("id_order", "id_product")
);

ALTER TABLE "products" ADD FOREIGN KEY ("id_category") REFERENCES "categories" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cart_items" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "cart_items" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "orders" ADD FOREIGN KEY ("id_user") REFERENCES "users" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("id_order") REFERENCES "orders" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;

ALTER TABLE "order_items" ADD FOREIGN KEY ("id_product") REFERENCES "products" ("id") ON DELETE RESTRICT ON UPDATE NO ACTION DEFERRABLE INITIALLY IMMEDIATE;
