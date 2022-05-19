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
	GetCV(cvID string) ([]byte, string, error)
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
	cvData.OwnerID = userID
	cvData.ID = helpers.GenerateUUID(8)
	return cs.Model.CreateCV(*cvData)
}

func (cs *CVView) GetCVs(userID string) (*[]models.CV, error) {
	return cs.Model.GetCVs(userID)
}

func (cs *CVView) GetCV(cvID string) ([]byte, string, error) {
	cvPdf, err := cs.Model.GetCV(cvID)
	if err != nil {
		return nil, "", err
	}

	pdf, err := base64.StdEncoding.DecodeString(cvPdf.PDFCV)
	if err != nil {
		return nil, "", err
	}

	return pdf, cvPdf.CVName, nil
}
