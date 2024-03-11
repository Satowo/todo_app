package router

import (
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
	archivedItemsHandler := handler.NewArchivedItemsHandler(itemsUsecase)

	// エンドポイントの設定
	r := mux.NewRouter()
	r.HandleFunc("/boards", boardsHandler.GetBoards).Methods("GET")
	r.HandleFunc("/boards", boardsHandler.CreateBoard).Methods("POST")
	r.HandleFunc("/boards/{boardID}", boardsHandler.UpdateBoard).Methods("POST")
	r.HandleFunc("/boards/{boardID}", boardsHandler.DeleteBoard).Methods("DELETE")

	r.HandleFunc("/categories", categoriesHandler.GetCategories).Methods("GET")
	r.HandleFunc("/categories", categoriesHandler.CreateCategory).Methods("POST")
	r.HandleFunc("/categories/{categoryID}", categoriesHandler.UpdateCategory).Methods("POST")
	r.HandleFunc("/categories/{categoryID}", categoriesHandler.DeleteCategory).Methods("DELETE")

	r.HandleFunc("/items", itemsHandler.GetItems).Methods("GET")
	r.HandleFunc("/items", itemsHandler.CreateItem).Methods("POST")
	r.HandleFunc("/items/{itemID}", itemsHandler.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{itemID}", itemsHandler.ArchiveItem).Methods("POST")

	r.HandleFunc("/items/archived", archivedItemsHandler.GetArchivedItems).Methods("GET")
	r.HandleFunc("/items/archived/{itemID}", archivedItemsHandler.UnArchiveItem).Methods("POST")

	return r
}

