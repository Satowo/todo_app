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
	ICategoriesHandler interface {
		GetCategories(w http.ResponseWriter, _ *http.Request)
		CreateCategory(w http.ResponseWriter, r *http.Request)
		UpdateCategory(w http.ResponseWriter, r *http.Request)
		DeleteCategory(w http.ResponseWriter, r *http.Request)
	}
	CategoriesHandler struct {
		CategoriesUsecase usecase.ICategoriesUsecase
	}
)

func NewCategoriesHandler(bu usecase.ICategoriesUsecase) *CategoriesHandler {
	return &CategoriesHandler{
		CategoriesUsecase: bu,
	}
}

func (bh *CategoriesHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	// クエリパラメータのboardIDを取得
	boardID := 	r.URL.Query().Get("board_id")
	convertedboardID, err := strconv.ParseUint(boardID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	Categories, err := bh.CategoriesUsecase.GetCategories(convertedboardID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(&Categories)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (bh *CategoriesHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req types.CreateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bh.CategoriesUsecase.CreateCategory(req.BoardID, req.CategoryTitle); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bh *CategoriesHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {	
	categoryID := mux.Vars(r)["categoryID"]

	// stringをuint64に変換
	convertedCategoryID, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディをデコード
	var req types.UpdateCategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bh.CategoriesUsecase.UpdateCategory(convertedCategoryID, req.CategoryTitle); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *CategoriesHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryID := mux.Vars(r)["categoryID"]

	// stringをuint64に変換
	convertedCategoryID, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bh.CategoriesUsecase.DeleteCategory(convertedCategoryID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}