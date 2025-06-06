// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/t-lunch/t-lunch-backend/internal/service (interfaces: UserRepo)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/user.go -package=mocks . UserRepo
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/t-lunch/t-lunch-backend/internal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockUserRepo is a mock of UserRepo interface.
type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
	isgomock struct{}
}

// MockUserRepoMockRecorder is the mock recorder for MockUserRepo.
type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

// NewMockUserRepo creates a new mock instance.
func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepo) CreateUser(ctx context.Context, user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepoMockRecorder) CreateUser(ctx, user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepo)(nil).CreateUser), ctx, user)
}

// GetUserByID mocks base method.
func (m *MockUserRepo) GetUserByID(ctx context.Context, id int64) (*models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(*models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockUserRepoMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserRepo)(nil).GetUserByID), ctx, id)
}

// GetUserPasswordByEmail mocks base method.
func (m *MockUserRepo) GetUserPasswordByEmail(ctx context.Context, email string) (int64, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserPasswordByEmail", ctx, email)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetUserPasswordByEmail indicates an expected call of GetUserPasswordByEmail.
func (mr *MockUserRepoMockRecorder) GetUserPasswordByEmail(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserPasswordByEmail", reflect.TypeOf((*MockUserRepo)(nil).GetUserPasswordByEmail), ctx, email)
}

// GetUsersByIDs mocks base method.
func (m *MockUserRepo) GetUsersByIDs(ctx context.Context, ids []int64) ([]*models.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersByIDs", ctx, ids)
	ret0, _ := ret[0].([]*models.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersByIDs indicates an expected call of GetUsersByIDs.
func (mr *MockUserRepoMockRecorder) GetUsersByIDs(ctx, ids any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersByIDs", reflect.TypeOf((*MockUserRepo)(nil).GetUsersByIDs), ctx, ids)
}

// IsUserWithEmailExist mocks base method.
func (m *MockUserRepo) IsUserWithEmailExist(ctx context.Context, email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserWithEmailExist", ctx, email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserWithEmailExist indicates an expected call of IsUserWithEmailExist.
func (mr *MockUserRepoMockRecorder) IsUserWithEmailExist(ctx, email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserWithEmailExist", reflect.TypeOf((*MockUserRepo)(nil).IsUserWithEmailExist), ctx, email)
}

// UpdateUserByID mocks base method.
func (m *MockUserRepo) UpdateUserByID(ctx context.Context, id int64, updates map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserByID", ctx, id, updates)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserByID indicates an expected call of UpdateUserByID.
func (mr *MockUserRepoMockRecorder) UpdateUserByID(ctx, id, updates any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserByID", reflect.TypeOf((*MockUserRepo)(nil).UpdateUserByID), ctx, id, updates)
}
