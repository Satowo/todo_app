package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satowo/todo-app/internal/usecase"
)

type (
	IArchivedItemsHandler interface {
		GetArchivedItems(w http.ResponseWriter, r *http.Request)
		UnArchiveItem(w http.ResponseWriter, r *http.Request)
	}
	ArchivedItemsHandler struct {
		ArchivedItemsUsecase usecase.IItemsUsecase
	}
)

func NewArchivedItemsHandler(au usecase.IItemsUsecase) *ArchivedItemsHandler {
	return &ArchivedItemsHandler{
		ArchivedItemsUsecase: au,
	}
}

func (ah *ArchivedItemsHandler) GetArchivedItems(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータのboardIDを取得
	boardID := 	r.URL.Query().Get("board_id")
	convertedBoardID, err := strconv.ParseUint(boardID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Archiveditems, err := ah.ArchivedItemsUsecase.GetArchivedItems(convertedBoardID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(&Archiveditems)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (ih *ArchivedItemsHandler) UnArchiveItem(w http.ResponseWriter, r *http.Request) {
	ArchiveditemID := mux.Vars(r)["itemID"]

	// stringをuint64に変換
	convertedArchivedItemID, err := strconv.ParseUint(ArchiveditemID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ih.ArchivedItemsUsecase.UnArchiveItem(convertedArchivedItemID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}