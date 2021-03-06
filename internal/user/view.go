package user

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/email"
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type View struct {
	Model UserModelInterface
}

type UserViewInterface interface {
	Register(register *models.User) error
	Login(login *models.Login, loginType string) (*string, error)
	ActivateUser(userID string) error
	ReSend(email string) error
}

func NewUserView(userModel UserModelInterface) *View {
	return &View{
		Model: userModel,
	}
}

var jwtKey = []byte("sdfk1lmhd2342sklgfjdhas634flkdshj23oır42o3euıw")

func (s *View) Register(register *models.User) error {
	_, err := s.Model.GetUserByEmail(register.Email)

	if err == nil {
		return models.EmailError
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), 14)
	if err != nil {
		return err
	}
	register.Password = string(hashedPassword)
	register.UserID = helpers.GenerateUUID(8)
	register.Activation = false

	err = s.Model.CreateUser(register)
	if err != nil {
		return err
	}

	return email.SendMail(register.Email, models.RegistirationMailContent+"https://cre-resume.herokuapp.com/login?userID="+register.UserID)
}

func (s *View) Login(login *models.Login, loginType string) (*string, error) {
	createdUser, err := s.Model.GetUserByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	if createdUser.Type != loginType {
		return nil, errors.New("user type is not correct")
	}

	err = bcrypt.CompareHashAndPassword([]byte(createdUser.Password), []byte(login.Password))
	if err != nil {
		return nil, errors.New("Password is not matched!!!")
	}

	if !createdUser.Activation {
		return nil, models.ActivationError
	}

	token, err := auth.CreateToken(createdUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *View) ActivateUser(userID string) error {

	return s.Model.Activation(userID)
}

func (s *View) ReSend(userEmail string) error {
	user, err := s.Model.GetUserByEmail(userEmail)
	if err != nil {
		return err
	}

	return email.SendMail(user.Email, models.RegistirationMailContent+"localhost:3000/login?userID="+user.UserID)
}
