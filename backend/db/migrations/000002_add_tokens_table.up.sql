CREATE TABLE IF NOT EXISTS tokens (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ NOT NULL,
    token_use TEXT NOT NULL CHECK ( token_use IN ('refresh') ),
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    token_string UUID
);
