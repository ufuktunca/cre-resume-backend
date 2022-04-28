// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/cv/repository.go

package mocks

import (
	gomock "github.com/golang/mock/gomock"
)

// MockCVRepositoryInterface is a mock of CVRepositoryInterface interface
type MockCVRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCVRepositoryInterfaceMockRecorder
}

// MockCVRepositoryInterfaceMockRecorder is the mock recorder for MockCVRepositoryInterface
type MockCVRepositoryInterfaceMockRecorder struct {
	mock *MockCVRepositoryInterface
}

// NewMockCVRepositoryInterface creates a new mock instance
func NewMockCVRepositoryInterface(ctrl *gomock.Controller) *MockCVRepositoryInterface {
	mock := &MockCVRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockCVRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockCVRepositoryInterface) EXPECT() *MockCVRepositoryInterfaceMockRecorder {
	return _m.recorder
}