// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/user/model.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserModelInterface is a mock of UserModelInterface interface
type MockUserModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserModelInterfaceMockRecorder
}

// MockUserModelInterfaceMockRecorder is the mock recorder for MockUserModelInterface
type MockUserModelInterfaceMockRecorder struct {
	mock *MockUserModelInterface
}

// NewMockUserModelInterface creates a new mock instance
func NewMockUserModelInterface(ctrl *gomock.Controller) *MockUserModelInterface {
	mock := &MockUserModelInterface{ctrl: ctrl}
	mock.recorder = &MockUserModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockUserModelInterface) EXPECT() *MockUserModelInterfaceMockRecorder {
	return _m.recorder
}

// GetUserByEmail mocks base method
func (_m *MockUserModelInterface) GetUserByEmail(email string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail
func (_mr *MockUserModelInterfaceMockRecorder) GetUserByEmail(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserModelInterface)(nil).GetUserByEmail), arg0)
}

// CreateUser mocks base method
func (_m *MockUserModelInterface) CreateUser(user *models.User) error {
	ret := _m.ctrl.Call(_m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser
func (_mr *MockUserModelInterfaceMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateUser", reflect.TypeOf((*MockUserModelInterface)(nil).CreateUser), arg0)
}

// Activation mocks base method
func (_m *MockUserModelInterface) Activation(userID string) error {
	ret := _m.ctrl.Call(_m, "Activation", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Activation indicates an expected call of Activation
func (_mr *MockUserModelInterfaceMockRecorder) Activation(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Activation", reflect.TypeOf((*MockUserModelInterface)(nil).Activation), arg0)
}

// GetUserByUserID mocks base method
func (_m *MockUserModelInterface) GetUserByUserID(userID string) (*models.User, error) {
	ret := _m.ctrl.Call(_m, "GetUserByUserID", userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUserID indicates an expected call of GetUserByUserID
func (_mr *MockUserModelInterfaceMockRecorder) GetUserByUserID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserByUserID", reflect.TypeOf((*MockUserModelInterface)(nil).GetUserByUserID), arg0)
}
