-- +goose Up
CREATE TABLE articles (
    id SERIAL PRIMARY KEY, -- use big serial here
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_articles_user_id ON articles(user_id);

-- +goose Down
DROP TABLE IF EXISTS articles;
