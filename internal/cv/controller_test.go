package cv_test

import (
	"bytes"
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/cv"
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

func Test_CreateCVController(t *testing.T) {
	controller := gomock.NewController(t)
	mockCVView := mocks.NewMockCVViewInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("23423423423")
		assert.Nil(t, err)
		cookie := &http.Cookie{
			Name:  "auth",
			Value: *token,
		}

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
		req.AddCookie(cookie)

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
