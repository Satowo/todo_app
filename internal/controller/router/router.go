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
	repo := repository.NewBoardsRepo(db)

	// usecase
	boardsUsecase := usecase.NewBoardsUsecase(repo)

	// handler
	boardsHandler := handler.NewBoardsHandler(boardsUsecase)

	// エンドポイントの設定
	r := mux.NewRouter()
	r.HandleFunc("/boards", boardsHandler.GetBoards).Methods("GET")
	r.HandleFunc("/boards", boardsHandler.CreateBoard).Methods("POST")
	r.HandleFunc("/boards/{boardID}", boardsHandler.UpdateBoard).Methods("POST")

	return r
}

