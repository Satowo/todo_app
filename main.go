package main

import (
	"github.com/satowo/todo-app/internal/server"
)

func main() {
	config := server.NewConfig()

	httpServer := server.NewServer(config)

	httpServer.Run()
}