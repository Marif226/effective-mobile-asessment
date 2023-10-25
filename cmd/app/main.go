package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Marif226/effective-mobile-assessment/internal/handler"
	"github.com/Marif226/effective-mobile-assessment/internal/repository"
	"github.com/Marif226/effective-mobile-assessment/internal/service"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

const (
	port = ":8080"
)

func main() {
	// Initialize config file
	dbConf, err := initConfig()
	if err != nil {
		log.Fatalf("error during initilizaing the config file, %s", err.Error())
	}

	// Load .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("error during loading .env file, %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.PGConfig{
		Host: dbConf.Host,
		Port: dbConf.Port,
		Username: dbConf.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName: dbConf.DBName,
		SSLMode: dbConf.SSLMode,	
	})

	if err != nil {
		log.Fatalf("error during connecting to the database: %s", err.Error())
	}

	repository := repository.New(db)
	services := service.New(repository)
	handlers := handler.New(services)

	httpServer := http.Server{
		Addr: port,
		Handler: handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

// initialize config file, return error if failed
func initConfig() (*repository.PGConfig, error) {
	var dbConf *repository.PGConfig
	// open and read config.yml file
	yamlFile, err := os.Open("configs/config.yml")
    if err != nil {
		return dbConf, err
    }
	yamlData, err := io.ReadAll(yamlFile)
	if err != nil {
		return dbConf, err
    }
	
    err = yaml.Unmarshal(yamlData, &dbConf)
    if err != nil {
		return dbConf, err
    }

    return dbConf, nil
}