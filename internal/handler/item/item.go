package item

import (
	"errors"
	"net/http"

	customerror "art-item/internal/error"
	middleware "art-item/internal/middleware"
	service "art-item/internal/service/item"

	"github.com/gofiber/fiber/v2"
)

type ItemHandlerImpl struct {
	service service.ItemService
}

func NewItemHandler(service service.ItemService) ItemHandler {
	return &ItemHandlerImpl{
		service: service,
	}
}

func (h *ItemHandlerImpl) RegisterRoutes(app *fiber.App) {
	app.Get("/items/:id", h.GetItemByID)
	app.Put("/items/normal", middleware.JWTMiddleware(), h.UpdateNormalItem)
	app.Put("/items/premium", middleware.JWTMiddleware(), h.UpdatePremiumItem)
}

func (h *ItemHandlerImpl) GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")

	item, err := h.service.GetItemByID(service.GetItemByIDInput{ID: id})

	if err != nil {
		if errors.Is(err, customerror.ErrItemNotFound) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrItemNotFound)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	return c.JSON(item)
}

func (h *ItemHandlerImpl) UpdateNormalItem(c *fiber.Ctx) error {
	var newItem map[string]interface{}

	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(customerror.ErrBadRequest)
	}

	userID := c.Locals("userID").(string)

	if err := h.service.UpdateNormalItem(service.UpdateNormalItemInput{UserID: userID, NewItem: newItem}); err != nil {
		if errors.Is(err, customerror.ErrItemNotFound) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrItemNotFound)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
}

func (h *ItemHandlerImpl) UpdatePremiumItem(c *fiber.Ctx) error {
	var newItem map[string]interface{}

	if err := c.BodyParser(&newItem); err != nil {
		return c.Status(http.StatusBadRequest).JSON(customerror.ErrBadRequest)
	}

	userID := c.Locals("userID").(string)
	if err := h.service.UpdatePremiumItem(service.UpdatePremiumItemInput{UserID: userID, NewItem: newItem}); err != nil {
		if errors.Is(err, customerror.ErrItemNotFound) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrItemNotFound)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
}
