package middleware

import (
	"context"
	"net/http"
	"strings"

	customerror "art-item/internal/error"
	"art-item/internal/outbound"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(authClient outbound.AuthServiceClient) fiber.Handler {
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

		ctx := context.Background()
		res, err := authClient.VerifyToken(ctx, tokenString, "access")
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		if !res.IsValid {
			return c.Status(http.StatusUnauthorized).JSON(customerror.ErrUnauthorized)
		}

		c.Locals("userID", res.UserId)

		return c.Next()
	}
}
