package dto

import (
	"art-item/internal/domain/item"
	"time"
)

type Item struct {
	ID          string                 `json:"id"`
	NormalItem  map[string]interface{} `json:"normal_item"`
	PremiumItem map[string]interface{} `json:"premium_item"`
	CreatedAt   time.Time              `json:"created_at"`
	UpdatedAt   time.Time              `json:"updated_at"`
}

func ItemDomainToDto(item *item.Item) Item {
	return Item{
		ID:          item.ID.Hex(),
		NormalItem:  item.NormalItem,
		PremiumItem: item.PremiumItem,
		CreatedAt:   item.CreatedAt,
		UpdatedAt:   item.UpdatedAt,
	}
}
