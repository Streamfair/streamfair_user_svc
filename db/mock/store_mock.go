// Code generated by MockGen. DO NOT EDIT.
// Source: db/sqlc/store.go
//
// Generated by this command:
//
//	mockgen -source=db/sqlc/store.go -destination=db/mock/store_mock.go
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"
	time "time"

	db "github.com/Streamfair/streamfair_user_svc/db/sqlc"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddAccountTypeToAccount mocks base method.
func (m *MockStore) AddAccountTypeToAccount(ctx context.Context, arg db.AddAccountTypeToAccountParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAccountTypeToAccount", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAccountTypeToAccount indicates an expected call of AddAccountTypeToAccount.
func (mr *MockStoreMockRecorder) AddAccountTypeToAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAccountTypeToAccount", reflect.TypeOf((*MockStore)(nil).AddAccountTypeToAccount), ctx, arg)
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(ctx context.Context, arg db.CreateAccountParams) (db.CreateAccountRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", ctx, arg)
	ret0, _ := ret[0].(db.CreateAccountRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), ctx, arg)
}

// CreateAccountTx mocks base method.
func (m *MockStore) CreateAccountTx(ctx context.Context, params db.CreateAccountTxParams) (db.CreateAccountTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountTx", ctx, params)
	ret0, _ := ret[0].(db.CreateAccountTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccountTx indicates an expected call of CreateAccountTx.
func (mr *MockStoreMockRecorder) CreateAccountTx(ctx, params any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountTx", reflect.TypeOf((*MockStore)(nil).CreateAccountTx), ctx, params)
}

// CreateAccountType mocks base method.
func (m *MockStore) CreateAccountType(ctx context.Context, arg db.CreateAccountTypeParams) (db.UserSvcAccountType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccountType", ctx, arg)
	ret0, _ := ret[0].(db.UserSvcAccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccountType indicates an expected call of CreateAccountType.
func (mr *MockStoreMockRecorder) CreateAccountType(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccountType", reflect.TypeOf((*MockStore)(nil).CreateAccountType), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.CreateUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(db.CreateUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, arg)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), ctx, id)
}

// DeleteAccountTx mocks base method.
func (m *MockStore) DeleteAccountTx(ctx context.Context, accountID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccountTx", ctx, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccountTx indicates an expected call of DeleteAccountTx.
func (mr *MockStoreMockRecorder) DeleteAccountTx(ctx, accountID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccountTx", reflect.TypeOf((*MockStore)(nil).DeleteAccountTx), ctx, accountID)
}

// DeleteAccountType mocks base method.
func (m *MockStore) DeleteAccountType(ctx context.Context, id int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccountType", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccountType indicates an expected call of DeleteAccountType.
func (mr *MockStoreMockRecorder) DeleteAccountType(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccountType", reflect.TypeOf((*MockStore)(nil).DeleteAccountType), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), ctx, id)
}

// GetAccountByID mocks base method.
func (m *MockStore) GetAccountByID(ctx context.Context, id int64) (db.UserSvcAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", ctx, id)
	ret0, _ := ret[0].(db.UserSvcAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockStoreMockRecorder) GetAccountByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockStore)(nil).GetAccountByID), ctx, id)
}

// GetAccountByOwner mocks base method.
func (m *MockStore) GetAccountByOwner(ctx context.Context, owner string) (db.UserSvcAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByOwner", ctx, owner)
	ret0, _ := ret[0].(db.UserSvcAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByOwner indicates an expected call of GetAccountByOwner.
func (mr *MockStoreMockRecorder) GetAccountByOwner(ctx, owner any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByOwner", reflect.TypeOf((*MockStore)(nil).GetAccountByOwner), ctx, owner)
}

// GetAccountType mocks base method.
func (m *MockStore) GetAccountType(ctx context.Context, id int32) (db.UserSvcAccountType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountType", ctx, id)
	ret0, _ := ret[0].(db.UserSvcAccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountType indicates an expected call of GetAccountType.
func (mr *MockStoreMockRecorder) GetAccountType(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountType", reflect.TypeOf((*MockStore)(nil).GetAccountType), ctx, id)
}

// GetAccountTypesForAccount mocks base method.
func (m *MockStore) GetAccountTypesForAccount(ctx context.Context, accountsID int64) ([]db.UserSvcAccountType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountTypesForAccount", ctx, accountsID)
	ret0, _ := ret[0].([]db.UserSvcAccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountTypesForAccount indicates an expected call of GetAccountTypesForAccount.
func (mr *MockStoreMockRecorder) GetAccountTypesForAccount(ctx, accountsID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountTypesForAccount", reflect.TypeOf((*MockStore)(nil).GetAccountTypesForAccount), ctx, accountsID)
}

// GetAccountsForAccountType mocks base method.
func (m *MockStore) GetAccountsForAccountType(ctx context.Context, accounttypesID int32) ([]db.UserSvcAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountsForAccountType", ctx, accounttypesID)
	ret0, _ := ret[0].([]db.UserSvcAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountsForAccountType indicates an expected call of GetAccountsForAccountType.
func (mr *MockStoreMockRecorder) GetAccountsForAccountType(ctx, accounttypesID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountsForAccountType", reflect.TypeOf((*MockStore)(nil).GetAccountsForAccountType), ctx, accounttypesID)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(ctx context.Context, id int64) (db.GetUserByIDRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(db.GetUserByIDRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), ctx, id)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(ctx context.Context, username string) (db.GetUserByUsernameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", ctx, username)
	ret0, _ := ret[0].(db.GetUserByUsernameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(ctx, username any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), ctx, username)
}

// ListAccountTypes mocks base method.
func (m *MockStore) ListAccountTypes(ctx context.Context, arg db.ListAccountTypesParams) ([]db.UserSvcAccountType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccountTypes", ctx, arg)
	ret0, _ := ret[0].([]db.UserSvcAccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccountTypes indicates an expected call of ListAccountTypes.
func (mr *MockStoreMockRecorder) ListAccountTypes(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccountTypes", reflect.TypeOf((*MockStore)(nil).ListAccountTypes), ctx, arg)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(ctx context.Context, arg db.ListAccountsParams) ([]db.ListAccountsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", ctx, arg)
	ret0, _ := ret[0].([]db.ListAccountsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), ctx, arg)
}

// ListUsers mocks base method.
func (m *MockStore) ListUsers(ctx context.Context, arg db.ListUsersParams) ([]db.ListUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, arg)
	ret0, _ := ret[0].([]db.ListUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockStoreMockRecorder) ListUsers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockStore)(nil).ListUsers), ctx, arg)
}

// Ping mocks base method.
func (m *MockStore) Ping(ctx context.Context, timeout time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx, timeout)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockStoreMockRecorder) Ping(ctx, timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockStore)(nil).Ping), ctx, timeout)
}

// RemoveAccountTypeFromAccount mocks base method.
func (m *MockStore) RemoveAccountTypeFromAccount(ctx context.Context, arg db.RemoveAccountTypeFromAccountParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAccountTypeFromAccount", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAccountTypeFromAccount indicates an expected call of RemoveAccountTypeFromAccount.
func (mr *MockStoreMockRecorder) RemoveAccountTypeFromAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAccountTypeFromAccount", reflect.TypeOf((*MockStore)(nil).RemoveAccountTypeFromAccount), ctx, arg)
}

// RemoveAllRelationshipsForAccountAccountType mocks base method.
func (m *MockStore) RemoveAllRelationshipsForAccountAccountType(ctx context.Context, accountsID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllRelationshipsForAccountAccountType", ctx, accountsID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllRelationshipsForAccountAccountType indicates an expected call of RemoveAllRelationshipsForAccountAccountType.
func (mr *MockStoreMockRecorder) RemoveAllRelationshipsForAccountAccountType(ctx, accountsID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllRelationshipsForAccountAccountType", reflect.TypeOf((*MockStore)(nil).RemoveAllRelationshipsForAccountAccountType), ctx, accountsID)
}

// RemoveAllRelationshipsForAccountTypeAccount mocks base method.
func (m *MockStore) RemoveAllRelationshipsForAccountTypeAccount(ctx context.Context, accounttypesID int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllRelationshipsForAccountTypeAccount", ctx, accounttypesID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllRelationshipsForAccountTypeAccount indicates an expected call of RemoveAllRelationshipsForAccountTypeAccount.
func (mr *MockStoreMockRecorder) RemoveAllRelationshipsForAccountTypeAccount(ctx, accounttypesID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllRelationshipsForAccountTypeAccount", reflect.TypeOf((*MockStore)(nil).RemoveAllRelationshipsForAccountTypeAccount), ctx, accounttypesID)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(ctx context.Context, arg db.UpdateAccountParams) (db.UserSvcAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", ctx, arg)
	ret0, _ := ret[0].(db.UserSvcAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), ctx, arg)
}

// UpdateAccountType mocks base method.
func (m *MockStore) UpdateAccountType(ctx context.Context, arg db.UpdateAccountTypeParams) (db.UserSvcAccountType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccountType", ctx, arg)
	ret0, _ := ret[0].(db.UserSvcAccountType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccountType indicates an expected call of UpdateAccountType.
func (mr *MockStoreMockRecorder) UpdateAccountType(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccountType", reflect.TypeOf((*MockStore)(nil).UpdateAccountType), ctx, arg)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.UpdateUserRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, arg)
	ret0, _ := ret[0].(db.UpdateUserRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), ctx, arg)
}

// UpdateUserEmail mocks base method.
func (m *MockStore) UpdateUserEmail(ctx context.Context, arg db.UpdateUserEmailParams) (db.UpdateUserEmailRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmail", ctx, arg)
	ret0, _ := ret[0].(db.UpdateUserEmailRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserEmail indicates an expected call of UpdateUserEmail.
func (mr *MockStoreMockRecorder) UpdateUserEmail(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmail", reflect.TypeOf((*MockStore)(nil).UpdateUserEmail), ctx, arg)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(ctx context.Context, arg db.UpdateUserPasswordParams) (db.UpdateUserPasswordRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, arg)
	ret0, _ := ret[0].(db.UpdateUserPasswordRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), ctx, arg)
}

// UpdateUsername mocks base method.
func (m *MockStore) UpdateUsername(ctx context.Context, arg db.UpdateUsernameParams) (db.UpdateUsernameRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUsername", ctx, arg)
	ret0, _ := ret[0].(db.UpdateUsernameRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUsername indicates an expected call of UpdateUsername.
func (mr *MockStoreMockRecorder) UpdateUsername(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUsername", reflect.TypeOf((*MockStore)(nil).UpdateUsername), ctx, arg)
}
