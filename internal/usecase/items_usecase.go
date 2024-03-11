package usecase

import (
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/repository"
)

type (
	IItemsUsecase interface {
		GetItems(categoryID uint64) ([]model.Item, error)
		GetArchivedItems(boardID uint64) ([]model.Item, error)
		CreateItem(item *model.Item) error
		UpdateItem(item *model.Item) error
		ArchiveItem(itemID uint64) error
		UnArchiveItem(itemID uint64) error
	}
	ItemsUsecase struct {
		repo repository.IItemsRepo
	}
)

func NewItemsUsecase(repo repository.IItemsRepo) *ItemsUsecase {
	return &ItemsUsecase{
		repo: repo,
	}
}

func (iu *ItemsUsecase) GetItems(categoryID uint64) ([]model.Item, error) {
	items, err := iu.repo.GetItems(categoryID)

	return 	items, err
}

func (iu *ItemsUsecase) GetArchivedItems(boardID uint64) ([]model.Item, error) {
	items, err := iu.repo.GetArchivedItems(boardID)
	
	return 	items, err
}

func (iu *ItemsUsecase) CreateItem(item *model.Item) error {
	err := iu.repo.CreateItem(item)

	return err
}

func (iu *ItemsUsecase) UpdateItem(item *model.Item) error {
	err := iu.repo.UpdateItem(item)

	return err
}

func (iu *ItemsUsecase) ArchiveItem(itemID uint64) error {
	err := iu.repo.ArchiveItem(itemID)

	return err
}

func (iu *ItemsUsecase) UnArchiveItem(itemID uint64) error {
	err := iu.repo.UnArchiveItem(itemID)

	return err
}