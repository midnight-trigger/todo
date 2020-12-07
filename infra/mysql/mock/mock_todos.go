// Code generated by MockGen. DO NOT EDIT.
// Source: todos.go

// Package mock_mysql is a generated GoMock package.
package mock_mysql

import (
	gomock "github.com/golang/mock/gomock"
)

// MockITodos is a mock of ITodos interface
type MockITodos struct {
	ctrl     *gomock.Controller
	recorder *MockITodosMockRecorder
}

// MockITodosMockRecorder is the mock recorder for MockITodos
type MockITodosMockRecorder struct {
	mock *MockITodos
}

// NewMockITodos creates a new mock instance
func NewMockITodos(ctrl *gomock.Controller) *MockITodos {
	mock := &MockITodos{ctrl: ctrl}
	mock.recorder = &MockITodosMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockITodos) EXPECT() *MockITodosMockRecorder {
	return m.recorder
}
