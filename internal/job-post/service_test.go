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
		mockJobPostRepository := mocks.NewMockJobPostRepositoryInterface(controller)
		mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

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

		mockUserRepository.
			EXPECT().
			GetUserByEmail("test@gmail.com").
			Return(testUser, nil)

		mockJobPostRepository.
			EXPECT().
			CreateJobPost(gomock.Any()).
			Return(nil)

		jobPostService := jobPost.NewJobPostService(mockJobPostRepository, mockUserRepository)
		jobPostData2, err := jobPostService.CreateJobPost(&jobPostData, "test@gmail.com")

		assert.NotNil(t, jobPostData2)
		assert.Nil(t, err)
	})
}

func Test_GetJobs(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostRepository := mocks.NewMockJobPostRepositoryInterface(controller)
	mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

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

		mockJobPostRepository.
			EXPECT().
			GetJobPosts("employee", "Developer", "3000", "5000", "salary").
			Return(expectedResult, nil)

		jobPostService := jobPost.NewJobPostService(mockJobPostRepository, mockUserRepository)

		result, err := jobPostService.GetJobPosts("employee", "Developer", "3000", "5000", "salary")

		assert.Nil(t, err)
		assert.NotNil(t, result)
	})
}

func Test_ApplyJob(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostRepository := mocks.NewMockJobPostRepositoryInterface(controller)
	mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

	t.Run("GivenUserWhenApplyJobThenShouldReturnNoError", func(t *testing.T) {
		jobPostData := models.JobPost{
			ID:      "293849238",
			OwnerID: "234",
		}

		applyJobDTO := models.ApplyJobPostDTO{
			CVID: "askdjkas",
		}

		testUser := models.User{
			UserID: "234u23423",
			Email:  "test@gmail.com",
		}

		mockUserRepository.
			EXPECT().
			GetUserByEmail("test@gmail.com").
			Return(&testUser, nil)

		mockJobPostRepository.
			EXPECT().
			GetJobPostByID("23874289374").
			Return(&jobPostData, nil)

		mockJobPostRepository.
			EXPECT().
			CreateApplyJobPost(gomock.Any()).
			Return(nil)

		jobPostService := jobPost.NewJobPostService(mockJobPostRepository, mockUserRepository)

		err := jobPostService.ApplyJobPost(&applyJobDTO, "test@gmail.com", "23874289374")

		assert.Nil(t, err)
	})
}
