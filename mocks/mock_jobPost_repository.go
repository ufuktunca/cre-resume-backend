// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/job-post/repository.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobPostRepositoryInterface is a mock of JobPostRepositoryInterface interface
type MockJobPostRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobPostRepositoryInterfaceMockRecorder
}

// MockJobPostRepositoryInterfaceMockRecorder is the mock recorder for MockJobPostRepositoryInterface
type MockJobPostRepositoryInterfaceMockRecorder struct {
	mock *MockJobPostRepositoryInterface
}

// NewMockJobPostRepositoryInterface creates a new mock instance
func NewMockJobPostRepositoryInterface(ctrl *gomock.Controller) *MockJobPostRepositoryInterface {
	mock := &MockJobPostRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockJobPostRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockJobPostRepositoryInterface) EXPECT() *MockJobPostRepositoryInterfaceMockRecorder {
	return _m.recorder
}

// CreateJobPost mocks base method
func (_m *MockJobPostRepositoryInterface) CreateJobPost(jobPost *models.JobPost) error {
	ret := _m.ctrl.Call(_m, "CreateJobPost", jobPost)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateJobPost indicates an expected call of CreateJobPost
func (_mr *MockJobPostRepositoryInterfaceMockRecorder) CreateJobPost(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateJobPost", reflect.TypeOf((*MockJobPostRepositoryInterface)(nil).CreateJobPost), arg0)
}
