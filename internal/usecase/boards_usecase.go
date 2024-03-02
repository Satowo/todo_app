package usecase

import (
	"context"

	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	IBoardsUsecase interface {
		GetBoards(context.Context) ([]model.Board, error)
	}
	IBoardsRepo interface {
		GetBoards() ([]model.Board, error)
	}
	BoardsUsecase struct {
		repo IBoardsRepo
	}
)

func NewBoardsUsecase(db *gorm.DB) *BoardsUsecase {
	return &BoardsUsecase{
		repo: repo.NewBoardsRepo(db),
	}
}

func (bu *BoardsUsecase)GetBoards(c context.Context) ([]model.Board, error) {
	boards, err := bu.repo.GetBoards()

	return 	boards, err
}