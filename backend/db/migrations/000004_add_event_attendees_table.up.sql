CREATE TABLE event_attendees (
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    event_id UUID NOT NULL REFERENCES events(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    PRIMARY KEY(event_id, user_id)
);
