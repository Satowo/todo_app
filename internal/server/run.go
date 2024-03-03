package server

import (
	"log"
	"net/http"
)

type Server struct {
	Address string
}

// APIサーバーの起動
func  (server *Server) Run(){
	log.Printf("Server is running on %s", server.Address)
	err := http.ListenAndServe(server.Address, nil)
	if err != nil {
		log.Fatalln(err)
	}
}