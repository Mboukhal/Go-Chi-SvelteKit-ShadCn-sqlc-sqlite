-- Session management


-- name: GetTokenByEmail :one
SELECT id, counter_request, created_at
FROM loginpl
WHERE email = ?
LIMIT 1;

-- name: UpdateTokenCount :one
UPDATE loginpl SET counter_request = counter_request + 1    
WHERE email = ?
RETURNING id;

-- name: CreateToken :one
INSERT INTO loginpl (email)
VALUES (?)
RETURNING id;

-- name: DeleteTokenByEmail :exec
DELETE FROM loginpl
WHERE email = ?;
