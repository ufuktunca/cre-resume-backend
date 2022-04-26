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

func Test_CreateCVHandler(t *testing.T) {
	controller := gomock.NewController(t)
	mockCVService := mocks.NewMockCVServiceInterface(controller)

	t.Run("GivenUserWhenSentGetJobPostRequestWithEmployeeParameterThenShouldReturnJobPosts", func(t *testing.T) {
		app := fiber.New()

		token, err := auth.CreateToken("test@asdasdas.com")
		assert.Nil(t, err)
		cookie := &http.Cookie{
			Name:  "auth",
			Value: *token,
		}

		cvData := models.CV{
			CVName:         "ufuk-cv-1",
			NameSurname:    "Ufuk tunca",
			PhoneNumber:    "5255204514",
			Hobbies:        "Test",
			Photo:          "klsdhfdljshfds",
			GraduateSchool: "Test",
			Experience:     "TEst exp",
			Github:         "asdhasdas",
			Linkedin:       "ldafjşasfas",
			OtherSM:        "şdlsfjkasfjas",
		}
		reqBody, err := json.Marshal(&cvData)
		assert.Nil(t, err)

		req, err := http.NewRequest(fiber.MethodPost, "/cv", bytes.NewReader(reqBody))
		assert.Nil(t, err)
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))
		req.AddCookie(cookie)

		cvHandler := cv.NewCVHandler(mockCVService)
		cvHandler.SetupRouteApp(app)

		// mockJobPostService.
		// 	EXPECT().
		// 	ApplyJobPost(&applyJobDTO, "test@asdasdas.com", "2938479").
		// 	Return(nil)

		resp, err := app.Test(req)
		assert.Nil(t, err)
		assert.Equal(t, resp.StatusCode, 200)
	})
}
