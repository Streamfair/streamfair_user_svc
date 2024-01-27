// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: accounts_accounttypes.sql

package db

import (
	"context"
)

const addAccountTypeToAccount = `-- name: AddAccountTypeToAccount :exec
INSERT INTO "user_svc"."AccountTypes_Accounts" ("Accounts_id", "AccountTypes_id")
VALUES ($1, $2)
`

type AddAccountTypeToAccountParams struct {
	AccountsID     int64 `json:"Accounts_id"`
	AccountTypesID int64 `json:"AccountTypes_id"`
}

func (q *Queries) AddAccountTypeToAccount(ctx context.Context, arg AddAccountTypeToAccountParams) error {
	_, err := q.db.Exec(ctx, addAccountTypeToAccount, arg.AccountsID, arg.AccountTypesID)
	return err
}

const getAccountTypesForAccount = `-- name: GetAccountTypesForAccount :many
SELECT at.id, at.type, at.permissions, at.is_artist, at.is_producer, at.is_writer, at.is_label, at.created_at, at.updated_at FROM "user_svc"."AccountTypes" at
JOIN "user_svc"."AccountTypes_Accounts" aat ON at.id = aat."AccountTypes_id"
WHERE aat."Accounts_id" = $1
`

func (q *Queries) GetAccountTypesForAccount(ctx context.Context, accountsID int64) ([]UserSvcAccountType, error) {
	rows, err := q.db.Query(ctx, getAccountTypesForAccount, accountsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserSvcAccountType{}
	for rows.Next() {
		var i UserSvcAccountType
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Permissions,
			&i.IsArtist,
			&i.IsProducer,
			&i.IsWriter,
			&i.IsLabel,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsForAccountType = `-- name: GetAccountsForAccountType :many
SELECT ac.id, ac.owner, ac.avatar_url, ac.plays, ac.likes, ac.follows, ac.shares, ac.created_at, ac.updated_at FROM "user_svc"."Accounts" ac
JOIN "user_svc"."AccountTypes_Accounts" aat ON ac.id = aat."Accounts_id"
WHERE aat."AccountTypes_id" = $1
`

func (q *Queries) GetAccountsForAccountType(ctx context.Context, accounttypesID int64) ([]UserSvcAccount, error) {
	rows, err := q.db.Query(ctx, getAccountsForAccountType, accounttypesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserSvcAccount{}
	for rows.Next() {
		var i UserSvcAccount
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.AvatarUrl,
			&i.Plays,
			&i.Likes,
			&i.Follows,
			&i.Shares,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeAccountTypeFromAccount = `-- name: RemoveAccountTypeFromAccount :exec
DELETE FROM "user_svc"."AccountTypes_Accounts"
WHERE "Accounts_id" = $1 AND "AccountTypes_id" = $2
`

type RemoveAccountTypeFromAccountParams struct {
	AccountsID     int64 `json:"Accounts_id"`
	AccountTypesID int64 `json:"AccountTypes_id"`
}

func (q *Queries) RemoveAccountTypeFromAccount(ctx context.Context, arg RemoveAccountTypeFromAccountParams) error {
	_, err := q.db.Exec(ctx, removeAccountTypeFromAccount, arg.AccountsID, arg.AccountTypesID)
	return err
}

const removeAllRelationshipsForAccountAccountType = `-- name: RemoveAllRelationshipsForAccountAccountType :exec
DELETE FROM "user_svc"."AccountTypes_Accounts"
WHERE "Accounts_id" = $1
`

func (q *Queries) RemoveAllRelationshipsForAccountAccountType(ctx context.Context, accountsID int64) error {
	_, err := q.db.Exec(ctx, removeAllRelationshipsForAccountAccountType, accountsID)
	return err
}

const removeAllRelationshipsForAccountTypeAccount = `-- name: RemoveAllRelationshipsForAccountTypeAccount :exec
DELETE FROM "user_svc"."AccountTypes_Accounts"
WHERE "AccountTypes_id" = $1
`

func (q *Queries) RemoveAllRelationshipsForAccountTypeAccount(ctx context.Context, accounttypesID int64) error {
	_, err := q.db.Exec(ctx, removeAllRelationshipsForAccountTypeAccount, accounttypesID)
	return err
}
