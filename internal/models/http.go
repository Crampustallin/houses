package models

import (
	"net/http"

	"github.com/Crampustallin/houses/internal/handlers"
	"github.com/Crampustallin/houses/internal/middleware"
)

type HTTPServer struct {
	port string
}

func NewServer(port string) *HTTPServer {
	return &HTTPServer{port}
}

func (s HTTPServer) Open() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /properties", handlers.GetListHandler)
	router.HandleFunc("GET /properties/{id}", handlers.FindHandler)

	postPropRouter := http.NewServeMux()
	postPropRouter.HandleFunc("POST /properties", handlers.PostHandler)

	router.Handle("/", middleware.ValidateProp(postPropRouter))

	server := http.Server{
		Addr: s.port,
		Handler: middleware.Auth(router),
	}

	server.ListenAndServe()

	return nil
}
