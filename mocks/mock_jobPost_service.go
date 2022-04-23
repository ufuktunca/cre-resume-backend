// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/job-post/service.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobPostServiceInterface is a mock of JobPostServiceInterface interface
type MockJobPostServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobPostServiceInterfaceMockRecorder
}

// MockJobPostServiceInterfaceMockRecorder is the mock recorder for MockJobPostServiceInterface
type MockJobPostServiceInterfaceMockRecorder struct {
	mock *MockJobPostServiceInterface
}

// NewMockJobPostServiceInterface creates a new mock instance
func NewMockJobPostServiceInterface(ctrl *gomock.Controller) *MockJobPostServiceInterface {
	mock := &MockJobPostServiceInterface{ctrl: ctrl}
	mock.recorder = &MockJobPostServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockJobPostServiceInterface) EXPECT() *MockJobPostServiceInterfaceMockRecorder {
	return _m.recorder
}

// CreateJobPost mocks base method
func (_m *MockJobPostServiceInterface) CreateJobPost(jobPost *models.JobPost) (*models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "CreateJobPost", jobPost)
	ret0, _ := ret[0].(*models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobPost indicates an expected call of CreateJobPost
func (_mr *MockJobPostServiceInterfaceMockRecorder) CreateJobPost(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateJobPost", reflect.TypeOf((*MockJobPostServiceInterface)(nil).CreateJobPost), arg0)
}

// GetJobPosts mocks base method
func (_m *MockJobPostServiceInterface) GetJobPosts(jobPostType string, category string, from string, to string, sort string) (*[]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobPosts", jobPostType, category, from, to, sort)
	ret0, _ := ret[0].(*[]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPosts indicates an expected call of GetJobPosts
func (_mr *MockJobPostServiceInterfaceMockRecorder) GetJobPosts(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobPosts", reflect.TypeOf((*MockJobPostServiceInterface)(nil).GetJobPosts), arg0, arg1, arg2, arg3, arg4)
}
