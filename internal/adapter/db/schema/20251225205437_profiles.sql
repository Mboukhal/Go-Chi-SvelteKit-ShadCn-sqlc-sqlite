-- +goose Up
-- +goose StatementBegin

CREATE TABLE profiles (
    id TEXT PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    phone VARCHAR(15) UNIQUE,
    email TEXT UNIQUE,
    role TEXT DEFAULT 'USER' NOT NULL CHECK(role IN ('ADMIN', 'USER', 'ORGANIZATION', 'ECOSYSTEM', 'EVALUATOR', 'BANNED', 'UNKNOWN')),
    -- organization association
    organization_id TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- insert default users
INSERT INTO profiles (id, username, phone, email, role) VALUES
(lower(hex(randomblob(16))), NULL, NULL, 'lios80466@gmail.com', 'ADMIN'),
(lower(hex(randomblob(16))), NULL, NULL, 'hami.doc2@gmail.com', 'USER'),
(lower(hex(randomblob(16))), NULL, NULL, 'hmamazeinab@gmail.com', 'ADMIN');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles;
-- +goose StatementEnd
