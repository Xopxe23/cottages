package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Xopxe23/cottages/internal/domain"
	"github.com/Xopxe23/cottages/internal/repository/sqlite"
	"github.com/Xopxe23/cottages/internal/service"
	"github.com/Xopxe23/cottages/internal/transport/rest"
	"github.com/Xopxe23/cottages/pkg/database"
	hasher "github.com/Xopxe23/cottages/pkg/hash"
)

const sqliteStoragePath = "./data/sqlite/storage.db"

func main() {
	storage, err := database.NewStorage(sqliteStoragePath)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.DB.Close()
	if err = storage.PrepareDatabase(); err != nil {
		log.Fatal(err)
	}
	userRepo := repository.NewUserRepository(storage.DB)
	hasher := hasher.NewSHA1Hasher("salt")
	authService := service.NewAuthService(userRepo, hasher)
	handler := rest.NewHandler(authService)
	srv := new(domain.Server)
	go func() {
		if err := srv.Run("8000", handler.InitRouter()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	log.Printf("Server started at %s", time.Now().Format("2006-01-02 15:04:05"))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("Server shutdown at %s", time.Now().Format("2006-01-02 15:04:05"))
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
}
