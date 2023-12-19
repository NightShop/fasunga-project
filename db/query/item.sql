-- name: CreateItem :one
INSERT INTO items (
    user_email, description, group_key
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;

-- name: ListItems :many
SELECT * FROM items
WHERE group_key = $1
ORDER BY id;

-- name: UpdateItem :one
UPDATE items
SET checked = $2
WHERE id = $1
RETURNING *;