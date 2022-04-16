package user

import (
	"cre-resume-backend/internal/models"

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
	app.Post("/login", u.LoginUserHandler)
}

func (u *User) RegisterUserHandler(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
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

func (u *User) LoginUserHandler(c *fiber.Ctx) error {
	login := &models.Login{}

	err := c.BodyParser(&login)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	token, err := u.Service.Login(login)
	if err != nil {
		return err
	}

	c.JSON(models.Auth{
		Token: *token,
	})

	return nil
}
