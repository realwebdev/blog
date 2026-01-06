-- +goose Up
-- Add the password column to match the Go struct. 
-- We use VARCHAR(255) for hashed passwords (e.g., bcrypt).
ALTER TABLE users ADD COLUMN password VARCHAR(255) NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE users DROP COLUMN password;
