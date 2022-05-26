// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/job-post/model.go

package mocks

import (
	models "cre-resume-backend/internal/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockJobPostModelInterface is a mock of JobPostModelInterface interface
type MockJobPostModelInterface struct {
	ctrl     *gomock.Controller
	recorder *MockJobPostModelInterfaceMockRecorder
}

// MockJobPostModelInterfaceMockRecorder is the mock recorder for MockJobPostModelInterface
type MockJobPostModelInterfaceMockRecorder struct {
	mock *MockJobPostModelInterface
}

// NewMockJobPostModelInterface creates a new mock instance
func NewMockJobPostModelInterface(ctrl *gomock.Controller) *MockJobPostModelInterface {
	mock := &MockJobPostModelInterface{ctrl: ctrl}
	mock.recorder = &MockJobPostModelInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockJobPostModelInterface) EXPECT() *MockJobPostModelInterfaceMockRecorder {
	return _m.recorder
}

// CreateJobPost mocks base method
func (_m *MockJobPostModelInterface) CreateJobPost(jobPost *models.JobPost) error {
	ret := _m.ctrl.Call(_m, "CreateJobPost", jobPost)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateJobPost indicates an expected call of CreateJobPost
func (_mr *MockJobPostModelInterfaceMockRecorder) CreateJobPost(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateJobPost", reflect.TypeOf((*MockJobPostModelInterface)(nil).CreateJobPost), arg0)
}

// GetJobPosts mocks base method
func (_m *MockJobPostModelInterface) GetJobPosts(jobPostType string, category string, from string, to string, sort string) (*[]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobPosts", jobPostType, category, from, to, sort)
	ret0, _ := ret[0].(*[]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPosts indicates an expected call of GetJobPosts
func (_mr *MockJobPostModelInterfaceMockRecorder) GetJobPosts(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobPosts", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetJobPosts), arg0, arg1, arg2, arg3, arg4)
}

// GetJobPostByID mocks base method
func (_m *MockJobPostModelInterface) GetJobPostByID(id string) (*models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobPostByID", id)
	ret0, _ := ret[0].(*models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPostByID indicates an expected call of GetJobPostByID
func (_mr *MockJobPostModelInterfaceMockRecorder) GetJobPostByID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobPostByID", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetJobPostByID), arg0)
}

// CreateApplyJobPost mocks base method
func (_m *MockJobPostModelInterface) CreateApplyJobPost(applyJobPost *models.ApplyJobPost) error {
	ret := _m.ctrl.Call(_m, "CreateApplyJobPost", applyJobPost)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateApplyJobPost indicates an expected call of CreateApplyJobPost
func (_mr *MockJobPostModelInterfaceMockRecorder) CreateApplyJobPost(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "CreateApplyJobPost", reflect.TypeOf((*MockJobPostModelInterface)(nil).CreateApplyJobPost), arg0)
}

// GetJobPostsWithUserID mocks base method
func (_m *MockJobPostModelInterface) GetJobPostsWithUserID(id string, postType string) (*[]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobPostsWithUserID", id, postType)
	ret0, _ := ret[0].(*[]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobPostsWithUserID indicates an expected call of GetJobPostsWithUserID
func (_mr *MockJobPostModelInterfaceMockRecorder) GetJobPostsWithUserID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobPostsWithUserID", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetJobPostsWithUserID), arg0, arg1)
}

// GetJobApplyWithUserIDAndJobID mocks base method
func (_m *MockJobPostModelInterface) GetJobApplyWithUserIDAndJobID(userId string, jobID string) (*models.ApplyJobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobApplyWithUserIDAndJobID", userId, jobID)
	ret0, _ := ret[0].(*models.ApplyJobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobApplyWithUserIDAndJobID indicates an expected call of GetJobApplyWithUserIDAndJobID
func (_mr *MockJobPostModelInterfaceMockRecorder) GetJobApplyWithUserIDAndJobID(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobApplyWithUserIDAndJobID", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetJobApplyWithUserIDAndJobID), arg0, arg1)
}

// GetUserApplies mocks base method
func (_m *MockJobPostModelInterface) GetUserApplies(userId string) ([]models.ApplyJobPost, error) {
	ret := _m.ctrl.Call(_m, "GetUserApplies", userId)
	ret0, _ := ret[0].([]models.ApplyJobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserApplies indicates an expected call of GetUserApplies
func (_mr *MockJobPostModelInterfaceMockRecorder) GetUserApplies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserApplies", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetUserApplies), arg0)
}

// GetUserJobPosts mocks base method
func (_m *MockJobPostModelInterface) GetUserJobPosts(userId string) ([]models.JobPost, error) {
	ret := _m.ctrl.Call(_m, "GetUserJobPosts", userId)
	ret0, _ := ret[0].([]models.JobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserJobPosts indicates an expected call of GetUserJobPosts
func (_mr *MockJobPostModelInterfaceMockRecorder) GetUserJobPosts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetUserJobPosts", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetUserJobPosts), arg0)
}

// GetJobApplies mocks base method
func (_m *MockJobPostModelInterface) GetJobApplies(jobId string) (*[]models.ApplyJobPost, error) {
	ret := _m.ctrl.Call(_m, "GetJobApplies", jobId)
	ret0, _ := ret[0].(*[]models.ApplyJobPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobApplies indicates an expected call of GetJobApplies
func (_mr *MockJobPostModelInterfaceMockRecorder) GetJobApplies(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetJobApplies", reflect.TypeOf((*MockJobPostModelInterface)(nil).GetJobApplies), arg0)
}

// DeleteJobPost mocks base method
func (_m *MockJobPostModelInterface) DeleteJobPost(jobId string) error {
	ret := _m.ctrl.Call(_m, "DeleteJobPost", jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJobPost indicates an expected call of DeleteJobPost
func (_mr *MockJobPostModelInterfaceMockRecorder) DeleteJobPost(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "DeleteJobPost", reflect.TypeOf((*MockJobPostModelInterface)(nil).DeleteJobPost), arg0)
}
