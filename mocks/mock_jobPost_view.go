// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/job-post/view.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobPostViewInterface is a mock of JobPostViewInterface interface
type MockJobPostViewInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobPostViewInterfaceMockRecorder
}

// MockJobPostViewInterfaceMockRecorder is the mock recorder for MockJobPostViewInterface
type MockJobPostViewInterfaceMockRecorder struct {
	mock *MockJobPostViewInterface
}

// NewMockJobPostViewInterface creates a new mock instance
func NewMockJobPostViewInterface(ctrl *gomock.Controller) *MockJobPostViewInterface {
	mock := &MockJobPostViewInterface{ctrl: ctrl}
	mock.recorder = &MockJobPostViewInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockJobPostViewInterface) EXPECT() *MockJobPostViewInterfaceMockRecorder {
	return _m.recorder
}

// CreateJobPost mocks base method
func (_m *MockJobPostViewInterface) CreateJobPost(jobPost *models.JobPost, ownerEmail string) (*models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "CreateJobPost", jobPost, ownerEmail)
	ret0, _ := ret[0].(*models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateJobPost indicates an expected call of CreateJobPost
func (_mr *MockJobPostViewInterfaceMockRecorder) CreateJobPost(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateJobPost", reflect.TypeOf((*MockJobPostViewInterface)(nil).CreateJobPost), arg0, arg1)
}

// GetJobPosts mocks base method
func (_m *MockJobPostViewInterface) GetJobPosts(jobPostType string, category string, from string, to string, sort string) (*[]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobPosts", jobPostType, category, from, to, sort)
	ret0, _ := ret[0].(*[]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPosts indicates an expected call of GetJobPosts
func (_mr *MockJobPostViewInterfaceMockRecorder) GetJobPosts(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobPosts", reflect.TypeOf((*MockJobPostViewInterface)(nil).GetJobPosts), arg0, arg1, arg2, arg3, arg4)
}

// ApplyJobPost mocks base method
func (_m *MockJobPostViewInterface) ApplyJobPost(jobPostDTO *models.ApplyJobPostDTO, applierEmail string, jobID string) error {
	ret := _m.ctrl.Call(_m, "ApplyJobPost", jobPostDTO, applierEmail, jobID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyJobPost indicates an expected call of ApplyJobPost
func (_mr *MockJobPostViewInterfaceMockRecorder) ApplyJobPost(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "ApplyJobPost", reflect.TypeOf((*MockJobPostViewInterface)(nil).ApplyJobPost), arg0, arg1, arg2)
}

// GetUserJobPosts mocks base method
func (_m *MockJobPostViewInterface) GetUserJobPosts(userEmail string, postType string) (*[]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetUserJobPosts", userEmail, postType)
	ret0, _ := ret[0].(*[]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserJobPosts indicates an expected call of GetUserJobPosts
func (_mr *MockJobPostViewInterfaceMockRecorder) GetUserJobPosts(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserJobPosts", reflect.TypeOf((*MockJobPostViewInterface)(nil).GetUserJobPosts), arg0, arg1)
}
