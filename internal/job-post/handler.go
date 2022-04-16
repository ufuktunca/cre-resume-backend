package jobPost

import (
	"cre-resume-backend/internal/auth"

	"github.com/gofiber/fiber/v2"
)

type JobPostHandler struct {
	Service JobPostServiceInterface
}

func SetupJobPostHandler(app *fiber.App) {
	app.Use(auth.VerifyToken)

}
