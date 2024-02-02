// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type UserSvcAccount struct {
	ID          int64              `json:"id"`
	AccountType int32              `json:"account_type"`
	Owner       string             `json:"owner"`
	AvatarUri   pgtype.Text        `json:"avatar_uri"`
	Plays       int64              `json:"plays"`
	Likes       int64              `json:"likes"`
	Follows     int64              `json:"follows"`
	Shares      int64              `json:"shares"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type UserSvcAccountType struct {
	ID          int32              `json:"id"`
	Type        string             `json:"type"`
	Permissions []byte             `json:"permissions"`
	IsArtist    bool               `json:"is_artist"`
	IsProducer  bool               `json:"is_producer"`
	IsWriter    bool               `json:"is_writer"`
	IsLabel     bool               `json:"is_label"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type UserSvcAccountTypesAccount struct {
	AccountTypesID int32 `json:"AccountTypes_id"`
	AccountsID     int64 `json:"Accounts_id"`
}

type UserSvcUser struct {
	ID                int64              `json:"id"`
	Username          string             `json:"username"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PasswordHash      string             `json:"password_hash"`
	PasswordSalt      string             `json:"password_salt"`
	CountryCode       string             `json:"country_code"`
	RoleID            pgtype.Int8        `json:"role_id"`
	Status            pgtype.Text        `json:"status"`
	LastLoginAt       pgtype.Timestamptz `json:"last_login_at"`
	UsernameChangedAt pgtype.Timestamptz `json:"username_changed_at"`
	EmailChangedAt    pgtype.Timestamptz `json:"email_changed_at"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
	UpdatedAt         pgtype.Timestamptz `json:"updated_at"`
}
