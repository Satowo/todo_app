package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	ICategoriesRepo interface {
		GetCategories(boardID uint64) ([]model.Category, error)
		CreateCategory(category *model.Category) error
		UpdateCategory(categoryID uint64, categoryTitle string) error
		DeleteCategory(category uint64) error
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
	err := br.db.Where("board_id = ? AND deleted = ?", boardID, false).Find(&categories).Error
	return categories, err
}

func (br *CategoriesRepo) CreateCategory(category *model.Category) error {
	err := br.db.Create(&category).Error
	return err
}

func (br *CategoriesRepo) UpdateCategory(categoryID uint64, categoryTitle string) error {
	err := br.db.Model(&model.Category{}).Where("id = ?", categoryID).Update("title", categoryTitle).Error
	return err
}

func (br *CategoriesRepo) DeleteCategory(categoryID uint64) error {
	err := br.db.Model(&model.Category{}).Where("id = ?", categoryID).Update("deleted", true).Error
	return err
}
