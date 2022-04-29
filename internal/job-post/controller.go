package jobPost

import (
	"cre-resume-backend/internal/auth"
	"cre-resume-backend/internal/models"

	"github.com/gofiber/fiber/v2"
)

type JobPostController struct {
	View JobPostViewInterface
}

func NewJobPostController(view JobPostViewInterface) *JobPostController {
	return &JobPostController{View: view}
}

func (j *JobPostController) SetupJobPostController(app *fiber.App) {
	app.Use(auth.VerifyToken)
	app.Post("/jobPost/:userType", j.CreateJobPost)
	app.Get("/jobPost/:type", j.GetJobPosts)
	app.Post("/jobPost/:jobId/apply", j.ApplyJobController)
	app.Get("/user/jobPost/:type", j.GetUserJobPosts)

}

func (j *JobPostController) CreateJobPost(c *fiber.Ctx) error {
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

	jobPostData, err := j.View.CreateJobPost(jobPost, ownerEmail)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPostData)
	c.Status(fiber.StatusCreated)
	return nil
}

func (j *JobPostController) GetJobPosts(c *fiber.Ctx) error {
	jobPostType := c.Params("type")
	category := c.Query("category", "")
	from := c.Query("from", "")
	to := c.Query("to", "")
	sort := c.Query("sort", "")

	jobPosts, err := j.View.GetJobPosts(jobPostType, category, from, to, sort)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPosts)
	return nil
}

func (j *JobPostController) ApplyJobController(c *fiber.Ctx) error {
	ownerEmail := c.Get("user-email", "")
	jobID := c.Params("jobId")

	applyJobDTO := models.ApplyJobPostDTO{}
	err := c.BodyParser(&applyJobDTO)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err = j.View.ApplyJobPost(&applyJobDTO, ownerEmail, jobID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	return nil
}

func (j *JobPostController) GetUserJobPosts(c *fiber.Ctx) error {
	postType := c.Params("type")
	userEmail := c.Get("user-email", "")

	jobPosts, err := j.View.GetUserJobPosts(userEmail, postType)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPosts)
	return nil
}
