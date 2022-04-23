package main

import (
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepository := user.NewUserRepository("mongodb://localhost:27017")
	userService := user.NewUserService(userRepository)
	userController := user.NewUserHandler(userService)

	jobPostRepository := jobPost.CreateJobPostRepository("mongodb://localhost:27017")
	jobPostService := jobPost.NewJobPostService(jobPostRepository)
	jobPostHandler := jobPost.NewJobPostHandler(jobPostService)

	userController.SetupUserHandler(app)
	jobPostHandler.SetupJobPostHandler(app)

	app.Listen(":8080")
}
