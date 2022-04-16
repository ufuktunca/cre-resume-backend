package user

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repository UserRepositoryInterface
}

type UserServiceInterface interface {
	Register(register *models.User) error
	Login(login *models.Login) (*string, error)
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

	return s.Repository.CreateUser(register)
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

	expirationTime := time.Now().Add(12 * time.Hour)
	claims := models.Claims{
		Username: login.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
