package user

import (
	"cre-resume-backend/internal/models"

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
	app.Post("/reSend", u.ReSendController)
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
		if err == models.EmailError {
			c.Status(fiber.StatusBadRequest)
			return nil
		}
		c.Status(fiber.StatusInternalServerError)
		return nil
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

	loginType := c.Query("type", "")

	token, err := u.View.Login(login, loginType)
	if err != nil {
		if err == models.ActivationError {
			c.Status(fiber.StatusUnauthorized)
			return nil
		}
		return err
	}

	c.JSON(models.Auth{
		Token: *token,
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

func (u *User) ReSendController(c *fiber.Ctx) error {
	reSend := models.ReSend{}

	err := c.BodyParser(&reSend)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}
	err = u.View.ReSend(reSend.Email)

	return err
}
