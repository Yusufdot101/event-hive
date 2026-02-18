CREATE TABLE IF NOT EXISTS events(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    starts_at TIMESTAMPTZ NOT NULL,
    ends_at TIMESTAMPTZ NOT NULL,
    last_updated_at TIMESTAMPTZ,
    creator_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    latitude  DOUBLE PRECISION NOT NULL CHECK ( latitude BETWEEN -90 AND 90 ),
    longitude DOUBLE PRECISION NOT NULL CHECK ( longitude BETWEEN -180 AND 180 )
);
