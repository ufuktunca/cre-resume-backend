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
}

func (cv *CVController) CreateCV(c *fiber.Ctx) error {
	cvData := models.CV{}
	err := c.BodyParser(&cvData)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	ownerEmail := c.Get("user-email", "")

	cv.View.CreateCV(&cvData, ownerEmail)

	return nil
}
