package user

import (
	"cre-resume-backend/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	View UserViewInterface
}

func NewUserController(userView UserViewInterface) *User {
	return &User{
		View: userView,
	}
}

func (u *User) SetupUserController(app *fiber.App) {
	app.Post("/login", u.LoginUserController)
	app.Post("/register", u.RegisterUserController)
	app.Get("/activation", u.ActivateUserController)

}

func (u *User) RegisterUserController(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	err = u.View.Register(&user)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	c.Status(fiber.StatusCreated)
	return nil
}

func (u *User) LoginUserController(c *fiber.Ctx) error {
	login := &models.Login{}

	err := c.BodyParser(&login)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return err
	}

	token, err := u.View.Login(login)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    "auth",
		Value:   *token,
		Expires: time.Now().Add(12 * time.Hour),
	})

	return nil
}

func (u *User) ActivateUserController(c *fiber.Ctx) error {
	userID := c.Query("userID")

	err := u.View.ActivateUser(userID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	return nil
}