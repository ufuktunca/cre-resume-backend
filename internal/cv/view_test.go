package cv_test

import (
	"cre-resume-backend/internal/cv"
	"cre-resume-backend/internal/models"
	"cre-resume-backend/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateCV(t *testing.T) {
	// t.Run("Given employee user When sent a create cv request Then should create a cv", func(t *testing.T) {
	// 	controller := gomock.NewController(t)
	// 	mockCVModel := mocks.NewMockCVModelInterface(controller)

	// 	cvData := models.CV{
	// 		CVName:      "ufuk-cv-1",
	// 		NameSurname: "Ufuk tunca",
	// 		PhoneNumber: "5255204514",
	// 		Photo:       "klsdhfdljshfds",
	// 		Github:      "asdhasdas",
	// 		Linkedin:    "ldafjşasfas",
	// 	}

	// 	mockCVModel.
	// 		EXPECT().
	// 		CreateCV(gomock.Any()).
	// 		Return(nil)

	// 	cvView := cv.NewCVView(mockCVModel)
	// 	err := cvView.CreateCV(&cvData, "test@gmail.com")

	// 	assert.Nil(t, err)
	// })
}

func Test_GetCV(t *testing.T) {
	t.Run("Given user When sent get CVs request Then should return user's CVs", func(t *testing.T) {
		controler := gomock.NewController(t)
		CVModel := mocks.NewMockCVModelInterface(controler)

		CVs := &[]models.CV{
			{
				ID:      "12312321",
				OwnerID: "3453453",
			},
		}

		CVModel.
			EXPECT().
			GetCVs("3453453").
			Return(CVs, nil)

		view := cv.NewCVView(CVModel)
		cvData, err := view.GetCVs("3453453")

		assert.Nil(t, err)
		assert.NotNil(t, cvData)
	})
}

func Test_GetSingleCV(t *testing.T) {
	t.Run("Given user When sent get CV request Then should return user's CV", func(t *testing.T) {
		controler := gomock.NewController(t)
		CVModel := mocks.NewMockCVModelInterface(controler)

		CVModel.
			EXPECT().
			GetCV("3453453").
			Return(&models.CV{}, nil)

		view := cv.NewCVView(CVModel)
		cvData, cvName, err := view.GetCV("3453453")

		assert.Nil(t, err)
		assert.NotNil(t, cvData)
		assert.NotNil(t, cvName)
	})
}
