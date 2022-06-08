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

	userModel := user.NewUserModel("mongodb+srv://dbUser:EMbUEVTOraH7YfpM@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	userView := user.NewUserView(userModel)
	userController := user.NewUserController(userView)

	cvModel := cv.CreateCVModel("mongodb+srv://dbUser:EMbUEVTOraH7YfpM@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	cvView := cv.NewCVView(cvModel, userModel)
	cvController := cv.NewCVController(cvView)

	jobPostModel := jobPost.NewJobModel("mongodb+srv://dbUser:EMbUEVTOraH7YfpM@cluster0.fg8ftbg.mongodb.net/?retryWrites=true&w=majority")
	jobPostView := jobPost.NewJobPostView(jobPostModel, userModel, cvModel)
	jobPostController := jobPost.NewJobPostController(jobPostView)

	userController.SetupUserController(app)
	cvController.SetupRouteApp(app)
	jobPostController.SetupJobPostController(app)

	app.Listen(":8080")
}
