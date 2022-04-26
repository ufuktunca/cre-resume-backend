package main

import (
	"cre-resume-backend/internal/cv"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepository := user.NewUserRepository("mongodb://localhost:27017")
	userService := user.NewUserService(userRepository)
	userController := user.NewUserHandler(userService)

	jobPostRepository := jobPost.NewJobRepository("mongodb://localhost:27017")
	jobPostService := jobPost.NewJobPostService(jobPostRepository, userRepository)
	jobPostHandler := jobPost.NewJobPostHandler(jobPostService)

	cvRepository := cv.CreateCVRepository("mongodb://localhost:27017")
	cvService := cv.NewCVService(cvRepository)
	cvHandler := cv.NewCVHandler(cvService)

	userController.SetupUserHandler(app)
	jobPostHandler.SetupJobPostHandler(app)
	cvHandler.SetupRouteApp(app)

	app.Listen(":8080")
}
