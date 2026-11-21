CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    token text,
    user_id INTEGER NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL
);