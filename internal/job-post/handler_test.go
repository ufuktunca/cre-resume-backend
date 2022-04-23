package jobPost_test

import (
	"bytes"
	"cre-resume-backend/internal/auth"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/mocks"
	"encoding/json"
	"io/ioutil"
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

		token, err := auth.CreateToken("test@asdasdas.com")
		assert.Nil(t, err)

		cookie := &http.Cookie{
			Name:  "auth",
			Value: *token,
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

func Test_GetJobPosts(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostService := mocks.NewMockJobPostServiceInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("test@asdasdas.com")
		assert.Nil(t, err)
		cookie := &http.Cookie{
			Name:  "auth",
			Value: *token,
		}

		expectedResult := &[]models.JobPost{
			{
				ID:       "1",
				Title:    "Tesat",
				Content:  "asdkasidasd",
				Salary:   400,
				Category: "TestC",
				Location: "İstanbul",
				Image:    "asişdkasid",
			},
		}

		req, err := http.NewRequest(fiber.MethodGet, "/jobPost/employee", nil)

		assert.Nil(t, err)
		req.AddCookie(cookie)

		jobPostHandler := jobPost.NewJobPostHandler(mockJobPostService)
		jobPostHandler.SetupJobPostHandler(app)

		mockJobPostService.
			EXPECT().
			GetJobPosts("employee").
			Return(expectedResult, nil)

		resp, _ := app.Test(req)

		actualResult := &[]models.JobPost{}
		respBody, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		err = json.Unmarshal(respBody, actualResult)
		assert.Nil(t, err)

		assert.Equal(t, resp.StatusCode, 200)
		assert.Equal(t, expectedResult, actualResult)
	})
}
