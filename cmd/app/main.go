package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/joho/godotenv"
	"github.com/Marif226/effective-mobile-assessment/internal/handler"
	"github.com/Marif226/effective-mobile-assessment/internal/repository"
	"github.com/Marif226/effective-mobile-assessment/internal/service"
	"gopkg.in/yaml.v2"
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
		Addr: ":8080",
		Handler: handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
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