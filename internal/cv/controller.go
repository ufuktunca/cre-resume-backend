package cv

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CVController struct {
	View CVViewInterface
}

func NewCVController(view CVViewInterface) *CVController {
	return &CVController{
		View: view,
	}
}

func (cv *CVController) SetupRouteApp(app *fiber.App) {
	app.Use(auth.VerifyToken)
	app.Post("/cv", cv.CreateCV)
	app.Get("/cv/:cvId", cv.GetCV)
	app.Get("/user/cv", cv.GetCVHandler)
}

func (cv *CVController) CreateCV(c *fiber.Ctx) error {
	cvData := models.CV{}
	err := c.BodyParser(&cvData)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	userID := c.Get("user-id", "")

	err = cv.View.CreateCV(&cvData, userID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	return nil
}

func (cv *CVController) GetCVHandler(c *fiber.Ctx) error {
	userID := c.Get("user-id", "")

	CVs, err := cv.View.GetCVs(userID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(CVs)

	return nil
}

func (cv *CVController) GetCV(c *fiber.Ctx) error {
	cvID := c.Params("cvId", "")

	pdf, err := cv.View.GetCV(cvID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.Set("Content-Type", "application/pdf")
	c.Write(pdf)

	return nil
}
