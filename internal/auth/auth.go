package auth

import (
	"cre-resume-backend/internal/models"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("sdfk1lmhd2342sklgfjdhas634flkdshj23oır42o3euıw")

func VerifyToken(c *fiber.Ctx) error {

	claims := &models.Claims{}
	auth := c.Cookies("auth")
	fmt.Println(auth)
	tkns, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !tkns.Valid {
		return err
	}

	fmt.Println(claims)

	c.Next()

	return nil
}
