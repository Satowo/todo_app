package main

import (
	"github.com/satowo/todo-app/internal/server"
)

func main() {
	server.NewConfig()

	httpServer := server.NewServer()

	httpServer.Run()
}