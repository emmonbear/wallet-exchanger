package main

import (
	"log/slog"
	"os"

	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/config"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/repository/postgres"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/server"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service"
	"github.com/emmonbear/wallet-exchanger/pkg/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.Env)

	log.Info("configuration", slog.Any("Config", cfg))

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
