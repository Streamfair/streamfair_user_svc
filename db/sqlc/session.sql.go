// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: session.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSession = `-- name: CreateSession :one
INSERT INTO "user_svc"."Sessions" (
 id,
 username,
 refresh_token,
 user_agent,
 client_ip,
 is_blocked,
 expires_at
) VALUES (
 $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
`

type CreateSessionParams struct {
	ID           pgtype.UUID        `json:"id"`
	Username     string             `json:"username"`
	RefreshToken string             `json:"refresh_token"`
	UserAgent    string             `json:"user_agent"`
	ClientIp     string             `json:"client_ip"`
	IsBlocked    bool               `json:"is_blocked"`
	ExpiresAt    pgtype.Timestamptz `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (UserSvcSession, error) {
	row := q.db.QueryRow(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i UserSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at FROM "user_svc"."Sessions"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id pgtype.UUID) (UserSvcSession, error) {
	row := q.db.QueryRow(ctx, getSession, id)
	var i UserSvcSession
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}
