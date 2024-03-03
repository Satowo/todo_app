package setup

import (
	"net/http"

	"github.com/satowo/todo-app/internal/controller/handler"
	"github.com/satowo/todo-app/internal/controller/router"
	"github.com/satowo/todo-app/internal/repository"
	"github.com/satowo/todo-app/internal/usecase"
)

func NewServer() {
	repo := repository.NewBoardsRepo(db)

	boardsUsecase := usecase.NewBoardsUsecase(repo)

	boardsHandler := handler.NewBoardsHandler(boardsUsecase)

	boardsRouter := router.NewBoardsRouter(boardsHandler)

	http.HandleFunc("/boards", boardsRouter.BoardsRouter)
}