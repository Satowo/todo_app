package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	IBoardsRepo interface {
		GetBoards() ([]model.Board, error)
		CreateBoard(board *model.Board) error
	}
	BoardsRepo struct {
		db *gorm.DB
	}
)

func NewBoardsRepo(db *gorm.DB) *BoardsRepo {
	return &BoardsRepo{db}
}

func (br *BoardsRepo) GetBoards() ([]model.Board, error) {
	var boards []model.Board
	err := br.db.Find(&boards).Error
	return boards, err
}

func (br *BoardsRepo) CreateBoard(board *model.Board) error {
	err := br.db.Create(&board).Error
	return err
}
