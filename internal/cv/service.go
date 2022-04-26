package cv

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
)

type CVService struct {
	Repository CVRepositoryInterface
}

type CVServiceInterface interface {
	CreateCV(cvData *models.CV, ownerEmail string) error
}

func NewCVService(cvRepository CVRepositoryInterface) *CVService {
	return &CVService{
		Repository: cvRepository,
	}
}

func (cs *CVService) CreateCV(cvData *models.CV, ownerEmail string) error {
	helpers.CreateCV(cvData)
	return nil
}
