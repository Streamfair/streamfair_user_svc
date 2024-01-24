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
	CreateAccountType(ctx context.Context, arg CreateAccountTypeParams) (UserServiceAccountType, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteAccountType(ctx context.Context, id int64) error
	GetAccountByAllParams(ctx context.Context, arg GetAccountByAllParamsParams) (UserServiceAccount, error)
	GetAccountByID(ctx context.Context, id int64) (GetAccountByIDRow, error)
	GetAccountByUsername(ctx context.Context, username string) (GetAccountByUsernameRow, error)
	GetAccountType(ctx context.Context, id int64) (UserServiceAccountType, error)
	GetAccountTypeByAllParams(ctx context.Context, arg GetAccountTypeByAllParamsParams) (UserServiceAccountType, error)
	GetAccountTypeIDsForAccount(ctx context.Context, accountsID int64) ([]GetAccountTypeIDsForAccountRow, error)
	GetAccountTypesForAccount(ctx context.Context, accountsID int64) ([]UserServiceAccountType, error)
	GetAccountsForAccountType(ctx context.Context, accounttypesID int64) ([]GetAccountsForAccountTypeRow, error)
	ListAccountTypes(ctx context.Context, arg ListAccountTypesParams) ([]UserServiceAccountType, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]ListAccountsRow, error)
	RemoveAccountTypeFromAccount(ctx context.Context, arg RemoveAccountTypeFromAccountParams) error
	RemoveAllRelationshipsForAccountAccountType(ctx context.Context, accountsID int64) error
	RemoveAllRelationshipsForAccountTypeAccount(ctx context.Context, accounttypesID int64) error
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (UpdateAccountRow, error)
	UpdateAccountPassword(ctx context.Context, arg UpdateAccountPasswordParams) (UpdateAccountPasswordRow, error)
	UpdateAccountType(ctx context.Context, arg UpdateAccountTypeParams) (UserServiceAccountType, error)
}

var _ Querier = (*Queries)(nil)
