package repository

import (
	"github.com/satowo/todo-app/internal/model"
	"gorm.io/gorm"
)

type (
	IBoardsRepo interface {
		GetBoards() ([]model.Board, error)
		CreateBoard(board *model.Board) error
		UpdateBoard(board *model.Board) error
		DeleteBoard(board uint64) error
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
	err := br.db.Where("deleted = ?", false).Find(&boards).Error
	return boards, err
}

func (br *BoardsRepo) CreateBoard(board *model.Board) error {
	err := br.db.Create(&board).Error
	return err
}

func (br *BoardsRepo) UpdateBoard(board *model.Board) error {
	err := br.db.Save(&board).Error
	return err
}

func (br *BoardsRepo) DeleteBoard(boardID uint64) error {
	err := br.db.Model(&model.Board{}).Where("id = ?", boardID).Update("deleted", true).Error
	return err
}
