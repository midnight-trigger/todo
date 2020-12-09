// Code generated by MockGen. DO NOT EDIT.
// Source: todos.go

// Package mock_mysql is a generated GoMock package.
package mock_mysql

import (
	gomock "github.com/golang/mock/gomock"
	mysql "github.com/midnight-trigger/todo/infra/mysql"
	reflect "reflect"
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

// FindById mocks base method
func (m *MockITodos) FindById(id int64) (mysql.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(mysql.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockITodosMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockITodos)(nil).FindById), id)
}

// Create mocks base method
func (m *MockITodos) Create(todo *mysql.Todos) (*mysql.Todos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", todo)
	ret0, _ := ret[0].(*mysql.Todos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockITodosMockRecorder) Create(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockITodos)(nil).Create), todo)
}

// Update mocks base method
func (m *MockITodos) Update(oldParams mysql.Todos, updateParams map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", oldParams, updateParams)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockITodosMockRecorder) Update(oldParams, updateParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockITodos)(nil).Update), oldParams, updateParams)
}

// Delete mocks base method
func (m *MockITodos) Delete(todo *mysql.Todos) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockITodosMockRecorder) Delete(todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockITodos)(nil).Delete), todo)
}
