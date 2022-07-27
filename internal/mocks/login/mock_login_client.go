// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/eneskzlcn/dictionary-app-cli/internal/login (interfaces: LoginClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	login "github.com/eneskzlcn/dictionary-app-cli/internal/login"
	gomock "github.com/golang/mock/gomock"
)

// MockLoginClient is a mock of LoginClient interface.
type MockLoginClient struct {
	ctrl     *gomock.Controller
	recorder *MockLoginClientMockRecorder
}

// MockLoginClientMockRecorder is the mock recorder for MockLoginClient.
type MockLoginClientMockRecorder struct {
	mock *MockLoginClient
}

// NewMockLoginClient creates a new mock instance.
func NewMockLoginClient(ctrl *gomock.Controller) *MockLoginClient {
	mock := &MockLoginClient{ctrl: ctrl}
	mock.recorder = &MockLoginClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoginClient) EXPECT() *MockLoginClientMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockLoginClient) Login(arg0 login.SignInRequest) login.SignInResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0)
	ret0, _ := ret[0].(login.SignInResponse)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockLoginClientMockRecorder) Login(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockLoginClient)(nil).Login), arg0)
}