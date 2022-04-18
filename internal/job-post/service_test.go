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

		jobPostData := models.JobPost{
			Title:    "test",
			Content:  "test cont",
			Salary:   123213,
			Category: "DEv",
			Location: "İstanbul",
			Image:    "lasşdkjşjasdşkasdj",
		}

		mockJobPostRepository.
			EXPECT().
			CreateJobPost(gomock.Any()).
			Return(nil)

		jobPostService := jobPost.NewJobPostService(mockJobPostRepository)
		jobPostData2, err := jobPostService.CreateJobPost(&jobPostData)

		assert.NotNil(t, jobPostData2)
		assert.Nil(t, err)
	})
}
