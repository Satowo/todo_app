package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/usecase"
)

type (
	IBoardsHandler interface {
		GetBoards(w http.ResponseWriter)
		CreateBoard(w http.ResponseWriter, r *http.Request)
	}
	BoardsHandler struct {
		boardsUsecase usecase.IBoardsUsecase
	}
)

func NewBoardsHandler(bu usecase.IBoardsUsecase) *BoardsHandler {
	return &BoardsHandler{
		boardsUsecase: bu,
	}
}

func (bh *BoardsHandler) GetBoards(w http.ResponseWriter) {
	boards, err := bh.boardsUsecase.GetBoards()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(&boards)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (bh *BoardsHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var board model.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bh.boardsUsecase.CreateBoard(&board); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}