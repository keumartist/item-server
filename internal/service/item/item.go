package item

import (
	dto "art-item/internal/dto/item"
	customerror "art-item/internal/error"
	repository "art-item/internal/repository/item"
)

type ItemServiceImpl struct {
	repo repository.ItemRepository
}

func NewItemService(repo repository.ItemRepository) ItemService {
	return &ItemServiceImpl{repo: repo}
}

func (s *ItemServiceImpl) GetItemByID(input GetItemByIDInput) (dto.Item, error) {
	item, err := s.repo.FindByID(input.ID)
	if err != nil {
		return dto.Item{}, customerror.ErrInternal
	}

	if item == nil {
		return dto.Item{}, customerror.ErrItemNotFound
	}

	return dto.ItemDomainToDto(item), nil
}

func (s *ItemServiceImpl) GetItemByUserID(input GetItemByUserIDInput) (dto.Item, error) {
	item, err := s.repo.FindByUserID(input.UserID)
	if err != nil {
		return dto.Item{}, err
	}

	if item == nil {
		return dto.Item{}, customerror.ErrItemNotFound
	}

	return dto.ItemDomainToDto(item), nil
}

func (s *ItemServiceImpl) UpdateNormalItem(input UpdateNormalItemInput) error {
	item, err := s.repo.FindByUserID(input.UserID)
	if err != nil {
		return customerror.ErrInternal
	}

	if item == nil {
		return customerror.ErrItemNotFound
	}

	return s.repo.UpdateNormalItem(item.ID.Hex(), input.NewItem)
}

func (s *ItemServiceImpl) UpdatePremiumItem(input UpdatePremiumItemInput) error {
	item, err := s.repo.FindByUserID(input.UserID)
	if err != nil {
		return customerror.ErrInternal
	}

	if item == nil {
		return customerror.ErrItemNotFound
	}

	return s.repo.UpdatePremiumItem(item.ID.Hex(), input.NewItem)
}

func (s *ItemServiceImpl) CreateItem(input CreateItemInput) (dto.Item, error) {
	existingItem, _ := s.repo.FindByUserID(input.UserID)
	if existingItem != nil {
		return dto.Item{}, customerror.ErrDuplicatedUserItem
	}

	newItem, err := s.repo.CreateItem(input.UserID, input.NormalItem, input.PremiumItem)
	if err != nil {
		return dto.Item{}, err
	}

	return dto.ItemDomainToDto(newItem), nil
}
