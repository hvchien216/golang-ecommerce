-- Drop the foreign key constraints on the product_categories table
ALTER TABLE product_categories DROP CONSTRAINT IF EXISTS product_categories_product_id_fkey;
ALTER TABLE product_categories DROP CONSTRAINT IF EXISTS product_categories_category_id_fkey;

-- Drop the product_categories join table
DROP TABLE IF EXISTS product_categories;

-- Drop the categories table
DROP TABLE IF EXISTS categories;

-- Drop the index on the products table
DROP INDEX IF EXISTS products_name_idx;

-- Drop the products table
DROP TABLE IF EXISTS products;

-- Drop the products_status enum type
DROP TYPE IF EXISTS products_status;
