// Code generated by MockGen. DO NOT EDIT.
// Source: actions.go

// Package auth0 is a generated GoMock package.
package auth0

import (
	gomock "github.com/golang/mock/gomock"
	management "gopkg.in/auth0.v5/management"
	reflect "reflect"
)

// MockActionsAPI is a mock of ActionsAPI interface
type MockActionsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockActionsAPIMockRecorder
}

// MockActionsAPIMockRecorder is the mock recorder for MockActionsAPI
type MockActionsAPIMockRecorder struct {
	mock *MockActionsAPI
}

// NewMockActionsAPI creates a new mock instance
func NewMockActionsAPI(ctrl *gomock.Controller) *MockActionsAPI {
	mock := &MockActionsAPI{ctrl: ctrl}
	mock.recorder = &MockActionsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionsAPI) EXPECT() *MockActionsAPIMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockActionsAPI) Create(a *management.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockActionsAPIMockRecorder) Create(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockActionsAPI)(nil).Create), a)
}

// Read mocks base method
func (m *MockActionsAPI) Read(id string) (*management.Action, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", id)
	ret0, _ := ret[0].(*management.Action)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockActionsAPIMockRecorder) Read(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockActionsAPI)(nil).Read), id)
}

// Update mocks base method
func (m *MockActionsAPI) Update(id string, a *management.Action) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockActionsAPIMockRecorder) Update(id, a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockActionsAPI)(nil).Update), id, a)
}

// Delete mocks base method
func (m *MockActionsAPI) Delete(id string, opts ...management.RequestOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{id}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockActionsAPIMockRecorder) Delete(id interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{id}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockActionsAPI)(nil).Delete), varargs...)
}

// List mocks base method
func (m *MockActionsAPI) List(opts ...management.RequestOption) (*management.ActionList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*management.ActionList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockActionsAPIMockRecorder) List(opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockActionsAPI)(nil).List), opts...)
}