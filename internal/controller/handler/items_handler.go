package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/satowo/todo-app/internal/model"
	"github.com/satowo/todo-app/internal/usecase"
)

type (
	IItemsHandler interface {
		GetItems(w http.ResponseWriter, r *http.Request)
		CreateItem(w http.ResponseWriter, r *http.Request)
		UpdateItem(w http.ResponseWriter, r *http.Request)
		ArchiveItem(w http.ResponseWriter, r *http.Request)
	}
	ItemsHandler struct {
		ItemsUsecase usecase.IItemsUsecase
	}
)

func NewItemsHandler(iu usecase.IItemsUsecase) *ItemsHandler {
	return &ItemsHandler{
		ItemsUsecase: iu,
	}
}

func (ih *ItemsHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータのboardIDを取得
	categoryID := r.URL.Query().Get("category_id")
	convertedCategoryID, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	items, err := ih.ItemsUsecase.GetItems(convertedCategoryID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(&items)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (ih *ItemsHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := ih.ItemsUsecase.CreateItem(&item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (ih *ItemsHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {	
	itemID := mux.Vars(r)["itemID"]

	// stringをuint64に変換
	convertedItemID, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディをデコード
	var item model.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// リクエストボディのItemIDとパスパラメータのItemIDが一致しない場合は400を返す
	if item.ID != convertedItemID{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ih.ItemsUsecase.UpdateItem(&item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (ih *ItemsHandler) ArchiveItem(w http.ResponseWriter, r *http.Request) {
	itemID := mux.Vars(r)["itemID"]

	// stringをuint64に変換
	convertedItemID, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := ih.ItemsUsecase.ArchiveItem(convertedItemID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}