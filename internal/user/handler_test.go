package user_test

import (
	"bytes"
	"cre-resume-backend/internal/user"
	user_models "cre-resume-backend/internal/user/models"
	"cre-resume-backend/mocks"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_UserRegister(t *testing.T) {
	t.Run("Given user When send a register request Then should get status 200", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserService := mocks.NewMockUserServiceInterface(controller)

		app := fiber.New()

		register := user_models.User{
			Name:     "Ufuk",
			Surname:  "Tunca",
			Email:    "ufutunca@gmail.com",
			Password: "123123",
			Type:     "Employer",
		}

		reqBody, err := json.Marshal(&register)
		assert.Nil(t, err)

		req, _ := http.NewRequest(fiber.MethodPost, "/register", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

		userHandler := user.NewUserHandler(mockUserService)
		userHandler.SetupUserHandler(app)

		mockUserService.
			EXPECT().
			Register(&register).
			Return(nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 201)
	})
}
