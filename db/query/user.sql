-- name: CreateUser :one
INSERT INTO users (
    email,
    group_key,
    hashed_password
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET group_key = $2
WHERE email = $1
RETURNING *;


-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;