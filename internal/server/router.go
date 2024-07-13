package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satowo/todo-app/internal/controller/handler"
	"github.com/satowo/todo-app/internal/repository"
	"github.com/satowo/todo-app/internal/usecase"
	"gorm.io/gorm"
)


func SetUpRouter(db *gorm.DB) *mux.Router {
	// repository
	boardsRepo := repository.NewBoardsRepo(db)
	categoriesRepo := repository.NewCategoriesRepo(db)
	itemsRepo := repository.NewItemsRepo(db)

	// usecase
	boardsUsecase := usecase.NewBoardsUsecase(boardsRepo)
	categoriesUsecase := usecase.NewCategoriesUsecase(categoriesRepo)
	itemsUsecase := usecase.NewItemsUsecase(itemsRepo)

	// handler
	boardsHandler := handler.NewBoardsHandler(boardsUsecase)
	categoriesHandler := handler.NewCategoriesHandler(categoriesUsecase)
	itemsHandler := handler.NewItemsHandler(itemsUsecase)

	// エンドポイントの設定
	r := mux.NewRouter()

	r.HandleFunc("/boards", boardsHandler.GetBoards).Methods(http.MethodGet)
	r.HandleFunc("/boards", boardsHandler.CreateBoard).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/boards/{boardID}", boardsHandler.UpdateBoard).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/boards/{boardID}", boardsHandler.DeleteBoard).Methods(http.MethodDelete)

	r.HandleFunc("/categories", categoriesHandler.GetCategories).Methods(http.MethodGet)
	r.HandleFunc("/categories", categoriesHandler.CreateCategory).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/categories/{categoryID}", categoriesHandler.UpdateCategory).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/categories/{categoryID}", categoriesHandler.DeleteCategory).Methods(http.MethodDelete)

	r.HandleFunc("/items", itemsHandler.GetItems).Methods(http.MethodGet)
	r.HandleFunc("/items", itemsHandler.CreateItem).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/items/{itemID}", itemsHandler.UpdateItem).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/items/{itemID}", itemsHandler.DeleteItem).Methods(http.MethodDelete)

	r.Use(CORSMiddleware)
	
	return r
}

