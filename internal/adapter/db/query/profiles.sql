-- Profile management

-- name: CreateProfile :one
INSERT INTO profiles (id, username, phone, email, role, created_at, updated_at)
VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, username, phone, email, role, created_at, updated_at;

-- name: GetProfileByID :one
SELECT id, username, phone, email, role, created_at, updated_at
FROM profiles
WHERE id = ?;

-- name: GetProfileByEmail :one
SELECT id, username, phone, email, role, created_at, updated_at
FROM profiles
WHERE email = ?
LIMIT 1;

-- name: GetProfileByUsername :one
SELECT id, username, phone, email, role, created_at, updated_at
FROM profiles
WHERE username = ?
LIMIT 1;

-- name: GetProfileByPhone :one
SELECT id, username, phone, email, role, created_at, updated_at
FROM profiles
WHERE phone = ?
LIMIT 1;

-- name: ListProfilesByRole :many
SELECT id, username, phone, email, role, created_at, updated_at
FROM profiles
WHERE role = ?
ORDER BY created_at DESC;

-- name: UpdateProfile :one
UPDATE profiles
SET username = COALESCE(?, username),
    phone = COALESCE(?, phone),
    email = COALESCE(?, email),
    role = COALESCE(?, role),
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING id, username, phone, email, role, created_at, updated_at;

-- name: UpdateProfileRole :one
UPDATE profiles
SET role = ?,
    updated_at = CURRENT_TIMESTAMP
WHERE id = ?
RETURNING id, username, phone, email, role, created_at, updated_at;

-- name: DeleteProfileByID :exec
DELETE FROM profiles
WHERE id = ?;
