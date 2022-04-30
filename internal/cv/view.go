package cv

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
)

type CVView struct {
	Model CVModelInterface
}

type CVViewInterface interface {
	CreateCV(cvData *models.CV, userID string) error
	GetCVs(userID string) (*[]models.CV, error)
}

func NewCVView(cvModel CVModelInterface) *CVView {
	return &CVView{
		Model: cvModel,
	}
}

func (cs *CVView) CreateCV(cvData *models.CV, userID string) error {
	bytePdf, err := helpers.CreateCV(cvData)
	if err != nil {
		return err
	}

	cvData.PDFCV = string(bytePdf)
	return nil
}

func (cs *CVView) GetCVs(userID string) (*[]models.CV, error) {
	return cs.Model.GetCVs(userID)
}
