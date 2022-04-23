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
	app.Get("/jobPost/:type", j.GetJobPosts)
	app.Get("/jobPost/:jobId/apply", j.ApplyJobHandler)

}

func (j *JobPostHandler) CreateJobPost(c *fiber.Ctx) error {
	jobPost := &models.JobPost{}
	ownerEmail := c.Get("user-email", "")
	if ownerEmail == "" {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err := c.BodyParser(jobPost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	jobPostData, err := j.Service.CreateJobPost(jobPost, ownerEmail)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPostData)
	c.Status(fiber.StatusCreated)
	return nil
}

func (j *JobPostHandler) GetJobPosts(c *fiber.Ctx) error {
	jobPostType := c.Params("type")
	category := c.Query("category", "")
	from := c.Query("from", "")
	to := c.Query("to", "")
	sort := c.Query("sort", "")

	jobPosts, err := j.Service.GetJobPosts(jobPostType, category, from, to, sort)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPosts)
	return nil
}

func (j *JobPostHandler) ApplyJobHandler(c *fiber.Ctx) error {
	return nil
}
