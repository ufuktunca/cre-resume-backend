package auth

import (
	"cre-resume-backend/internal/models"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var jwtKey = []byte("sdfk1lmhd2342sklgfjdhas634flkdshj23oır42o3euıw")

func VerifyToken(c *fiber.Ctx) error {

	claims := &models.Claims{}
	auth := c.Get("Authorization")
	fmt.Println(auth)
	tkns, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	fmt.Println(err)
	if err != nil || !tkns.Valid {
		c.Status(fiber.StatusUnauthorized)
		return nil
	}
	c.Request().Header.Set("user-id", claims.UserID)
	c.Next()

	return nil
}

func CreateToken(id string) (*string, error) {

	expirationTime := time.Now().Add(12 * time.Hour)
	claims := models.Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
