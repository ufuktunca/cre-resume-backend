package main

import (
	"cre-resume-backend/internal/cv"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userModel := user.NewUserModel("mongodb://localhost:27017")
	userView := user.NewUserView(userModel)
	userController := user.NewUserController(userView)

	jobPostModel := jobPost.NewJobModel("mongodb://localhost:27017")
	jobPostView := jobPost.NewJobPostView(jobPostModel, userModel)
	jobPostController := jobPost.NewJobPostController(jobPostView)

	cvModel := cv.CreateCVModel("mongodb://localhost:27017")
	cvView := cv.NewCVView(cvModel)
	cvController := cv.NewCVController(cvView)

	userController.SetupUserController(app)
	jobPostController.SetupJobPostController(app)
	cvController.SetupRouteApp(app)

	app.Listen(":8080")
}
