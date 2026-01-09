-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    price       INTEGER NOT NULL CHECK (price >= 0),
    image_url   VARCHAR(512),
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);