package main

import (
	"context"
	"github.com/joho/godotenv"
	SarkorTest "github.com/qsnake66/ProductWerehouse"
	"github.com/qsnake66/ProductWerehouse/pkg/handler"
	"github.com/qsnake66/ProductWerehouse/pkg/repository"
	"github.com/qsnake66/ProductWerehouse/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %s", err.Error())
	}

	DbConfig := repository.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
	}

	db := repository.InitDB(DbConfig)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(SarkorTest.Server)

	go func() {
		if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error run server: %s", err.Error())
		}
	}()

	log.Println("starting server")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("stopping server")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("error shutting down server: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to retrieve database connection: %v", err)
	}

	if err = sqlDB.Close(); err != nil {
		log.Fatalf("failed to close database connection: %v", err)
	}
}
