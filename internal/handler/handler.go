package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/satowo/todo-app/internal/usecase"
)

func GetBoards(w http.ResponseWriter) {
	boards, err := usecase.GetBoards()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(boards)
	if err != nil {
		log.Printf("fail: json.Marshal, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}