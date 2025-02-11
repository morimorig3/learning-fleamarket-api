package services

import (
	"learning-fleamarket-api/dto"
	"learning-fleamarket-api/models"
	"learning-fleamarket-api/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemId uint, userId uint) (*models.Item, error)
	Create(createItemInput dto.CreateItemInput, userId uint) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput, userId uint) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

// 戻り値の型をインターフェースにしているので、具体的な実装（ItemService）に
// FindAllの実装を忘れているとコンパイルエラーになる
func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{
		repository: repository,
	}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindByID(itemId uint, userId uint) (*models.Item, error) {
	return s.repository.FindById(itemId, userId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput, userId uint) (*models.Item, error) {
	newItem := models.Item{
		UserID:      userId,
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput, userId uint) (*models.Item, error) {
	targetItem, err := s.FindByID(itemId, userId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}
	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint, userId uint) error {
	return s.repository.Delete(itemId, userId)
}
