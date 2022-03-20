package user_test

import (
	"cre-resume-backend/internal/user"
	"cre-resume-backend/internal/user/models"
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
