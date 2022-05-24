package helpers

import (
	"bytes"
	"cre-resume-backend/internal/models"
	"encoding/base64"
	"image"
	"image/png"
	"log"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/signintech/gopdf"
)

func GenerateUUID(length int) string {
	uuid := uuid.New().String()

	uuid = strings.ReplaceAll(uuid, "-", "")

	if length < 1 {
		return uuid
	}
	if length > len(uuid) {
		length = len(uuid)
	}

	return uuid[0:length]
}

func CreateCV(cvData *models.CV, user *models.User) ([]byte, error) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 840, H: 1188}})

	err := pdf.AddTTFFont("wts11", "./internal/helpers/Roboto-Regular.ttf")
	if err != nil {
		return nil, err
	}

	err = pdf.AddTTFFont("robotoBold", "./internal/helpers/Roboto-Bold.ttf")
	if err != nil {
		return nil, err
	}

	pdf.AddPage()

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	pdf.SetFillColor(53, 59, 69)
	pdf.RectFromUpperLeftWithStyle(0, 0, 220, 1188, "FD")

	byteImage, err := base64.StdEncoding.DecodeString(cvData.Photo)
	if err != nil {
		log.Fatal("error:", err)
	}

	byteImage, imageWidth, imageHeight := resizeImage(byteImage, 75, 75)

	imgData, err := gopdf.ImageHolderByBytes(byteImage)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	pdf.ImageByHolder(imgData, 72, 50, &gopdf.Rect{W: float64(imageWidth), H: float64(imageHeight)})

	err = pdf.SetFont("robotoBold", "", 16)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	pdf.SetTextColor(255, 255, 255)

	jobTitleLength, err := pdf.MeasureTextWidth(cvData.JobTitle)
	if err != nil {
		return nil, err
	}

	pdf.SetX((220 - jobTitleLength) / 2)
	pdf.SetY(140)
	pdf.Cell(nil, cvData.JobTitle)

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	nameSurnameLength, _ := pdf.MeasureTextWidth(cvData.NameSurname)
	pdf.SetX((220 - nameSurnameLength) / 2)
	pdf.SetY(160)
	pdf.Cell(nil, cvData.NameSurname)

	pdf.SetFont("robotoBold", "", 16)
	pdf.SetX(20)
	pdf.SetY(195)
	pdf.Cell(nil, "Personal Info")

	pdf.SetFillColor(200, 200, 200)
	pdf.RectFromUpperLeftWithStyle(20, 210, 180, 1, "F")

	err = pdf.SetFont("robotoBold", "", 14)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}

	pdf.SetX(20)
	pdf.SetY(pdf.GetY() + 25)
	pdf.Cell(nil, "Phone")

	pdf.SetFont("wts11", "", 10)

	pdf.SetY(pdf.GetY() + 15)
	pdf.SetX(20)
	pdf.Cell(nil, cvData.PhoneNumber)

	pdf.SetFont("robotoBold", "", 14)
	pdf.SetX(20)
	pdf.SetY(pdf.GetY() + 20)
	pdf.Cell(nil, "Birth Date")

	pdf.SetFont("wts11", "", 10)
	pdf.SetY(pdf.GetY() + 15)
	pdf.SetX(20)
	pdf.Cell(nil, user.BirthDate)

	pdf.SetFont("robotoBold", "", 14)
	pdf.SetX(20)
	pdf.SetY(pdf.GetY() + 20)
	pdf.Cell(nil, "E-mail")

	pdf.SetFont("wts11", "", 10)
	pdf.SetY(pdf.GetY() + 15)
	pdf.SetX(20)
	pdf.Cell(nil, cvData.Email)

	if cvData.Github != "" {
		pdf.SetFont("robotoBold", "", 14)
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 20)
		pdf.Cell(nil, "Github")

		pdf.SetFont("wts11", "", 10)
		pdf.SetY(pdf.GetY() + 15)
		pdf.SetX(20)
		pdf.Cell(nil, cvData.Github)
		githubWidth, _ := pdf.MeasureTextWidth(cvData.Github)
		pdf.AddExternalLink(cvData.Github, 20, pdf.GetY(), githubWidth, 14)
	}

	if cvData.Linkedin != "" {
		pdf.SetFont("robotoBold", "", 14)
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 20)
		pdf.Cell(nil, "Linkedin")

		pdf.SetFont("wts11", "", 10)
		pdf.SetY(pdf.GetY() + 15)
		pdf.SetX(20)
		pdf.Cell(nil, cvData.Linkedin)

		linkWidth, _ := pdf.MeasureTextWidth(cvData.Linkedin)
		pdf.AddExternalLink(cvData.Linkedin, 20, pdf.GetY(), linkWidth, 14)
	}

	if len(cvData.OtherSM) > 0 {
		pdf.SetFont("robotoBold", "", 14)
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 20)
		pdf.Cell(nil, "Other Links")

		pdf.SetFont("wts11", "", 10)

		for _, link := range cvData.OtherSM {
			pdf.SetY(pdf.GetY() + 15)
			pdf.SetX(20)
			pdf.Cell(nil, link)

			linkWidth, _ := pdf.MeasureTextWidth(link)
			pdf.AddExternalLink(link, 20, pdf.GetY(), linkWidth, 14)
		}
	}

	if len(cvData.Hobbies) > 0 {
		pdf.SetFont("robotoBold", "", 16)
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 25)
		pdf.Cell(nil, "Hobbies")

		pdf.SetFillColor(200, 200, 200)
		pdf.RectFromUpperLeftWithStyle(20, pdf.GetY()+15, 180, 1, "F")

		pdf.SetFont("wts11", "", 10)
		for _, hobby := range cvData.Hobbies {
			pdf.SetX(20)
			pdf.SetY(pdf.GetY() + 20)
			pdf.Cell(nil, hobby)
		}
	}

	if len(cvData.Languages) > 0 {
		pdf.SetFont("robotoBold", "", 16)
		pdf.SetX(20)
		pdf.SetY(pdf.GetY() + 25)
		pdf.Cell(nil, "Languages")

		pdf.SetFillColor(200, 200, 200)
		pdf.RectFromUpperLeftWithStyle(20, pdf.GetY()+15, 180, 1, "F")

		pdf.SetFont("wts11", "", 10)
		for _, language := range cvData.Languages {
			pdf.SetX(20)
			pdf.SetY(pdf.GetY() + 20)
			pdf.Cell(nil, language.Language+" - "+language.Level)
		}
	}

	pdf.SetX(250)
	pdf.SetY(50)

	if cvData.AboutMe != "" {
		pdf.SetTextColor(31, 31, 31)
		pdf.SetFont("robotoBold", "", 25)
		pdf.Cell(nil, "About Me")

		pdf.SetFillColor(31, 31, 31)
		pdf.SetY(pdf.GetY() + 27)
		pdf.RectFromUpperLeftWithStyle(250, pdf.GetY(), 500, 1, "F")

		pdf.SetFont("wts11", "", 17)
		pdf.SetY(pdf.GetY() + 15)
		splittedAboutMe, _ := pdf.SplitText(cvData.AboutMe, 500)
		for _, aboutMeLine := range splittedAboutMe {
			pdf.SetX(250)
			pdf.Cell(nil, aboutMeLine)
			pdf.Br(18)
		}
	}

	if len(cvData.Education) > 0 {
		pdf.SetTextColor(31, 31, 31)
		pdf.SetFont("robotoBold", "", 25)
		pdf.SetX(250)
		pdf.SetY(pdf.GetY() + 27)
		pdf.Cell(nil, "Education")

		pdf.SetFillColor(31, 31, 31)
		pdf.SetY(pdf.GetY() + 27)
		pdf.RectFromUpperLeftWithStyle(250, pdf.GetY(), 500, 1, "F")

		for _, education := range cvData.Education {
			pdf.SetTextColor(145, 145, 145)
			pdf.SetX(250)
			pdf.SetY(pdf.GetY() + 15)

			longestText := findLongerText(education.StartDate, education.EndDate, &pdf)
			pdf.SetFont("robotoBold", "", 17)
			pdf.Cell(nil, education.StartDate)

			pdf.SetTextColor(31, 31, 31)
			pdf.SetX(250 + longestText + 15)
			universityStartX := pdf.GetX()
			pdf.Cell(nil, education.School)

			pdf.SetTextColor(145, 145, 145)
			pdf.Br(20)
			pdf.SetX(250)
			pdf.Cell(nil, education.EndDate)

			pdf.SetTextColor(31, 31, 31)
			pdf.SetX(universityStartX)
			pdf.Cell(nil, education.Department)

			pdf.SetY(pdf.GetY() + 15)
		}
	}

	if len(cvData.Experience) > 0 {
		pdf.SetTextColor(31, 31, 31)
		pdf.SetFont("robotoBold", "", 25)
		pdf.SetX(250)
		pdf.SetY(pdf.GetY() + 27)
		pdf.Cell(nil, "Experience")

		pdf.SetFillColor(31, 31, 31)
		pdf.SetY(pdf.GetY() + 27)
		pdf.RectFromUpperLeftWithStyle(250, pdf.GetY(), 500, 1, "F")

		for _, experience := range cvData.Experience {
			pdf.SetTextColor(145, 145, 145)
			pdf.SetX(250)
			pdf.SetY(pdf.GetY() + 15)

			pdf.SetFont("robotoBold", "", 17)
			pdf.Cell(nil, experience.StartDate)

			pdf.SetTextColor(31, 31, 31)
			pdf.SetX(pdf.GetX() + 15)
			universityStartX := pdf.GetX()
			pdf.Cell(nil, experience.Company)

			pdf.SetTextColor(145, 145, 145)
			pdf.Br(20)
			pdf.SetX(250)
			pdf.Cell(nil, experience.EndDate)

			pdf.SetTextColor(31, 31, 31)
			pdf.SetX(universityStartX)
			pdf.Cell(nil, experience.Title)

			splittedDescription, _ := pdf.SplitText(experience.Description, 500)
			pdf.SetFont("wts11", "", 17)
			pdf.SetY(pdf.GetY() + 10)
			for _, descriptionLine := range splittedDescription {
				pdf.Br(17)
				pdf.SetX(250)
				pdf.Cell(nil, descriptionLine)
			}

			pdf.SetY(pdf.GetY() + 15)
		}
	}

	if len(cvData.Skills) > 0 {
		pdf.SetTextColor(31, 31, 31)
		pdf.SetFont("robotoBold", "", 25)
		pdf.SetX(250)
		pdf.SetY(pdf.GetY() + 27)
		pdf.Cell(nil, "Skills")

		pdf.SetFillColor(31, 31, 31)
		pdf.SetY(pdf.GetY() + 27)
		pdf.RectFromUpperLeftWithStyle(250, pdf.GetY(), 500, 1, "F")

		for _, skill := range cvData.Skills {
			pdf.Br(25)
			pdf.SetX(250)
			pdf.SetFont("robotoBold", "", 20)
			pdf.Cell(nil, skill)
		}
	}

	return pdf.GetBytesPdf(), nil
}

