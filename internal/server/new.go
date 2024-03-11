package server

import (
	"fmt"
	"log"

	"github.com/satowo/todo-app/internal/controller/router"
)

func NewServer(config *Config) *Server {
	db, err := SetUpDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	router := router.SetUpRouter(db)

	return &Server{
		Address: fmt.Sprintf(":%s", config.APIConfig.Port),
		router:  router,
	}
}
