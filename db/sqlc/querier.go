// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (UserSvcSession, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (UserSvcUser, error)
	DeleteUser(ctx context.Context, id int64) error
	GetSession(ctx context.Context, id uuid.UUID) (UserSvcSession, error)
	GetUserByID(ctx context.Context, id int64) (UserSvcUser, error)
	GetUserByUsername(ctx context.Context, username string) (UserSvcUser, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (UpdateUserEmailRow, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (UpdateUserPasswordRow, error)
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (UpdateUsernameRow, error)
}

var _ Querier = (*Queries)(nil)
