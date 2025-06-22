CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR NOT NULL,
    description VARCHAR,
    price NUMERIC(10, 2) NOT NULL CHECK (price >= 0),
    category_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now()
);
