-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd