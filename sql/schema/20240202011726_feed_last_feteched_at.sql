-- +goose Up
-- +goose StatementBegin
ALTER TABLE feed ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE feed DROP COLUMN last_fetched_at;
-- +goose StatementEnd
