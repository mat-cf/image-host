CREATE TABLE IF NOT EXISTS images (
    id          TEXT PRIMARY KEY,
    filename    TEXT NOT NULL,
    url         TEXT NOT NULL,
    size        BIGINT NOT NULL,
    content_type TEXT NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL
);