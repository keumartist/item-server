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

	return s.repo.UpdateNormalItem(input.UserID, input.NewItem)
}

func (s *ItemServiceImpl) UpdatePremiumItem(input UpdatePremiumItemInput) error {
	item, err := s.repo.FindByUserID(input.UserID)
	if err != nil {
		return customerror.ErrInternal
	}

	if item == nil {
		return customerror.ErrItemNotFound
	}

	return s.repo.UpdatePremiumItem(input.UserID, input.NewItem)
}
