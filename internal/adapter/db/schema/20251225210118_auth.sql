-- +goose Up
-- +goose StatementBegin
CREATE TABLE revoked_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    jti TEXT UNIQUE,       -- token ID from JWT
    revoked_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME
);
CREATE TABLE sessions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT,
    refresh_token TEXT UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    expires_at DATETIME
);

CREATE TABLE IF NOT EXISTS loginpl (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(32)))) UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS revoked_tokens;
DROP TABLE IF EXISTS sessions;
-- +goose StatementEnd
