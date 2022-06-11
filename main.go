package main

import (
	"cre-resume-backend/internal/cv"
	jobPost "cre-resume-backend/internal/job-post"
	"cre-resume-backend/internal/user"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	userModel := user.NewUserModel("mongodb+srv://dbUser:X9ZE1PrfJfBg7T4m@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	userView := user.NewUserView(userModel)
	userController := user.NewUserController(userView)

	cvModel := cv.CreateCVModel("mongodb+srv://dbUser:X9ZE1PrfJfBg7T4m@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	cvView := cv.NewCVView(cvModel, userModel)
	cvController := cv.NewCVController(cvView)

	jobPostModel := jobPost.NewJobModel("mongodb+srv://dbUser:X9ZE1PrfJfBg7T4m@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	jobPostView := jobPost.NewJobPostView(jobPostModel, userModel, cvModel)
	jobPostController := jobPost.NewJobPostController(jobPostView)

	userController.SetupUserController(app)
	cvController.SetupRouteApp(app)
	jobPostController.SetupJobPostController(app)

	app.Listen(":" + SetPort())
}

func SetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return "8080"
	}

	return port
}
