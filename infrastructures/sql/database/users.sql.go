// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
insert into users (id, username, password, fullname, created_at, updated_at) values(?, ?, ?, ?, current_timestamp, current_timestamp)
`

type CreateUserParams struct {
	ID       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Fullname string `db:"fullname" json:"fullname"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Fullname,
	)
	return err
}

const getByUsername = `-- name: GetByUsername :one
select id, username, password, fullname, created_at, updated_at from users where username = ?
`

func (q *Queries) GetByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Fullname,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
select id, username, password, fullname, created_at, updated_at from users where id = ?
`

func (q *Queries) GetUserByID(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Fullname,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
