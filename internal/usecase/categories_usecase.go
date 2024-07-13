package usecase

import (
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/repository"
)

type (
	ICategoriesUsecase interface {
		GetCategories(boardID uint64) ([]model.Category, error)
		CreateCategory(boardID uint64, categoryTitle string) error
		UpdateCategory(categoryID uint64, categoryTitle string) error
		DeleteCategory(categoryID uint64) error
	}
	CategoriesUsecase struct {
		repo repository.ICategoriesRepo
	}
)

func NewCategoriesUsecase(repo repository.ICategoriesRepo) *CategoriesUsecase {
	return &CategoriesUsecase{
		repo: repo,
	}
}

func (bu *CategoriesUsecase) GetCategories(boardID uint64) ([]model.Category, error) {
	categories, err := bu.repo.GetCategories(boardID)

	return categories, err
}

func (bu *CategoriesUsecase) CreateCategory(boardID uint64, categoryTitle string) error {
	category := &model.Category{
		BoardID: boardID,
		Title: categoryTitle,
		Deleted: false,
	}

	err := bu.repo.CreateCategory(category)

	return err
}

func (bu *CategoriesUsecase) UpdateCategory(categoryID uint64, categoryTitle string) error {
	err := bu.repo.UpdateCategory(categoryID, categoryTitle)

	return err
}

func (bu *CategoriesUsecase) DeleteCategory(categoryID uint64) error {
	err := bu.repo.DeleteCategory(categoryID)

	return err
}