func resizeImage(byteImage []byte, widthCondition int, heightCondition int) ([]byte, int, int) {
	tempImage, _, _ := image.Decode(bytes.NewReader(byteImage))
	imageWidth := tempImage.Bounds().Max.X
	imageHeight := tempImage.Bounds().Max.Y

	if imageHeight > heightCondition {
		var resizedImage *image.NRGBA
		resizedImage = imaging.Resize(tempImage, 0, heightCondition, imaging.Lanczos)

		buf := new(bytes.Buffer)
		err := png.Encode(buf, resizedImage)

		if err != nil {
			//return nil nil, err
		}

		byteImage = buf.Bytes()

		tempImage, _, _ = image.Decode(bytes.NewReader(byteImage))
		imageWidth = tempImage.Bounds().Max.X
		imageHeight = tempImage.Bounds().Max.Y
	}
	if imageWidth > widthCondition {
		var resizedImage *image.NRGBA
		resizedImage = imaging.Resize(tempImage, widthCondition, 0, imaging.Lanczos)

		buf := new(bytes.Buffer)
		err := png.Encode(buf, resizedImage)

		if err != nil {
			//return nil nil, err
		}

		byteImage = buf.Bytes()

		tempImage, _, _ = image.Decode(bytes.NewReader(byteImage))
		imageWidth = tempImage.Bounds().Max.X
		imageHeight = tempImage.Bounds().Max.Y
	}

	return byteImage, imageWidth, imageHeight
}

func findLongerText(text1, text2 string, pdf *gopdf.GoPdf) float64 {
	text1Length, _ := pdf.MeasureTextWidth(text1)
	text2Length, _ := pdf.MeasureTextWidth(text2)

	if text1Length > text2Length {
		return text1Length
	}

	return text2Length
}
