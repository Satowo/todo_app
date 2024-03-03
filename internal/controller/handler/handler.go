package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/satowo/todo-app/internal/usecase"
)

type (
	IBoardsHandler interface {
		GetBoards(w http.ResponseWriter)
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

	bytes, err := json.Marshal(boards)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}