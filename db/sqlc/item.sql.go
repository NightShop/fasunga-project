// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: item.sql

package db

import (
	"context"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (
    user_email, description, group_key
) VALUES (
    $1, $2, $3
) RETURNING id, user_email, description, group_key, checked
`

type CreateItemParams struct {
	UserEmail   string `json:"user_email"`
	Description string `json:"description"`
	GroupKey    string `json:"group_key"`
}

func (q *Queries) CreateItem(ctx context.Context, arg CreateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, createItem, arg.UserEmail, arg.Description, arg.GroupKey)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.UserEmail,
		&i.Description,
		&i.GroupKey,
		&i.Checked,
	)
	return i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1
`

func (q *Queries) DeleteItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const listItems = `-- name: ListItems :many
SELECT id, user_email, description, group_key, checked FROM items
WHERE group_key = $1
ORDER BY id
`

func (q *Queries) ListItems(ctx context.Context, groupKey string) ([]Item, error) {
	rows, err := q.db.QueryContext(ctx, listItems, groupKey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Item{}
	for rows.Next() {
		var i Item
		if err := rows.Scan(
			&i.ID,
			&i.UserEmail,
			&i.Description,
			&i.GroupKey,
			&i.Checked,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET checked = $2
WHERE id = $1
RETURNING id, user_email, description, group_key, checked
`

type UpdateItemParams struct {
	ID      int64 `json:"id"`
	Checked bool  `json:"checked"`
}

func (q *Queries) UpdateItem(ctx context.Context, arg UpdateItemParams) (Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem, arg.ID, arg.Checked)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.UserEmail,
		&i.Description,
		&i.GroupKey,
		&i.Checked,
	)
	return i, err
}
