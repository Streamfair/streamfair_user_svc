// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user_svc"."Users" (
 username,
 full_name,
 email,
 password_hash,
 country_code
) VALUES (
 $1, $2, $3, $4, $5
)
RETURNING id, username, full_name, email, country_code, created_at
`

type CreateUserParams struct {
	Username     string `json:"username"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	CountryCode  string `json:"country_code"`
}

type CreateUserRow struct {
	ID          pgtype.Int8        `json:"id"`
	Username    string             `json:"username"`
	FullName    string             `json:"full_name"`
	Email       string             `json:"email"`
	CountryCode string             `json:"country_code"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.PasswordHash,
		arg.CountryCode,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.CountryCode,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT
 username,
 full_name,
 email,
 country_code,
 username_changed_at,
 email_changed_at,
 password_changed_at,
 created_at,
 updated_at
FROM "user_svc"."Users"
WHERE username = $1 LIMIT 1
`

type GetUserByUsernameRow struct {
	Username          string             `json:"username"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	CountryCode       string             `json:"country_code"`
	UsernameChangedAt pgtype.Timestamptz `json:"username_changed_at"`
	EmailChangedAt    pgtype.Timestamptz `json:"email_changed_at"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.CountryCode,
		&i.UsernameChangedAt,
		&i.EmailChangedAt,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
