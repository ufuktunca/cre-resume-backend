package helpers

import (
	"bytes"
	"cre-resume-backend/internal/models"
	"encoding/base64"
	"fmt"
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

func CreateCV(cvData *models.CV) {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 840, H: 1188}})

	err := pdf.AddTTFFont("wts11", "./Roboto-Regular.ttf")
	if err != nil {
		fmt.Println(err)
	}

	err = pdf.AddTTFFont("robotoBold", "./Roboto-Bold.ttf")
	if err != nil {
		fmt.Println(err)
	}

	pdf.AddPage()

	pdf.SetX(500)
	pdf.SetY(140)

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.Cell(nil, "test")
	if err != nil {
		fmt.Println(err)
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
		return
	}

	pdf.ImageByHolder(imgData, 72, 50, &gopdf.Rect{W: float64(imageWidth), H: float64(imageHeight)})

	err = pdf.SetFont("robotoBold", "", 16)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetTextColor(255, 255, 255)

	jobTitleLength, err := pdf.MeasureTextWidth(cvData.JobTitle)
	if err != nil {
		return
	}

	pdf.SetX((220 - jobTitleLength) / 2)
	pdf.SetY(140)
	pdf.Cell(nil, cvData.JobTitle)

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
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
		return
	}

	pdf.SetX(20)
	pdf.SetY(220)
	pdf.Cell(nil, "Phone")

	pdf.SetFont("wts11", "", 14)

	pdf.SetY(235)
	pdf.SetX(20)
	pdf.Cell(nil, cvData.PhoneNumber)

	pdf.SetFont("robotoBold", "", 14)
	pdf.SetX(20)
	pdf.SetY(255)
	pdf.Cell(nil, "E-mail")

	pdf.SetFont("wts11", "", 14)
	pdf.SetY(270)
	pdf.SetX(20)
	pdf.Cell(nil, cvData.Email)

	pdf.SetFont("robotoBold", "", 14)
	pdf.SetX(20)
	pdf.SetY(290)
	pdf.Cell(nil, "Github")

	pdf.SetFont("wts11", "", 14)
	pdf.SetY(305)
	pdf.SetX(20)
	pdf.Cell(nil, cvData.Github)

	err = pdf.WritePdf("./test.pdf")
	if err != nil {
		fmt.Println(err)
	}
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
			//return nil, err
		}

		byteImage = buf.Bytes()

		tempImage, _, _ = image.Decode(bytes.NewReader(byteImage))
		imageWidth = tempImage.Bounds().Max.X
		imageHeight = tempImage.Bounds().Max.Y
	} else if imageWidth > widthCondition {
		var resizedImage *image.NRGBA
		resizedImage = imaging.Resize(tempImage, widthCondition, 0, imaging.Lanczos)

		buf := new(bytes.Buffer)
		err := png.Encode(buf, resizedImage)

		if err != nil {
			//return nil, err
		}

		byteImage = buf.Bytes()

		tempImage, _, _ = image.Decode(bytes.NewReader(byteImage))
		imageWidth = tempImage.Bounds().Max.X
		imageHeight = tempImage.Bounds().Max.Y
	}

	return byteImage, imageWidth, imageHeight
}
