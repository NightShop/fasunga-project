package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}

type CreateUserTxParams struct {
	Email    string `json:"email"`
	GroupKey string `json:"group_key"`
}

type CreateUserTxResult struct {
	User  User  `json:"user"`
	Group Group `json:"group"`
}

// CreateUserTx creates an user and a group it the latter doesn't exist already
func (store *Store) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		_, err = q.GetGroup(ctx, arg.GroupKey)
		if err != nil {
			result.Group.GroupKey, err = q.CreateGroup(ctx, arg.GroupKey)
			if err != nil {
				return err
			}
		}

		result.User, err = q.CreateUser(ctx, CreateUserParams{
			Email:          arg.Email,
			HashedPassword: "testing",
			GroupKey:       arg.GroupKey,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
