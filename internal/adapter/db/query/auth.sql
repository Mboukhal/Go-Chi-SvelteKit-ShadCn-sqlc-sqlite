-- Session management


-- name: GetTokenByEmail :one
SELECT id, created_at
FROM loginpl
WHERE email = ?
LIMIT 1;


-- name: CreateToken :one
INSERT INTO loginpl (email)
VALUES (?)
RETURNING id;

-- name: DeleteTokenByEmail :exec
DELETE FROM loginpl
WHERE email = ?;
