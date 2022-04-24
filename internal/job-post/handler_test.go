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
			Type:     "employee",
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
			CreateJobPost(&jobPostData, "test@asdasdas.com").
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
				Salary:   4000,
				Category: "TestC",
				Location: "İstanbul",
				Image:    "asişdkasid",
			},
		}

		req, err := http.NewRequest(fiber.MethodGet, "/jobPost/employee?category=Developer&from=3000&to=5000&sort=salary", nil)

		assert.Nil(t, err)
		req.AddCookie(cookie)

		jobPostHandler := jobPost.NewJobPostHandler(mockJobPostService)
		jobPostHandler.SetupJobPostHandler(app)

		mockJobPostService.
			EXPECT().
			GetJobPosts("employee", "Developer", "3000", "5000", "salary").
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

func Test_ApplyJobHandler(t *testing.T) {
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

		applyJobDTO := models.ApplyJobPostDTO{
			CVID: "askdjkas",
		}
		reqBody, err := json.Marshal(&applyJobDTO)
		assert.Nil(t, err)

		req, err := http.NewRequest(fiber.MethodPost, "/jobPost/2938479/apply", bytes.NewReader(reqBody))
		assert.Nil(t, err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.AddCookie(cookie)

		jobPostHandler := jobPost.NewJobPostHandler(mockJobPostService)
		jobPostHandler.SetupJobPostHandler(app)

		mockJobPostService.
			EXPECT().
			ApplyJobPost(&applyJobDTO, "test@asdasdas.com", "2938479").
			Return(nil)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	})
}

func Test_GetUsersJobPostsHanlder(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostService := mocks.NewMockJobPostServiceInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPostOnlyUserJobPosts", func(t *testing.T) {
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
				Salary:   4000,
				Category: "TestC",
				Location: "İstanbul",
				Image:    "asişdkasid",
			},
		}

		req, err := http.NewRequest(fiber.MethodGet, "/user/jobPost/employee", nil)

		assert.Nil(t, err)
		req.AddCookie(cookie)

		jobPostHandler := jobPost.NewJobPostHandler(mockJobPostService)
		jobPostHandler.SetupJobPostHandler(app)

		mockJobPostService.
			EXPECT().
			GetUserJobPosts("test@asdasdas.com", "employee").
			Return(expectedResult, nil)

		resp, _ := app.Test(req)

		actualResult := &[]models.JobPost{}
		respBody, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		err = json.Unmarshal(respBody, &actualResult)
		assert.Nil(t, err)

		assert.Equal(t, resp.StatusCode, 200)
	})
}
