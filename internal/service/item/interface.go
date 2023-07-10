package item

import dto "art-item/internal/dto/item"

type ItemService interface {
	GetItemByID(input GetItemByIDInput) (dto.Item, error)
	UpdateNormalItem(input UpdateNormalItemInput) error
	UpdatePremiumItem(input UpdatePremiumItemInput) error
}
