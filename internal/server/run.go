package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Address string
	router *mux.Router
}

// APIサーバーの起動
func  (s *Server) Run(){
	log.Printf("Server is running on %s", s.Address)

	log.Printf("web url: %s",config.WebConfig.WebURL)
	err := http.ListenAndServe(s.Address, s.router)
	if err != nil {
		log.Fatalln(err)
	}
}