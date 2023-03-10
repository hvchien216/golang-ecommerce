-- Create the products_status enum type
CREATE TYPE products_status AS ENUM (
    'OUT_OF_STOCK',
    'IN_STOCK',
    'RUNNING_LOW'
    );

-- Create the products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR,
    slug VARCHAR,
    code VARCHAR,
    image VARCHAR,
    images VARCHAR,
    price FLOAT,
    price_min FLOAT,
    price_max FLOAT,
    status products_status,
    is_complete BOOLEAN,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Create the categories table
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    description VARCHAR,
    parent_id INT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- Create the product_categories join table
CREATE TABLE product_categories (
    product_id INT,
    category_id INT,
    PRIMARY KEY (product_id, category_id)
);

-- Create the index on the products table
CREATE INDEX ON products (name);

-- Create the foreign key constraints on the product_categories table
ALTER TABLE product_categories ADD FOREIGN KEY (product_id) REFERENCES products (id);
ALTER TABLE product_categories ADD FOREIGN KEY (category_id) REFERENCES categories (id);
