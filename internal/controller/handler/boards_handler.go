package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satowo/todo-app/internal/controller/types"
	"github.com/satowo/todo-app/internal/usecase"
)

type (
	IBoardsHandler interface {
		GetBoards(w http.ResponseWriter, _ *http.Request)
		CreateBoard(w http.ResponseWriter, r *http.Request)
		UpdateBoard(w http.ResponseWriter, r *http.Request)
		DeleteBoard(w http.ResponseWriter, r *http.Request)
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

func (bh *BoardsHandler) GetBoards(w http.ResponseWriter, _ *http.Request) {
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
	var req types.CreateBoardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bh.boardsUsecase.CreateBoard(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bh *BoardsHandler) UpdateBoard(w http.ResponseWriter, r *http.Request) {
	boardID := 	mux.Vars(r)["boardID"]

	// stringをuint64に変換
	convertedboardID, err := strconv.ParseUint(boardID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディをデコード
	var req types.UpdateBoardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bh.boardsUsecase.UpdateBoard(convertedboardID, req.BoardTitle); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *BoardsHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	boardID := 	mux.Vars(r)["boardID"]

	// stringをuint64に変換
	convertedboardID, err := strconv.ParseUint(boardID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bh.boardsUsecase.DeleteBoard(convertedboardID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}