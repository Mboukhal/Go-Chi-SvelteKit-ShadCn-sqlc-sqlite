-- +goose Up
-- +goose StatementBegin

CREATE TABLE profiles (
    id TEXT PRIMARY KEY DEFAULT (lower(hex(randomblob(16)))) UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE,
    phone VARCHAR(15) UNIQUE,
    email TEXT UNIQUE NOT NULL,
    role TEXT DEFAULT 'USER' NOT NULL CHECK(role IN ('ADMIN', 'USER', 'ORGANIZATION', 'ECOSYSTEM', 'EVALUATOR', 'BANNED', 'UNKNOWN')),
    -- organization association
    organization_id TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- insert default users
INSERT INTO profiles (email, role) VALUES ('lios80466@gmail.com', 'ADMIN');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
