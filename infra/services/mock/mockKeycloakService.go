// Code generated by MockGen. DO NOT EDIT.
// Source: infra/services/keycloakService.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	requests "github.com/TulioGuaraldoB/cross-roads/core/dtos/requests"
	gomock "github.com/golang/mock/gomock"
)

// MockIKeyCloakService is a mock of IKeyCloakService interface.
type MockIKeyCloakService struct {
	ctrl     *gomock.Controller
	recorder *MockIKeyCloakServiceMockRecorder
}

// MockIKeyCloakServiceMockRecorder is the mock recorder for MockIKeyCloakService.
type MockIKeyCloakServiceMockRecorder struct {
	mock *MockIKeyCloakService
}

// NewMockIKeyCloakService creates a new mock instance.
func NewMockIKeyCloakService(ctrl *gomock.Controller) *MockIKeyCloakService {
	mock := &MockIKeyCloakService{ctrl: ctrl}
	mock.recorder = &MockIKeyCloakServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIKeyCloakService) EXPECT() *MockIKeyCloakServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockIKeyCloakService) CreateUser(ctx context.Context, user *requests.UserRequest) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIKeyCloakServiceMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIKeyCloakService)(nil).CreateUser), ctx, user)
}

// Login mocks base method.
func (m *MockIKeyCloakService) Login(ctx context.Context, credentials *requests.Credentials) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, credentials)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockIKeyCloakServiceMockRecorder) Login(ctx, credentials interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockIKeyCloakService)(nil).Login), ctx, credentials)
}
