// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	AddAccountTypeToAccount(ctx context.Context, arg AddAccountTypeToAccountParams) error
	CreateAccount(ctx context.Context, arg CreateAccountParams) (CreateAccountRow, error)
	CreateAccountType(ctx context.Context, arg CreateAccountTypeParams) (UserSvcAccountType, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAccountType(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id int64) error
	GetAccountByID(ctx context.Context, id int64) (UserSvcAccount, error)
	GetAccountByOwner(ctx context.Context, owner string) (UserSvcAccount, error)
	GetAccountType(ctx context.Context, id int32) (UserSvcAccountType, error)
	GetAccountTypesForAccount(ctx context.Context, accountsID int64) ([]UserSvcAccountType, error)
	GetAccountsForAccountType(ctx context.Context, accounttypesID int32) ([]UserSvcAccount, error)
	GetUserByID(ctx context.Context, id int64) (GetUserByIDRow, error)
	GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error)
	ListAccountTypes(ctx context.Context, arg ListAccountTypesParams) ([]UserSvcAccountType, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]ListUsersRow, error)
	RemoveAccountTypeFromAccount(ctx context.Context, arg RemoveAccountTypeFromAccountParams) error
	RemoveAllRelationshipsForAccountAccountType(ctx context.Context, accountsID int64) error
	RemoveAllRelationshipsForAccountTypeAccount(ctx context.Context, accounttypesID int32) error
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (UserSvcAccount, error)
	UpdateAccountType(ctx context.Context, arg UpdateAccountTypeParams) (UserSvcAccountType, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
	UpdateUserEmail(ctx context.Context, arg UpdateUserEmailParams) (UpdateUserEmailRow, error)
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) (UpdateUserPasswordRow, error)
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) (UpdateUsernameRow, error)
}

var _ Querier = (*Queries)(nil)
