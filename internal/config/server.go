package config

import (
	"net/http"

	"github.com/shironxn/astra/internal/domain"
)

type Server struct {
	Host    string
	Port    string
	Handler Handler
}

type Handler struct {
	Author domain.AuthorHandler
	Book   domain.BookHandler
}

func NewServer(cfg ServerConfig, handler Handler) *Server {
	return &Server{
		Host:    cfg.Host,
		Port:    cfg.Port,
		Handler: handler,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	// author routes
	router.HandleFunc("POST /authors", s.Handler.Author.Create)
	router.HandleFunc("GET /authors", s.Handler.Author.GetAll)
	router.HandleFunc("GET /authors/{id}", s.Handler.Author.GetById)
	router.HandleFunc("PUT /authors/{id}", s.Handler.Author.Update)
	router.HandleFunc("DELETE /authors/{id}", s.Handler.Author.Delete)

	// book routes
	router.HandleFunc("POST /books", s.Handler.Book.Create)
	router.HandleFunc("GET /books", s.Handler.Book.GetAll)
	router.HandleFunc("GET /books/{id}", s.Handler.Book.GetById)
	router.HandleFunc("PUT /books/{id}", s.Handler.Book.Update)
	router.HandleFunc("DELETE /books/{id}", s.Handler.Book.Delete)

	server := http.Server{
		Addr:    s.Host + ":" + s.Port,
		Handler: router,
	}

	return server.ListenAndServe()
}
