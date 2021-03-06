// Code generated by MockGen. DO NOT EDIT.
// Source: dao.go

// Package users is a generated GoMock package.
package users

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMysqlDao is a mock of MysqlDao interface.
type MockMysqlDao struct {
	ctrl     *gomock.Controller
	recorder *MockMysqlDaoMockRecorder
}

// MockMysqlDaoMockRecorder is the mock recorder for MockMysqlDao.
type MockMysqlDaoMockRecorder struct {
	mock *MockMysqlDao
}

// NewMockMysqlDao creates a new mock instance.
func NewMockMysqlDao(ctrl *gomock.Controller) *MockMysqlDao {
	mock := &MockMysqlDao{ctrl: ctrl}
	mock.recorder = &MockMysqlDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMysqlDao) EXPECT() *MockMysqlDaoMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockMysqlDao) GetUser(arg0 context.Context, arg1 int64) (user, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(user)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockMysqlDaoMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockMysqlDao)(nil).GetUser), arg0, arg1)
}

// InsertUser mocks base method.
func (m *MockMysqlDao) InsertUser(arg0 context.Context, arg1 user) (user, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", arg0, arg1)
	ret0, _ := ret[0].(user)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockMysqlDaoMockRecorder) InsertUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockMysqlDao)(nil).InsertUser), arg0, arg1)
}
