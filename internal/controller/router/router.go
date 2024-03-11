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
	boards_repo := repository.NewBoardsRepo(db)
	categories_repo := repository.NewCategoriesRepo(db)

	// usecase
	boardsUsecase := usecase.NewBoardsUsecase(boards_repo)
	categoriesUsecase := usecase.NewCategoriesUsecase(categories_repo)

	// handler
	boardsHandler := handler.NewBoardsHandler(boardsUsecase)
	categoriesHandler := handler.NewCategoriesHandler(categoriesUsecase)

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

	return r
}

