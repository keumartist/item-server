package item

import (
	"art-item/internal/outbound"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler interface {
	RegisterRoutes(app *fiber.App, authClient outbound.AuthServiceClient)
	GetItemByID(c *fiber.Ctx) error
	UpdateNormalItem(c *fiber.Ctx) error
	UpdatePremiumItem(c *fiber.Ctx) error
	CreateItem(c *fiber.Ctx) error
}
