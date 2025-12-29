-- +goose Up
-- +goose StatementBegin
-- Triggers not supported by sqlc for SQLite, handle updated_at in application code
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- No triggers to drop
-- +goose StatementEnd
