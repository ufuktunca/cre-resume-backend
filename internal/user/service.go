package user

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/email"
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository UserRepositoryInterface
}

type UserServiceInterface interface {
	Register(register *models.User) error
	Login(login *models.Login) (*string, error)
	ActivateUser(userID string) error
}

func NewUserService(userRepository UserRepositoryInterface) *Service {
	return &Service{
		Repository: userRepository,
	}
}

var jwtKey = []byte("sdfk1lmhd2342sklgfjdhas634flkdshj23oır42o3euıw")

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
	register.UserID = helpers.GenerateUUID(8)
	register.UserActivate = false

	err = s.Repository.CreateUser(register)
	if err != nil {
		return err
	}

	return email.SendMail(register.Email, models.RegistirationMailContent+"localhost:8080/verify?userID="+register.UserID)
}

func (s *Service) Login(login *models.Login) (*string, error) {
	createdUser, err := s.Repository.GetUserByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(createdUser.Password), []byte(login.Password))
	if err != nil {
		return nil, errors.New("Password is not matched!!!")
	}

	token, err := auth.CreateToken(createdUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *Service) ActivateUser(userID string) error {

	return s.Repository.Activation(userID)
}
