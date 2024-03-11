package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	IItemsRepo interface {
		GetItems(categoryID uint64) ([]model.Item, error)
		GetArchivedItems(boardID uint64) ([]model.Item, error)
		CreateItem(item *model.Item) error
		UpdateItem(item *model.Item) error
		ArchiveItem(itemID uint64) error
		UnArchiveItem(itemID uint64) error
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
	err := ir.db.Model(&[]model.Item{}).Where("category_id = ? AND archived = ?", categoryID, false).Find(&items).Error
	return items, err
}

func (ir *ItemsRepo) GetArchivedItems(boardID uint64) ([]model.Item, error) {
	var items []model.Item
	err := ir.db.Model(&[]model.Item{}).Joins("JOIN categories ON items.category_id = categories.id").
        Where("categories.board_id = ? AND items.archived = ?", boardID, true).
        Find(&items).Error
		
	return items, err
}

func (ir *ItemsRepo) CreateItem(item *model.Item) error {
	err := ir.db.Model(&model.Item{}).Create(&item).Error
	return err
}

func (ir *ItemsRepo) UpdateItem(item *model.Item) error {
	err := ir.db.Model(&model.Item{}).Where("id = ?", item.ID).Save(&item).Error
	return err
}

func (ir *ItemsRepo) ArchiveItem(itemID uint64) error {
	err := ir.db.Model(&model.Item{}).Where("id = ?", itemID).Update("archived", true).Error
	return err
}

func (ir *ItemsRepo) UnArchiveItem(itemID uint64) error {
	err := ir.db.Model(&model.Item{}).Where("id = ?", itemID).Update("archived", false).Error
	return err
}