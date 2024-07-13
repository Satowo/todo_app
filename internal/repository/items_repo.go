package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	IItemsRepo interface {
		GetItems(categoryID uint64) ([]model.Item, error)
		CreateItem(item *model.Item) error
		UpdateItem(itemID, categoryID uint64, item_title, content, expired_at string) error
		DeleteItem(itemID uint64) error
	}
	ItemsRepo struct {
		db *gorm.DB
	}
)

func NewItemsRepo(db *gorm.DB) *ItemsRepo {
	return &ItemsRepo{db}
}

func (ir *ItemsRepo) GetItems(categoryID uint64) ([]model.Item, error) {
	var items []model.Item
	err := ir.db.Model(&[]model.Item{}).Where("category_id = ? AND deleted = ?", categoryID, false).Find(&items).Error
	return items, err
}

func (ir *ItemsRepo) CreateItem(item *model.Item) error {
	err := ir.db.Model(&model.Item{}).Create(&item).Error
	return err
}

func (ir *ItemsRepo) UpdateItem(itemID, categoryID uint64, item_title, content, expired_at string) error {
	err := ir.db.Model(&model.Item{}).Where("id = ?", itemID).Update("category_id", categoryID).Update("title", item_title).Update("content", content).Update("expired_at", expired_at).Error
	return err
}

func (ir *ItemsRepo) DeleteItem(itemID uint64) error {
	err := ir.db.Model(&model.Item{}).Where("id = ?", itemID).Update("deleted", true).Error
	return err
}
