package server

import (
	"fmt"
	"log"
)

func NewServer(config *Config) *Server {
	db, err := SetUpDB(config.DBConfig)
	if err != nil {
		log.Fatal(err)
	}

	setUpRouter(db)

	return &Server{
		Address: fmt.Sprintf(":%s", config.APIConfig.Port),
	}
}
