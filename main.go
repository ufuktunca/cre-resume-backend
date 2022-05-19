package main

import (
	"cre-resume-backend/internal/cv"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

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
	cvController.SetupRouteApp(app)
	jobPostController.SetupJobPostController(app)

	app.Listen(":8080")
}
