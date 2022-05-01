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
		mockJobPostView := mocks.NewMockJobPostViewInterface(controller)

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

		token, err := auth.CreateToken("2342342352")
		assert.Nil(t, err)

		reqBody, err := json.Marshal(&jobPostData)
		assert.Nil(t, err)

		req, _ := http.NewRequest(fiber.MethodPost, "/jobPost/employee", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.Header.Add("Authorization", *token)

		jobPostController := jobPost.NewJobPostController(mockJobPostView)
		jobPostController.SetupJobPostController(app)

		mockJobPostView.
			EXPECT().
			CreateJobPost(&jobPostData, "2342342352").
			Return(nil, nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 201)
	})
}

func Test_GetJobPosts(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostView := mocks.NewMockJobPostViewInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("234234234")
		assert.Nil(t, err)

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
		req.Header.Add("Authorization", *token)
		assert.Nil(t, err)

		jobPostController := jobPost.NewJobPostController(mockJobPostView)
		jobPostController.SetupJobPostController(app)

		mockJobPostView.
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

func Test_ApplyJobController(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostView := mocks.NewMockJobPostViewInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("234234234")
		assert.Nil(t, err)

		applyJobDTO := models.ApplyJobPostDTO{
			CVID: "askdjkas",
		}
		reqBody, err := json.Marshal(&applyJobDTO)
		assert.Nil(t, err)

		req, err := http.NewRequest(fiber.MethodPost, "/jobPost/2938479/apply", bytes.NewReader(reqBody))
		assert.Nil(t, err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.Header.Add("Authorization", *token)

		jobPostController := jobPost.NewJobPostController(mockJobPostView)
		jobPostController.SetupJobPostController(app)

		mockJobPostView.
			EXPECT().
			ApplyJobPost(&applyJobDTO, "234234234", "2938479").
			Return(nil)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	})
}

func Test_GetUsersJobPostsHanlder(t *testing.T) {
	controller := gomock.NewController(t)
	mockJobPostView := mocks.NewMockJobPostViewInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPostOnlyUserJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("234234234")
		assert.Nil(t, err)

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
		req.Header.Add("Authorization", *token)

		jobPostController := jobPost.NewJobPostController(mockJobPostView)
		jobPostController.SetupJobPostController(app)

		mockJobPostView.
			EXPECT().
			GetUserJobPosts("234234234", "employee").
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
