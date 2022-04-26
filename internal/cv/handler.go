package cv

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type CVHandler struct {
	Service CVServiceInterface
}

func NewCVHandler(service CVServiceInterface) *CVHandler {
	return &CVHandler{
		Service: service,
	}
}

func (cv *CVHandler) SetupRouteApp(app *fiber.App) {
	app.Use(auth.VerifyToken)
	app.Post("/cv", cv.CreateCV)
}

func (cv *CVHandler) CreateCV(c *fiber.Ctx) error {
	cvData := models.CV{}
	err := c.BodyParser(&cvData)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	ownerEmail := c.Get("user-email", "")

	cv.Service.CreateCV(&cvData, ownerEmail)

	return nil
}
