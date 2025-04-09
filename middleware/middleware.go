package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateMiddleware(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing token",
			})
	}
	token, err := jwt.Parse(tokenString, func ()token *jwt.Token) (interface{}, error)  {
		if _, ok:= token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err!=nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Expired token",
			})
	}

	claims := token.Claims.(jwt.MapClaims)
	c.Locals("user", claims)
	return c.Next()

}