package usecase

import (
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/repository"
)

type (
	IItemsUsecase interface {
		GetItems(categoryID uint64) ([]model.Item, error)
		CreateItem(categoryID uint64, itemTitle, content, expiredAt string) error
		UpdateItem(itemID, categoryID uint64, item_title, content, expired_at string) error
		DeleteItem(itemID uint64) error
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

func (iu *ItemsUsecase) CreateItem(categoryID uint64, itemTitle, content, expiredAt string) error {
	item := &model.Item{
		CategoryID: categoryID,
		Title: itemTitle,
		Content: content,
		ExpiredAt: expiredAt,
		Deleted: false,
	}

	err := iu.repo.CreateItem(item)

	return err
}

func (iu *ItemsUsecase) UpdateItem(itemID, categoryID uint64, item_title, content, expired_at string) error {
	err := iu.repo.UpdateItem(itemID, categoryID, item_title, content, expired_at)

	return err
}

func (iu *ItemsUsecase) DeleteItem(itemID uint64) error {
	err := iu.repo.DeleteItem(itemID)

	return err
}
