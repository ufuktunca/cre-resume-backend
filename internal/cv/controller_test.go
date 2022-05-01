package cv_test

import (
	"bytes"
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/cv"
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

func Test_CreateCVController(t *testing.T) {
	controller := gomock.NewController(t)
	mockCVView := mocks.NewMockCVViewInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("23423423423")
		assert.Nil(t, err)

		cvData := models.CV{
			CVName:      "ufuk-cv-1",
			NameSurname: "Ufuk tunca",
			PhoneNumber: "5255204514",
			Photo:       "klsdhfdljshfds",
			Github:      "asdhasdas",
			Linkedin:    "ldafjşasfas",
		}
		reqBody, err := json.Marshal(&cvData)
		assert.Nil(t, err)

		req, err := http.NewRequest(fiber.MethodPost, "/cv", bytes.NewReader(reqBody))
		assert.Nil(t, err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.Header.Add("Authorization", *token)

		cvController := cv.NewCVController(mockCVView)
		cvController.SetupRouteApp(app)

		mockCVView.
			EXPECT().
			CreateCV(&cvData, "23423423423").
			Return(nil)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	})
}

func Test_GetCVs(t *testing.T) {
	controller := gomock.NewController(t)
	mockCVView := mocks.NewMockCVViewInterface(controller)

	t.Run("GivenUserWhenSentGetCVsRequestThenShouldReturnUsersCVs", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("234234234")
		assert.Nil(t, err)

		expectedResult := &[]models.CV{
			{
				ID:      "2131231",
				OwnerID: "234234234",
			},
		}

		req, err := http.NewRequest(fiber.MethodGet, "/user/cv", nil)
		assert.Nil(t, err)
		req.Header.Add("Authorization", *token)

		cvController := cv.NewCVController(mockCVView)
		cvController.SetupRouteApp(app)

		mockCVView.
			EXPECT().
			GetCVs("234234234").
			Return(expectedResult, nil)

		resp, _ := app.Test(req)

		actualResult := &[]models.CV{}
		respBody, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)
		err = json.Unmarshal(respBody, &actualResult)
		assert.Nil(t, err)

		assert.Equal(t, 200, resp.StatusCode)
		assert.NotNil(t, actualResult)
	})
}
