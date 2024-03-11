package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	ICategoriesRepo interface {
		GetCategories(boardID uint64) ([]model.Category, error)
		CreateCategory(category *model.Category) error
		UpdateCategory(category *model.Category) error
		DeleteCategory(category *model.Category) error
	}
	CategoriesRepo struct {
		db *gorm.DB
	}
)

func NewCategoriesRepo(db *gorm.DB) *CategoriesRepo {
	return &CategoriesRepo{db}
}

func (br *CategoriesRepo) GetCategories(boardID uint64) ([]model.Category, error) {
	var categories []model.Category
	err := br.db.Where("board_id = ?", boardID).Find(&categories).Error
	return categories, err
}

func (br *CategoriesRepo) CreateCategory(category *model.Category) error {
	err := br.db.Create(&category).Error
	return err
}

func (br *CategoriesRepo) UpdateCategory(category *model.Category) error {
	err := br.db.Save(&category).Error
	return err
}

func (br *CategoriesRepo) DeleteCategory(category *model.Category) error {
	err := br.db.Delete(&category).Error
	return err
}
