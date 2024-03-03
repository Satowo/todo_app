package server

import (
	"net/http"

	"github.com/satowo/todo-app/internal/controller/handler"
	"github.com/satowo/todo-app/internal/controller/router"
	"github.com/satowo/todo-app/internal/repository"
	"github.com/satowo/todo-app/internal/usecase"
	"gorm.io/gorm"
)

func setUpRouter(db *gorm.DB) {
	// repository
	repo := repository.NewBoardsRepo(db)

	// usecase
	boardsUsecase := usecase.NewBoardsUsecase(repo)

	// handler
	boardsHandler := handler.NewBoardsHandler(boardsUsecase)

	// router
	boardsRouter := router.NewBoardsRouter(boardsHandler)

	// エンドポイントの設定
	http.HandleFunc("/boards", boardsRouter.BoardsRouter)
}

