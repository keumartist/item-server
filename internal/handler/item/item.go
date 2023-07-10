package item

import (
	"errors"
	"log"
	"net/http"

	customerror "art-item/internal/error"
	middleware "art-item/internal/middleware"
	"art-item/internal/outbound"
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

func (h *ItemHandlerImpl) RegisterRoutes(app *fiber.App, authClient outbound.AuthServiceClient) {
	app.Get("/api/v1/items/:id", h.GetItemByID)
	app.Post("/api/v1/items", middleware.AuthMiddleware(authClient), h.CreateItem)
	app.Put("/api/v1/items/normal", middleware.AuthMiddleware(authClient), h.UpdateNormalItem)
	app.Put("/api/v1/items/premium", middleware.AuthMiddleware(authClient), h.UpdatePremiumItem)
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
	var input struct {
		NewItem map[string]interface{} `json:"new_item"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(customerror.ErrBadRequest)
	}

	userID := c.Locals("userID").(string)

	if err := h.service.UpdateNormalItem(service.UpdateNormalItemInput{UserID: userID, NewItem: input.NewItem}); err != nil {
		if errors.Is(err, customerror.ErrItemNotFound) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrItemNotFound)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
}

func (h *ItemHandlerImpl) UpdatePremiumItem(c *fiber.Ctx) error {
	var input struct {
		NewItem map[string]interface{} `json:"new_item"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(customerror.ErrBadRequest)
	}

	userID := c.Locals("userID").(string)
	if err := h.service.UpdatePremiumItem(service.UpdatePremiumItemInput{UserID: userID, NewItem: input.NewItem}); err != nil {
		if errors.Is(err, customerror.ErrItemNotFound) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrItemNotFound)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	log.Println("업뎃됭나")

	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Item updated successfully"})
}

func (h *ItemHandlerImpl) CreateItem(c *fiber.Ctx) error {
	var input struct {
		NormalItem  map[string]interface{} `json:"normal_item"`
		PremiumItem map[string]interface{} `json:"premium_item"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(customerror.ErrBadRequest)
	}

	userID := c.Locals("userID").(string)

	newItem, err := h.service.CreateItem(service.CreateItemInput{UserID: userID, NormalItem: input.NormalItem, PremiumItem: input.PremiumItem})
	if err != nil {
		if errors.Is(err, customerror.ErrDuplicatedUserItem) {
			return c.Status(http.StatusBadRequest).JSON(customerror.ErrDuplicatedUserItem)
		}
		return c.Status(http.StatusInternalServerError).JSON(customerror.ErrInternal)
	}

	return c.Status(http.StatusCreated).JSON(newItem)
}
