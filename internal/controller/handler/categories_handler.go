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

func (bh *CategoriesHandler) GetCategories(w http.ResponseWriter, _ *http.Request) {
	HeaderSet(w)

	Categories, err := bh.CategoriesUsecase.GetCategories()
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
	HeaderSet(w)

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := bh.CategoriesUsecase.CreateCategory(&category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (bh *CategoriesHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w)
	
	categoryID := 	mux.Vars(r)["CategoryID"]

	// stringをuint64に変換
	convertedCategoryID, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// リクエストボディをデコード
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		log.Printf("fail: json.NewDecoder, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// リクエストボディのCategoryIDとパスパラメータのCategoryIDが一致しない場合は400を返す
	if category.ID != convertedCategoryID{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := bh.CategoriesUsecase.UpdateCategory(&category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (bh *CategoriesHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w)

	categoryID := 	mux.Vars(r)["CategoryID"]

	// stringをuint64に変換
	convertedCategoryID, err := strconv.ParseUint(categoryID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	category := model.Category{
		ID: convertedCategoryID,
	}

	if err := bh.CategoriesUsecase.DeleteCategory(&category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}