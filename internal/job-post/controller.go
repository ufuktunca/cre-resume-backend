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
	app.Use("/jobPost", auth.VerifyToken)
	app.Use("/user", auth.VerifyToken)
	app.Post("/jobPost/:userType", j.CreateJobPost)
	app.Get("/jobPost/:type", j.GetJobPosts)
	app.Post("/jobPost/:jobId/apply", j.ApplyJobController)
	app.Get("/user/jobPost/apply", j.GetAppliedJobs)
	app.Get("/jobPost/user/:type", j.GetUserJobPosts)
	app.Get("/jobs/:jobId/apply", j.GetJobApplies)
	app.Delete("/jobs/:jobId", j.DeleteJobPostHandler)

}

func (j *JobPostController) CreateJobPost(c *fiber.Ctx) error {
	jobPost := &models.JobPost{}
	userID := c.Get("user-id", "")
	jobPostType := c.Params("userType", "")

	if userID == "" {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err := c.BodyParser(jobPost)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	jobPost.Type = jobPostType
	jobPostData, err := j.View.CreateJobPost(jobPost, userID)
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
	userID := c.Get("user-id", "")
	jobID := c.Params("jobId")

	applyJobDTO := models.ApplyJobPostDTO{}
	err := c.BodyParser(&applyJobDTO)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err = j.View.ApplyJobPost(&applyJobDTO, userID, jobID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	return nil
}

func (j *JobPostController) GetUserJobPosts(c *fiber.Ctx) error {
	postType := c.Params("type")
	userID := c.Get("user-id", "")
	category := c.Query("category", "")
	from := c.Query("from", "")
	to := c.Query("to", "")
	sort := c.Query("sort", "")

	jobPosts, err := j.View.GetUserJobPosts(userID, postType, category, from, to, sort)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.JSON(jobPosts)
	return nil
}

func (j *JobPostController) GetAppliedJobs(c *fiber.Ctx) error {
	userID := c.Get("user-id", "")

	appliedJobs, err := j.View.GetUserAppliedJobs(userID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.JSON(appliedJobs)
	return nil
}

func (j *JobPostController) GetJobApplies(c *fiber.Ctx) error {
	jobId := c.Params("jobId", "")
	jobApplies, err := j.View.GetJobApplies(jobId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.JSON(jobApplies)
	return nil
}

func (j *JobPostController) DeleteJobPostHandler(c *fiber.Ctx) error {
	jobId := c.Params("jobId", "")

	err := j.View.DeleteJobPost(jobId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	c.Status(fiber.StatusNoContent)
	return nil
}
