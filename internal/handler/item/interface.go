package item

import "github.com/gofiber/fiber/v2"

type ItemHandler interface {
	RegisterRoutes(app *fiber.App)
	GetItemByID(c *fiber.Ctx) error
	UpdateNormalItem(c *fiber.Ctx) error
	UpdatePremiumItem(c *fiber.Ctx) error
}
