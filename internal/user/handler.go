package user

import (
	"cre-resume-backend/internal/user/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	Service UserServiceInterface
}

func NewUserHandler(userService UserServiceInterface) *User {
	return &User{
		Service: userService,
	}
}

func (u *User) SetupUserHandler(app *fiber.App) {
	app.Post("/register", u.RegisterUserHandler)
}

func (u *User) RegisterUserHandler(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		fmt.Println(err, "@@@")
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err = u.Service.Register(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	c.Status(fiber.StatusCreated)
	return nil
}
