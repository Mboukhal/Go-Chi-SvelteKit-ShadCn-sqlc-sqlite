-- +goose Up
-- +goose StatementBegin

CREATE TABLE organizations (
    id TEXT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT DEFAULT '',
    -- max users licensed in this company
    holder_account INTEGER NOT NULL DEFAULT 0,
    holder_account_evaluator INTEGER NOT NULL DEFAULT 0,
    -- program type PCP+ or PCP360 or PCP360+
    program_type TEXT NOT NULL,
    logo TEXT DEFAULT '',
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS organizations;
-- +goose StatementEnd
