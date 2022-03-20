package user

import (
	"cre-resume-backend/internal/user/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository UserRepositoryInterface
}

type UserServiceInterface interface {
	Register(register *models.User) error
}

func NewUserService(userRepository UserRepositoryInterface) *Service {
	return &Service{
		Repository: userRepository,
	}
}

func (s *Service) Register(register *models.User) error {
	_, err := s.Repository.GetUserByEmail(register.Email)

	if err == nil {
		return errors.New("This email address already used")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), 14)
	if err != nil {
		return err
	}
	register.Password = string(hashedPassword)

	return s.Repository.CreateUser(register)
}
