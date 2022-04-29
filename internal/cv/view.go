package cv

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
)

type CVView struct {
	Model CVModelInterface
}

type CVViewInterface interface {
	CreateCV(cvData *models.CV, ownerEmail string) error
}

func NewCVView(cvModel CVModelInterface) *CVView {
	return &CVView{
		Model: cvModel,
	}
}

func (cs *CVView) CreateCV(cvData *models.CV, ownerEmail string) error {
	helpers.CreateCV(cvData)
	return nil
}
