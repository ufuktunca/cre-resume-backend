package user_test

import (
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"cre-resume-backend/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_Register(t *testing.T) {
	t.Run("Given user register When sent valid data and unique email Then should register ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

		register := models.User{
			UserID:   "123123",
			Name:     "Ufuk",
			Surname:  "Tunca",
			Email:    "ufuk.tunca@gmail.com",
			Password: "asdşfljsdaşkfjsd",
			Type:     "Employer",
		}

		mockUserRepository.
			EXPECT().
			GetUserByEmail(register.Email).
			Return(nil, errors.New("asdasd"))

		mockUserRepository.
			EXPECT().
			CreateUser(&register).
			Return(nil)

		userService := user.NewUserService(mockUserRepository)

		err := userService.Register(&register)

		assert.Nil(t, err)
	})

	t.Run("Given user register When sent valid data with used email Then should return error ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

		register := models.User{
			UserID:   "123123",
			Name:     "Ufuk",
			Surname:  "Tunca",
			Email:    "ufuk.tunca@gmail.com",
			Password: "asdşfljsdaşkfjsd",
			Type:     "Employer",
		}

		mockUserRepository.
			EXPECT().
			GetUserByEmail(register.Email).
			Return(nil, nil)

		userService := user.NewUserService(mockUserRepository)

		err := userService.Register(&register)

		assert.NotNil(t, err)
	})
}

func Test_Login(t *testing.T) {
	t.Run("Given user login When sent valid data Then should return jwt token ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

		login := &models.Login{
			Email:    "ufuktunca@gmail.com",
			Password: "qwe123",
		}

		userData := &models.User{
			Email:    "ufuk.tunca@gmail.com",
			Password: "$2a$14$6Vad1pGBrdI6FuWZKUfImutaCfJL8BNgqWJEBLReyyts6gLWQ64h.",
		}

		mockUserRepository.
			EXPECT().
			GetUserByEmail(login.Email).
			Return(userData, nil)
		userService := user.NewUserService(mockUserRepository)

		jwtToken, err := userService.Login(login)

		assert.Nil(t, err)
		assert.NotNil(t, jwtToken)
	})

	t.Run("Given user register When sent invalid data with unused email Then should return error ", func(t *testing.T) {
		controller := gomock.NewController(t)
		mockUserRepository := mocks.NewMockUserRepositoryInterface(controller)

		login := models.Login{
			Email:    "ufuk.tunca@gmail.com",
			Password: "asdşfljsdaşkfjsd",
		}

		mockUserRepository.
			EXPECT().
			GetUserByEmail(login.Email).
			Return(nil, errors.New("error"))

		userService := user.NewUserService(mockUserRepository)

		jwtToken, err := userService.Login(&login)

		assert.NotNil(t, err)
		assert.Nil(t, jwtToken)
	})
}
