package cv

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"encoding/base64"
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
	byteData, err := helpers.CreateCV(cvData)
	if err != nil {
		return err
	}
	base64Data := base64.StdEncoding.EncodeToString(byteData)
	cvData.PDFCV = base64Data
	return cs.Model.CreateCV(*cvData)
}

func (cs *CVView) GetCVs(userID string) (*[]models.CV, error) {
	return cs.Model.GetCVs(userID)
}
