CREATE TABLE "product_types" (
  "id" bigserial PRIMARY KEY,
  "name" varchar
);

CREATE TABLE "customers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "contact_number" varchar
);

CREATE TABLE "delivery_statuses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL
);

CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "product_type_id" bigint NOT NULL,
  "customer_id" bigint NOT NULL,
  "delivery_status" bigint NOT NULL,
  "delivery_address" varchar NOT NULL,
  "estimated_delivery_date" TIMESTAMPTZ DEFAULT Now() 
);

ALTER TABLE "products" ADD FOREIGN KEY ("product_type_id") REFERENCES "product_types" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");

ALTER TABLE "products" ADD FOREIGN KEY ("delivery_status") REFERENCES "delivery_statuses" ("id");

CREATE INDEX ON "customers" ("name");

CREATE INDEX ON "delivery_statuses" ("name");

CREATE INDEX ON "products" ("product_type_id", "customer_id");


COMMENT ON COLUMN "customers"."name" IS 'customer name required';

COMMENT ON COLUMN "delivery_statuses"."name" IS 'delivery status required';

COMMENT ON COLUMN "products"."delivery_address" IS 'delivery address required';

INSERT INTO delivery_statuses (name) VALUES ('PENDING'), ('ORDERED'), ('SHIPPED'), ('CANCELLED');

INSERT INTO customers(name, contact_number) VALUES ('Cutomer-1', '123'), ('Cutomer-2', '234'), ('Cutomer-3', '345'), ('Cutomer-4', '');

INSERT INTO product_types(name) VALUES ('Product-type-1'), ('Product-type-2'), ('Product-type-3'), ('Product-type-4');

INSERT INTO products(product_type_id, customer_id, delivery_status, delivery_address, estimated_delivery_date) 
VALUES 
(1, 1, 1, 'address-one', now()), 
(2, 2, 2, 'address-two', now()), 
(3, 3, 3, 'address-three', now()), 
(4, 4, 4, 'address-four', now());
