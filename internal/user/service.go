package user

import "cre-resume-backend/internal/user/models"

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
	return nil
}
