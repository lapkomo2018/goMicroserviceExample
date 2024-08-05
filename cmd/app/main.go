package main

import (
	"github.com/joho/godotenv"
	"github.com/lapkomo2018/goServiceExample/config"
	"github.com/lapkomo2018/goServiceExample/internal/service"
	"github.com/lapkomo2018/goServiceExample/internal/storage/mysql"
	grpcServer "github.com/lapkomo2018/goServiceExample/internal/transport/grpc"
	restServer "github.com/lapkomo2018/goServiceExample/internal/transport/rest"
	"github.com/lapkomo2018/goServiceExample/pkg/hash"
	"github.com/lapkomo2018/goServiceExample/pkg/validation"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var cfg *config.Config

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file, proceeding without it: %v", err)
	}

	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// @title           Example Service
// @version         1.0
// @description     Example description

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	storages, err := mysql.New(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	hasher := hash.NewSHA1Hasher(cfg.Hash)
	services := service.New(storages.Struct, hasher)

	validator, err := validation.NewValidator(cfg.Validator)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		server := restServer.New(cfg.RestServer).Init(services.Struct, validator)
		if err := server.Run(); err != nil {
			log.Fatalf("Rest server err: %v", err)
		}
	}()

	go func() {
		server := grpcServer.New(cfg.GrpcServer).Init()
		if err := server.Run(); err != nil {
			log.Fatalf("Grpc server err: %v", err)
		}
	}()

	<-quit
	log.Println("Shutting down servers...")
}
