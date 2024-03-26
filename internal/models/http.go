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
	http.HandleFunc("GET /properties", handlers.Home)
	http.HandleFunc("GET /properties/{id}", handlers.FindHandler)
	http.HandleFunc("POST /properties/{id}", handlers.PostHandler)
	http.ListenAndServe(s.port, nil)
	return nil
}
