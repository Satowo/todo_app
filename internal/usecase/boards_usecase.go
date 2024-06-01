package usecase

import (
	"github.com/satowo/todo-app/internal/controller/types"
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/repository"
)

type (
	IBoardsUsecase interface {
		GetBoards() ([]model.Board, error)
		CreateBoard(param *types.CreateBoardRequest) error
		UpdateBoard(board *model.Board) error
		DeleteBoard(board *model.Board) error
	}
	BoardsUsecase struct {
		repo repository.IBoardsRepo
	}
)

func NewBoardsUsecase(repo repository.IBoardsRepo) *BoardsUsecase {
	return &BoardsUsecase{
		repo: repo,
	}
}

func (bu *BoardsUsecase) GetBoards() ([]model.Board, error) {
	boards, err := bu.repo.GetBoards()

	return 	boards, err
}

func (bu *BoardsUsecase) CreateBoard(param *types.CreateBoardRequest) error {
	board := &model.Board{
		Title: param.BoardTitle,
		Deleted: false,
	}

	err := bu.repo.CreateBoard(board)

	return err
}

func (bu *BoardsUsecase) UpdateBoard(board *model.Board) error {
	err := bu.repo.UpdateBoard(board)

	return err
}

func (bu *BoardsUsecase) DeleteBoard(board *model.Board) error {
	err := bu.repo.DeleteBoard(board)

	return err
}