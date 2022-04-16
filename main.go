package main

import (
	"cre-resume-backend/internal/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	userRepository := user.NewUserRepository("mongodb://localhost:27017")
	userService := user.NewUserService(userRepository)
	userController := user.NewUserHandler(userService)

	userController.SetupUserHandler(app)

	app.Listen(":8080")
}
