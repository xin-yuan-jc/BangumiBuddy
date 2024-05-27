// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package auth is a generated GoMock package.
package auth

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthenticator is a mock of Authenticator interface.
type MockAuthenticator struct {
	ctrl     *gomock.Controller
	recorder *MockAuthenticatorMockRecorder
}

// MockAuthenticatorMockRecorder is the mock recorder for MockAuthenticator.
type MockAuthenticatorMockRecorder struct {
	mock *MockAuthenticator
}

// NewMockAuthenticator creates a new mock instance.
func NewMockAuthenticator(ctrl *gomock.Controller) *MockAuthenticator {
	mock := &MockAuthenticator{ctrl: ctrl}
	mock.recorder = &MockAuthenticatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthenticator) EXPECT() *MockAuthenticatorMockRecorder {
	return m.recorder
}

// Authorize mocks base method.
func (m *MockAuthenticator) Authorize(ctx context.Context, username, password string) (Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorize", ctx, username, password)
	ret0, _ := ret[0].(Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authorize indicates an expected call of Authorize.
func (mr *MockAuthenticatorMockRecorder) Authorize(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorize", reflect.TypeOf((*MockAuthenticator)(nil).Authorize), ctx, username, password)
}

// CheckAccessToken mocks base method.
func (m *MockAuthenticator) CheckAccessToken(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAccessToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAccessToken indicates an expected call of CheckAccessToken.
func (mr *MockAuthenticatorMockRecorder) CheckAccessToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAccessToken", reflect.TypeOf((*MockAuthenticator)(nil).CheckAccessToken), ctx, token)
}

// RefreshCredentials mocks base method.
func (m *MockAuthenticator) RefreshCredentials(ctx context.Context, refreshToken string) (Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshCredentials", ctx, refreshToken)
	ret0, _ := ret[0].(Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshCredentials indicates an expected call of RefreshCredentials.
func (mr *MockAuthenticatorMockRecorder) RefreshCredentials(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshCredentials", reflect.TypeOf((*MockAuthenticator)(nil).RefreshCredentials), ctx, refreshToken)
}

// UpdateUser mocks base method.
func (m *MockAuthenticator) UpdateUser(ctx context.Context, username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockAuthenticatorMockRecorder) UpdateUser(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockAuthenticator)(nil).UpdateUser), ctx, username, password)
}
