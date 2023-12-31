// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    email,
    group_key,
    hashed_password
) VALUES (
    $1, $2, $3
) RETURNING email, hashed_password, group_key
`

type CreateUserParams struct {
	Email          string `json:"email"`
	GroupKey       string `json:"group_key"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.GroupKey, arg.HashedPassword)
	var i User
	err := row.Scan(&i.Email, &i.HashedPassword, &i.GroupKey)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1
`

func (q *Queries) DeleteUser(ctx context.Context, email string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, email)
	return err
}

const getUser = `-- name: GetUser :one
SELECT email, hashed_password, group_key FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, email)
	var i User
	err := row.Scan(&i.Email, &i.HashedPassword, &i.GroupKey)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET group_key = $2
WHERE email = $1
RETURNING email, hashed_password, group_key
`

type UpdateUserParams struct {
	Email    string `json:"email"`
	GroupKey string `json:"group_key"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser, arg.Email, arg.GroupKey)
	var i User
	err := row.Scan(&i.Email, &i.HashedPassword, &i.GroupKey)
	return i, err
}
