package main

import (
	"log/slog"
	"os"

	"github.com/emmonbear/wallet-exchanger/internal/config"
	"github.com/emmonbear/wallet-exchanger/internal/handler"
	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/handlers/slogpretty"
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

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("configuration",
		slog.String("env", cfg.Env),
		slog.Any("listen", cfg.Listen),
		slog.Any("storage_config", cfg.StorageConfig),
	)

	log.Debug("debug messages are enabled")

	db, err := postgres.New(cfg)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	log.Info("successful database initialization")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services, log)
	srv := new(server.Server)
	if err := srv.Run(cfg.Listen.PortEndpoint, handlers.InitRoutes()); err != nil {
		log.Error("error occured while running http server", sl.Err(err))
		os.Exit(1)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	return slog.New(handler)
}
