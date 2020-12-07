// Code generated by MockGen. DO NOT EDIT.
// Source: users.go

// Package mock_mysql is a generated GoMock package.
package mock_mysql

import (
	gomock "github.com/golang/mock/gomock"
)

// MockIUsers is a mock of IUsers interface
type MockIUsers struct {
	ctrl     *gomock.Controller
	recorder *MockIUsersMockRecorder
}

// MockIUsersMockRecorder is the mock recorder for MockIUsers
type MockIUsersMockRecorder struct {
	mock *MockIUsers
}

// NewMockIUsers creates a new mock instance
func NewMockIUsers(ctrl *gomock.Controller) *MockIUsers {
	mock := &MockIUsers{ctrl: ctrl}
	mock.recorder = &MockIUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUsers) EXPECT() *MockIUsersMockRecorder {
	return m.recorder
}
