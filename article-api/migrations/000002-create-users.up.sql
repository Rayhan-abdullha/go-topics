-- +migrate Up
ALTER TABLE users
ADD COLUMN phone VARCHAR(20);