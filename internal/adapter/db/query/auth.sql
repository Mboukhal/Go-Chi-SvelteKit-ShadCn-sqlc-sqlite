-- Session management

-- name: CreateSession :one
INSERT INTO sessions (user_id, refresh_token, expires_at)
VALUES (?, ?, ?)
RETURNING id, user_id, refresh_token, created_at, expires_at;

-- name: GetSessionByRefreshToken :one
SELECT id, user_id, refresh_token, created_at, expires_at
FROM sessions
WHERE refresh_token = ?
LIMIT 1;

-- name: GetSessionByID :one
SELECT id, user_id, refresh_token, created_at, expires_at
FROM sessions
WHERE id = ?;

-- name: ListSessionsByUserID :many
SELECT id, user_id, refresh_token, created_at, expires_at
FROM sessions
WHERE user_id = ?
ORDER BY created_at DESC;

-- name: DeleteSessionByID :exec
DELETE FROM sessions
WHERE id = ?;

-- name: PurgeExpiredSessions :execrows
DELETE FROM sessions
WHERE expires_at IS NOT NULL AND expires_at <= CURRENT_TIMESTAMP;

-- Revoked token management

-- name: RevokeToken :one
INSERT INTO revoked_tokens (jti, expires_at)
VALUES (?, ?)
ON CONFLICT(jti) DO UPDATE
SET revoked_at = CURRENT_TIMESTAMP,
    expires_at = EXCLUDED.expires_at
RETURNING id, jti, revoked_at, expires_at;

-- name: IsTokenRevoked :one
SELECT EXISTS (
    SELECT 1 FROM revoked_tokens WHERE jti = ?
) AS revoked;

-- name: CleanupExpiredRevokedTokens :execrows
DELETE FROM revoked_tokens
WHERE expires_at IS NOT NULL AND expires_at <= CURRENT_TIMESTAMP;
