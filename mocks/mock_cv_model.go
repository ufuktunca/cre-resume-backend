// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/cv/model.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCVModelInterface is a mock of CVModelInterface interface
type MockCVModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCVModelInterfaceMockRecorder
}

// MockCVModelInterfaceMockRecorder is the mock recorder for MockCVModelInterface
type MockCVModelInterfaceMockRecorder struct {
	mock *MockCVModelInterface
}

// NewMockCVModelInterface creates a new mock instance
func NewMockCVModelInterface(ctrl *gomock.Controller) *MockCVModelInterface {
	mock := &MockCVModelInterface{ctrl: ctrl}
	mock.recorder = &MockCVModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCVModelInterface) EXPECT() *MockCVModelInterfaceMockRecorder {
	return _m.recorder
}

// CreateCV mocks base method
func (_m *MockCVModelInterface) CreateCV(cvData models.CV) error {
	ret := _m.ctrl.Call(_m, "CreateCV", cvData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCV indicates an expected call of CreateCV
func (_mr *MockCVModelInterfaceMockRecorder) CreateCV(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateCV", reflect.TypeOf((*MockCVModelInterface)(nil).CreateCV), arg0)
}

// GetCVs mocks base method
func (_m *MockCVModelInterface) GetCVs(userID string) (*[]models.CV, error) {
	ret := _m.ctrl.Call(_m, "GetCVs", userID)
	ret0, _ := ret[0].(*[]models.CV)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCVs indicates an expected call of GetCVs
func (_mr *MockCVModelInterfaceMockRecorder) GetCVs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetCVs", reflect.TypeOf((*MockCVModelInterface)(nil).GetCVs), arg0)
}