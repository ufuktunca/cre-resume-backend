// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/cv/service.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCVServiceInterface is a mock of CVServiceInterface interface
type MockCVServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCVServiceInterfaceMockRecorder
}

// MockCVServiceInterfaceMockRecorder is the mock recorder for MockCVServiceInterface
type MockCVServiceInterfaceMockRecorder struct {
	mock *MockCVServiceInterface
}

// NewMockCVServiceInterface creates a new mock instance
func NewMockCVServiceInterface(ctrl *gomock.Controller) *MockCVServiceInterface {
	mock := &MockCVServiceInterface{ctrl: ctrl}
	mock.recorder = &MockCVServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCVServiceInterface) EXPECT() *MockCVServiceInterfaceMockRecorder {
	return _m.recorder
}

// CreateCV mocks base method
func (_m *MockCVServiceInterface) CreateCV(cvData *models.CV, ownerEmail string) error {
	ret := _m.ctrl.Call(_m, "CreateCV", cvData, ownerEmail)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCV indicates an expected call of CreateCV
func (_mr *MockCVServiceInterfaceMockRecorder) CreateCV(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateCV", reflect.TypeOf((*MockCVServiceInterface)(nil).CreateCV), arg0, arg1)
}