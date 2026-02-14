CREATE TABLE IF NOT EXISTS users (
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    last_updated_at TIMESTAMPTZ,
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL CHECK ( LENGTH(name) >= 2 ),
    email CITEXT UNIQUE NOT NULL,
    password_hash BYTEA
);
