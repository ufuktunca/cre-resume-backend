package jobPost_test

import (
	"bytes"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/mocks"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateEmployeeJobPost(t *testing.T) {
	t.Run("Given employee When send a request to create job post Then should get status 201", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockJobPostService := mocks.NewMockJobPostServiceInterface(controller)

		app := fiber.New()

		jobPostData := models.JobPost{
			Title:    "Job",
			Content:  "test",
			Salary:   3000,
			Category: "Developer",
			Location: "İstanbul",
			Image:    "şalsdjşlasdas",
		}

		cookie := &http.Cookie{
			Name:  "auth",
			Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVmdWt0dW5jYUBnbWFpbC5jb20iLCJleHAiOjUyNTAxODEzNjF9.7HzHor9YdC0Jbwi939cQ5W4kotbdomA_OlMLj9KVd8U",
		}

		reqBody, err := json.Marshal(&jobPostData)
		assert.Nil(t, err)

		req, _ := http.NewRequest(fiber.MethodPost, "/jobPost/employee", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.AddCookie(cookie)

		jobPostHandler := jobPost.NewJobPostHandler(mockJobPostService)
		jobPostHandler.SetupJobPostHandler(app)

		mockJobPostService.
			EXPECT().
			CreateJobPost(&jobPostData).
			Return(nil, nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 201)
	})
}
