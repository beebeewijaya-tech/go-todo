// Code generated by MockGen. DO NOT EDIT.
// Source: .\db\sql\queries.go

// Package mock_sql is a generated GoMock package.
package mock_sql

import (
	context "context"
	reflect "reflect"

	sql "beebeewijaya.com/db/sql"
	gomock "github.com/golang/mock/gomock"
)

// MockQueries is a mock of Queries interface.
type MockQueries struct {
	ctrl     *gomock.Controller
	recorder *MockQueriesMockRecorder
}

// MockQueriesMockRecorder is the mock recorder for MockQueries.
type MockQueriesMockRecorder struct {
	mock *MockQueries
}

// NewMockQueries creates a new mock instance.
func NewMockQueries(ctrl *gomock.Controller) *MockQueries {
	mock := &MockQueries{ctrl: ctrl}
	mock.recorder = &MockQueriesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueries) EXPECT() *MockQueriesMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockQueries) CreateTodo(ctx context.Context, args sql.CreateTodoArgs) (sql.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", ctx, args)
	ret0, _ := ret[0].(sql.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockQueriesMockRecorder) CreateTodo(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockQueries)(nil).CreateTodo), ctx, args)
}

// CreateUser mocks base method.
func (m *MockQueries) CreateUser(ctx context.Context, args sql.CreateUserArgs) (sql.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, args)
	ret0, _ := ret[0].(sql.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQueriesMockRecorder) CreateUser(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQueries)(nil).CreateUser), ctx, args)
}

// DeleteTodo mocks base method.
func (m *MockQueries) DeleteTodo(ctx context.Context, args sql.DeleteTodoArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", ctx, args)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockQueriesMockRecorder) DeleteTodo(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockQueries)(nil).DeleteTodo), ctx, args)
}

// GetTodo mocks base method.
func (m *MockQueries) GetTodo(ctx context.Context, id int64) (sql.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodo", ctx, id)
	ret0, _ := ret[0].(sql.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo.
func (mr *MockQueriesMockRecorder) GetTodo(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockQueries)(nil).GetTodo), ctx, id)
}

// GetTodos mocks base method.
func (m *MockQueries) GetTodos(ctx context.Context, args sql.GetTodosArgs) ([]sql.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodos", ctx, args)
	ret0, _ := ret[0].([]sql.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodos indicates an expected call of GetTodos.
func (mr *MockQueriesMockRecorder) GetTodos(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodos", reflect.TypeOf((*MockQueries)(nil).GetTodos), ctx, args)
}

// GetUser mocks base method.
func (m *MockQueries) GetUser(ctx context.Context, email string) (sql.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, email)
	ret0, _ := ret[0].(sql.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockQueriesMockRecorder) GetUser(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockQueries)(nil).GetUser), ctx, email)
}

// UpdateTodo mocks base method.
func (m *MockQueries) UpdateTodo(ctx context.Context, args sql.UpdateTodoArgs) (sql.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", ctx, args)
	ret0, _ := ret[0].(sql.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockQueriesMockRecorder) UpdateTodo(ctx, args interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockQueries)(nil).UpdateTodo), ctx, args)
}
