package jobPost_test

import (
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateJobPost(t *testing.T) {
	t.Run("Given employee user When sent a create job post request Then should return job post data", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockJobPostModel := mocks.NewMockJobPostModelInterface(controller)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		testUser := &models.User{
			UserID: "dşasljşkas",
			Email:  "test@gmail.com",
		}

		jobPostData := models.JobPost{
			OwnerID:  testUser.UserID,
			Title:    "test",
			Content:  "test cont",
			Salary:   123213,
			Category: "DEv",
			Location: "İstanbul",
			Image:    "lasşdkjşjasdşkasdj",
			Type:     "employer",
		}

		mockJobPostModel.
			EXPECT().
			CreateJobPost(gomock.Any()).
			Return(nil)

		jobPostView := jobPost.NewJobPostView(mockJobPostModel, mockUserModel)
		jobPostData2, err := jobPostView.CreateJobPost(&jobPostData, "test@gmail.com")

		assert.NotNil(t, jobPostData2)
		assert.Nil(t, err)
	})
}

func Test_GetJobs(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostModel := mocks.NewMockJobPostModelInterface(controller)
	mockUserModel := mocks.NewMockUserModelInterface(controller)

	t.Run("GivenUserWhenGetJobPostsThenShouldReturnJobPosts", func(t *testing.T) {
		expectedResult := &[]models.JobPost{
			{
				ID:       "12",
				Title:    "test",
				Content:  "asdasdas",
				Salary:   4500,
				Category: "Developer",
				Location: "İstanbul",
				Image:    "asdasdasd",
				Type:     "employee",
			},
		}

		mockJobPostModel.
			EXPECT().
			GetJobPosts("employee", "Developer", "3000", "5000", "salary").
			Return(expectedResult, nil)

		jobPostView := jobPost.NewJobPostView(mockJobPostModel, mockUserModel)

		result, err := jobPostView.GetJobPosts("employee", "Developer", "3000", "5000", "salary")

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

//func Test_ApplyJob(t *testing.T) {
// controller := gomock.NewController(t)
// mockJobPostModel := mocks.NewMockJobPostModelInterface(controller)
// mockUserModel := mocks.NewMockUserModelInterface(controller)

// t.Run("GivenUserWhenApplyJobThenShouldReturnNoError", func(t *testing.T) {
// jobPostData := models.JobPost{
// 	ID:      "293849238",
// 	OwnerID: "234",
// }

// applyJobDTO := models.ApplyJobPostDTO{
// 	CVID: "askdjkas",
// }

// // mockJobPostModel.
// // 	EXPECT().
// // 	GetJobPostsWithUserID("test@gmail.com", "23874289374").
// // 	Return(nil, errors.New("error"))

// mockJobPostModel.
// 	EXPECT().
// 	GetJobPostByID("23874289374").
// 	Return(&jobPostData, nil)

// mockJobPostModel.
// 	EXPECT().
// 	CreateApplyJobPost(gomock.Any()).
// 	Return(nil)

// jobPostView := jobPost.NewJobPostView(mockJobPostModel, mockUserModel)

// err := jobPostView.ApplyJobPost(&applyJobDTO, "test@gmail.com", "23874289374")

// assert.Nil(t, err)
//})
//}

func Test_GetUserJobPosts(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostModel := mocks.NewMockJobPostModelInterface(controller)
	mockUserModel := mocks.NewMockUserModelInterface(controller)

	t.Run("Given User When gets jobpost with user id then should return jobposts", func(t *testing.T) {
		jobPostData := []models.JobPost{
			{ID: "293849238",
				OwnerID: "234"},
		}

		testUser := models.User{
			UserID: "234u23423",
			Email:  "test@gmail.com",
		}

		mockUserModel.
			EXPECT().
			GetUserByEmail("test@gmail.com").
			Return(&testUser, nil)

		mockJobPostModel.
			EXPECT().
			GetJobPostsWithUserID("test@gmail.com", "employee").
			Return(&jobPostData, nil)

		jobPostView := jobPost.NewJobPostView(mockJobPostModel, mockUserModel)

		jobPostData2, err := jobPostView.GetUserJobPosts("test@gmail.com", "employee")

		assert.Nil(t, err)
		assert.NotNil(t, jobPostData2)
	})
}

func Test_GetUserAppliedJobPosts(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostModel := mocks.NewMockJobPostModelInterface(controller)
	mockUserModel := mocks.NewMockUserModelInterface(controller)

	t.Run("Given User When gets applied posts then should return applied posts", func(t *testing.T) {

		mockJobPostModel.
			EXPECT().
			GetUserApplies("test@gmail.com").
			Return([]models.ApplyJobPost{}, nil)

		mockJobPostModel.
			EXPECT().
			GetUserJobPosts("test@gmail.com").
			Return([]models.JobPost{}, nil)

		jobPostView := jobPost.NewJobPostView(mockJobPostModel, mockUserModel)

		jobPostData2, err := jobPostView.GetUserAppliedJobs("test@gmail.com")

		assert.Nil(t, err)
		assert.NotNil(t, jobPostData2)
	})
}
