package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"template/internal/config"
	storage "template/internal/storages"
	postgresconnector "template/pkg/client/postgresql"
	"template/pkg/logs"
	"template/pkg/printer"
)

type Connector struct {
	Client *pgxpool.Pool
	logger *logging.Logger
	cfg    config.Config
}

func NewRepositoryPostgres(logger *logging.Logger, cfg config.Config) storage.Repository {
	return &Connector{
		logger: logger,
		cfg:    cfg,
	}
}

func (c *Connector) Start(_ context.Context) error {
	postgresClient, err := postgresconnector.NewClient(
		5,
		c.cfg.Storage.PostgresUsername,
		c.cfg.Storage.PostgresPassword,
		c.cfg.Storage.PostgresHost,
		c.cfg.Storage.PostgresPort,
		c.cfg.Storage.PostgresDatabase,
	)
	if err != nil {
		return err
	}
	c.Client = postgresClient

	printer.PrintGreen(fmt.Sprintf("Connect postgres %s:%s", c.cfg.Storage.PostgresHost, c.cfg.Storage.PostgresPort))
	return err
}

func (c *Connector) Stop(_ context.Context) error {
	c.Client.Close()
	return nil
}
