package middleware

import (
	"net/http"
	"strings"

	customerror "art-item/internal/error"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		tokenString := splitToken[1]
		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		token, _ := jwt.Parse(tokenString, nil)

		if token == nil {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := claims["sub"].(string)
			c.Locals("userID", userID)
			return c.Next()
		} else {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}
	}
}
