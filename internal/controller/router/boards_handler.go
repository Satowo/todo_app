package router

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satowo/todo-app/internal/controller/handler"
)

type (
	IBoardsRouter interface {
		BoardsRouter(w http.ResponseWriter, r *http.Request)
	}
	BoardsRouter struct {
		boardsHandler handler.IBoardsHandler
	}
)

func NewBoardsRouter(bh handler.IBoardsHandler) *BoardsRouter {
	return &BoardsRouter{
		boardsHandler: bh,
	}
}

func (br *BoardsRouter) BoardsRouter(w http.ResponseWriter, r *http.Request) {
	HeaderSet(w)
	switch r.Method {
	case http.MethodOptions:
		w.Header()
	case http.MethodGet:
		br.boardsHandler.GetBoards(w)

	case http.MethodPost:
		// br.boardsHandler.UserRegisterController(w, r)

	default:
		log.Printf("fail: HTTP Method is %s\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}