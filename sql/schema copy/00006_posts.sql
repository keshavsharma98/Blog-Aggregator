-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    url TEXT UNIQUE NOT NULL,
    description TEXT,
    feed_id UUID NOT NULL REFERENCES feed(id) ON DELETE CASCADE,
    published_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts;
-- +goose StatementEnd
