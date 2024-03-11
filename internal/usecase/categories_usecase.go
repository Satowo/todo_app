package usecase

import (
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/repository"
)

type (
	ICategoriesUsecase interface {
		GetCategories(boardID uint64) ([]model.Category, error)
		CreateCategory(category *model.Category) error
		UpdateCategory(category *model.Category) error
		DeleteCategory(category *model.Category) error
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

	return 	categories, err
}

func (bu *CategoriesUsecase) CreateCategory(category *model.Category) error {
	err := bu.repo.CreateCategory(category)

	return err
}

func (bu *CategoriesUsecase) UpdateCategory(category *model.Category) error {
	err := bu.repo.UpdateCategory(category)

	return err
}

func (bu *CategoriesUsecase) DeleteCategory(category *model.Category) error {
	err := bu.repo.DeleteCategory(category)

	return err
}