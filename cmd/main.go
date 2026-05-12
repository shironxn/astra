package main

import (
	"log"

	"github.com/shironxn/astra/internal/config"
	"github.com/shironxn/astra/internal/config/db"
	"github.com/shironxn/astra/internal/handler"
	"github.com/shironxn/astra/internal/repository"
	"github.com/shironxn/astra/internal/service"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	db, err := db.NewDatabase(cfg.Database).Connection()
	if err != nil {
		panic(err)
	}

	authorRepository := repository.NewAuthorRepository(db)
	authorService := service.NewAuthorService(authorRepository)
	authorHandler := handler.NewAuthorHandler(authorService)

	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	server := config.NewServer(cfg.Server, config.Handler{Author: authorHandler, Book: bookHandler})
	log.Printf("Server is running on %s...\n", cfg.Server.Port)
	if err := server.Run(); err != nil {
		panic(err)
	}
}
