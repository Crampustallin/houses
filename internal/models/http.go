package models

import (
	"net/http"
	"github.com/Crampustallin/houses/internal/handlers"
)

type HTTPServer struct {
	port string
}

func NewServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

func (s HTTPServer) Open() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /properties", handlers.Home)
	router.HandleFunc("GET /properties/{id}", handlers.FindHandler)
	router.HandleFunc("POST /properties/{id}", handlers.PostHandler)

	server := http.Server{
		Addr: s.port,
		Handler: router,
	}
	server.ListenAndServe()
	return nil
}