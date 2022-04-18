package jobPost

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type JobPostHandler struct {
	Service JobPostServiceInterface
}

func NewJobPostHandler(service JobPostServiceInterface) *JobPostHandler {
	return &JobPostHandler{Service: service}
}

func (j *JobPostHandler) SetupJobPostHandler(app *fiber.App) {
	app.Use(auth.VerifyToken)
	app.Post("/jobPost/:userType", j.CreateJobPost)

}

func (j *JobPostHandler) CreateJobPost(c *fiber.Ctx) error {
	jobPost := &models.JobPost{}

	err := c.BodyParser(jobPost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	jobPostData, err := j.Service.CreateJobPost(jobPost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPostData)
	c.Status(fiber.StatusCreated)
	return nil
}
