package cv

import (
	"cre-resume-backend/internal/helpers"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/internal/user"
	"encoding/base64"
	"fmt"
)

type CVView struct {
	Model     CVModelInterface
	UserModel user.UserModelInterface
}

type CVViewInterface interface {
	CreateCV(cvData *models.CV, userID string) error
	GetCVs(userID string) (*[]models.CV, error)
	GetCV(cvID string) ([]byte, string, error)
}

func NewCVView(cvModel CVModelInterface, userModel user.UserModelInterface) *CVView {
	return &CVView{
		Model:     cvModel,
		UserModel: userModel,
	}
}

func (cs *CVView) CreateCV(cvData *models.CV, userID string) error {
	user, err := cs.UserModel.GetUserByEmail(userID)
	if err != nil {
		fmt.Println("err", err, userID)
		return err
	}

	byteData, err := helpers.CreateCV(cvData, user)
	if err != nil {
		fmt.Println("err2", err)
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
