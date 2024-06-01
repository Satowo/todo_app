package server

import (
	"fmt"
	"log"
)

func NewServer() *Server {
	db, err := SetUpDB()
	if err != nil {
		log.Fatal(err)
	}

	router := SetUpRouter(db)

	return &Server{
		Address: fmt.Sprintf(":%s", config.APIConfig.APIPort),
		router:  router,
	}
}
