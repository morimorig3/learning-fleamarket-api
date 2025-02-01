package services

import (
	"learning-freemarket-api/models"
	"learning-freemarket-api/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindByID(itemId uint) (*models.Item, error)
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

func (s *ItemService) FindByID(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}
