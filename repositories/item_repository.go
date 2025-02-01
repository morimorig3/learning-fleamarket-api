package repositories

import (
	"errors"
	"learning-freemarket-api/models"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
}

// ↑↑↑↑↑のFindAllを満たすための実装が↓↓↓↓↓↓
type ItemMemoryRepository struct {
	items []models.Item
}

// ファクトリ関数の戻り値の方をインターフェースにすることで
// 具体的な実装がインターフェースの型を満たしていないときに
// エラーが発生するので実装もれを防ぐことができる
func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint) (*models.Item, error) {
	for _, v := range r.items {
		if v.ID == itemId {
			return &v, nil
		}
	}
	return nil, errors.New("Item not found")
}
