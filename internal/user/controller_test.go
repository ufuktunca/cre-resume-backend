package user_test

import (
	"bytes"
	user_models "cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
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
		mockUserView := mocks.NewMockUserViewInterface(controller)

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

		userController := user.NewUserController(mockUserView)
		userController.SetupUserController(app)

		mockUserView.
			EXPECT().
			Register(&register).
			Return(nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 201)
	})
}

func Test_UserAcivation(t *testing.T) {
	t.Run("Given user When send a activation request Then should get status 200", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserView := mocks.NewMockUserViewInterface(controller)

		app := fiber.New()
		req, _ := http.NewRequest(fiber.MethodGet, "/activation?userID=askdjasd", nil)
		userController := user.NewUserController(mockUserView)
		userController.SetupUserController(app)

		mockUserView.
			EXPECT().
			ActivateUser("askdjasd").
			Return(nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 200)
	})
}

func Test_LoginController(t *testing.T) {
	t.Run("Given user When send a login request Then should get status 200", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserView := mocks.NewMockUserViewInterface(controller)

		app := fiber.New()

		login := user_models.Login{
			Email:    "ufutunca@gmail.com",
			Password: "123123",
		}

		reqBody, err := json.Marshal(&login)
		assert.Nil(t, err)

		req, _ := http.NewRequest(fiber.MethodPost, "/login?type=employee", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

		userController := user.NewUserController(mockUserView)
		userController.SetupUserController(app)

		data := "123213"
		mockUserView.
			EXPECT().
			Login(&login, "employee").
			Return(&data, nil)

		resp, _ := app.Test(req)

		assert.Equal(t, resp.StatusCode, 200)
	})
}

func Test_ReSendHandler(t *testing.T) {
	t.Run("Given user When send a resend activation email request Then should get status 200", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserView := mocks.NewMockUserViewInterface(controller)

		app := fiber.New()

		reSend := user_models.ReSend{
			Email: "ufutunca@gmail.com",
		}

		reqBody, err := json.Marshal(&reSend)
		assert.Nil(t, err)

		req, _ := http.NewRequest(fiber.MethodPost, "/reSend", bytes.NewReader(reqBody))
		req.Header.Add("Content-Type", "application/json")
		req.Header.Set("Content-Length", strconv.Itoa(len(reqBody)))

		userController := user.NewUserController(mockUserView)
		userController.SetupUserController(app)

		mockUserView.
			EXPECT().
			ReSend(reSend.Email).
			Return(nil)

		resp, _ := app.Test(req)
		assert.Equal(t, resp.StatusCode, 200)
	})
}
