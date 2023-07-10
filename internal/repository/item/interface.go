package item

import "art-item/internal/domain/item"

type ItemRepository interface {
	FindByID(id string) (*item.Item, error)
	FindByUserID(userID string) (*item.Item, error)
	UpdateNormalItem(id string, newItem map[string]interface{}) error
	UpdatePremiumItem(id string, newItem map[string]interface{}) error
	CreateItem(userID string, normalItem map[string]interface{}, premiumItem map[string]interface{}) (*item.Item, error)
}
