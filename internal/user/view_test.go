package user_test

import (
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"cre-resume-backend/mocks"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	t.Run("Given user register When sent valid data and unique email Then should register ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		register := models.User{
			UserID:       "123123",
			Name:         "Ufuk",
			Surname:      "Tunca",
			Email:        "uftunca72@gmail.com",
			Password:     "asdşfljsdaşkfjsd",
			Type:         "Employer",
			UserActivate: false,
		}

		mockUserModel.
			EXPECT().
			GetUserByEmail(register.Email).
			Return(nil, errors.New("asdasd"))

		mockUserModel.
			EXPECT().
			CreateUser(&register).
			Return(nil)

		userView := user.NewUserView(mockUserModel)

		err := userView.Register(&register)

		assert.Nil(t, err)
	})

	t.Run("Given user register When sent valid data with used email Then should return error ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		register := models.User{
			UserID:   "123123",
			Name:     "Ufuk",
			Surname:  "Tunca",
			Email:    "uftunca72@gmail.com",
			Password: "asdşfljsdaşkfjsd",
			Type:     "Employer",
		}

		mockUserModel.
			EXPECT().
			GetUserByEmail(register.Email).
			Return(nil, nil)

		userView := user.NewUserView(mockUserModel)

		err := userView.Register(&register)

		assert.NotNil(t, err)
	})
}

func Test_Login(t *testing.T) {
	t.Run("Given user login When sent valid data Then should return jwt token ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		login := &models.Login{
			Email:    "ufuktunca@gmail.com",
			Password: "qwe123",
		}

		userData := &models.User{
			Email:        "ufuk.tunca@gmail.com",
			Password:     "$2a$14$6Vad1pGBrdI6FuWZKUfImutaCfJL8BNgqWJEBLReyyts6gLWQ64h.",
			Type:         "employee",
			UserActivate: true,
		}

		mockUserModel.
			EXPECT().
			GetUserByEmail(login.Email).
			Return(userData, nil)
		userView := user.NewUserView(mockUserModel)

		jwtToken, err := userView.Login(login, "employee")

		assert.Nil(t, err)
		assert.NotNil(t, jwtToken)
	})

	t.Run("Given user register When sent invalid data with unused email Then should return error ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		login := models.Login{
			Email:    "ufuk.tunca@gmail.com",
			Password: "asdşfljsdaşkfjsd",
		}

		mockUserModel.
			EXPECT().
			GetUserByEmail(login.Email).
			Return(nil, errors.New("error"))

		userView := user.NewUserView(mockUserModel)

		jwtToken, err := userView.Login(&login, "employee")

		assert.NotNil(t, err)
		assert.Nil(t, jwtToken)
	})
}

func Test_Activation(t *testing.T) {
	t.Run("Given user login When sent valid data Then should change activatin ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		mockUserModel.
			EXPECT().
			Activation("kasjdklasd").
			Return(nil)

		userView := user.NewUserView(mockUserModel)
		err := userView.ActivateUser("kasjdklasd")

		assert.Nil(t, err)
	})
}

func Test_ReSend(t *testing.T) {
	t.Run("Given user login When sent valid data Then should change activatin ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserModel := mocks.NewMockUserModelInterface(controller)

		mockUserModel.
			EXPECT().
			GetUserByEmail("ufukbaristunca@windowslive.com").
			Return(&models.User{
				Email: "ufukbaristunca@windowslive.com",
			}, nil)

		userView := user.NewUserView(mockUserModel)
		err := userView.ReSend("ufukbaristunca@windowslive.com")
		fmt.Println(err)
		assert.Nil(t, err)
	})
}
