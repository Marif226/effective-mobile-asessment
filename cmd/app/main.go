package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Marif226/effective-mobile-assessment/internal/handler"
	"github.com/Marif226/effective-mobile-assessment/internal/repository"
	"github.com/Marif226/effective-mobile-assessment/internal/service"
)

func main() {
	repository := repository.New(nil)
	services := service.New(repository)
	handlers := handler.New(services)

	httpServer := http.Server{
		Addr: ":8080",
		Handler: handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}