-- name: CreateGroup :one
INSERT INTO groups (
    group_key
) VALUES (
    $1
) RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE group_key = $1;

-- name: GetGroup :one
SELECT * FROM groups
WHERE group_key = $1 LIMIT 1;