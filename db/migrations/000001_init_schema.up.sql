CREATE TYPE "products_status" AS ENUM
(
    'out_of_stock',
    'in_stock',
    'running_low'
);

CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar NOT NULL,
    "description" varchar,
    "slug" varchar,
    "code" varchar,
    "image" varchar,
    "images" varchar,
    "price" float,
    "price_min" float,
    "price_max" float,
    "status" products_status,
    "is_complete" boolean,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamp
);

CREATE INDEX ON "products" ("name");

