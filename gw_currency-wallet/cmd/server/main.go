package main

import (
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/emmonbear/wallet-exchanger/internal/config"
	"github.com/emmonbear/wallet-exchanger/internal/handler"
	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/repository"
	"github.com/emmonbear/wallet-exchanger/internal/repository/postgres"
	"github.com/emmonbear/wallet-exchanger/internal/server"
	"github.com/emmonbear/wallet-exchanger/internal/service"
	_ "github.com/lib/pq"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "", "Path to config file")
	flag.Parse()

	if configPath == "" {
		log.Fatalf("Config file path not provided. Use -c flag to specify the path")
	}
}

func main() {
	cfg := config.MustLoad(configPath)
	fmt.Println(cfg)

	log := setupLogger(cfg.Env)
	log.Info("starting wallet-server", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := postgres.New(cfg)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, log)
	srv := new(server.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Error("error occured while running http server", sl.Err(err))
	}

	// router := chi.NewRouter()

	// TODO: middleware - Добавить логгер запросов
	// TODO: middleware - Реализовать recoverer
	// TODO: slog - Реализовать PrettyLogger

	// TODO: run server:
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
