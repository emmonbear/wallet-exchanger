package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"template/internal/config"
	"template/internal/delivery"
	"template/internal/fiberserver"
	storagePostgres "template/internal/storages/db/postgres"
	logging "template/pkg/logs"
	"template/pkg/utils/lifecycle"
	"time"
)

type (
	App struct {
		cfg  *config.Config
		log  *logging.Logger
		cmps []cmp
	}
	cmp struct {
		Service lifecycle.Lifecycle
		Name    string
	}
)

func New(cfg *config.Config, logger *logging.Logger) *App {
	return &App{
		cfg: cfg,
		log: logger,
	}
}

func (a *App) Start(ctx context.Context) error {
	handler, err := delivery.New()
	if err != nil {
		return err
	}

	pg := storagePostgres.NewRepositoryPostgres(a.log, *a.cfg)

	fiberServer := fiberserver.New(a.log, a.cfg.FiberServer, handler)

	a.cmps = append(a.cmps,
		cmp{fiberServer, fmt.Sprintf("HTTP fiberServer (address: %s)", a.cfg.FiberServer.Host)},
		cmp{pg, fmt.Sprintf("Postgres %s:%s", a.cfg.Storage.PostgresHost, a.cfg.Storage.PostgresPort)},
	)

	okCh, errCh := make(chan struct{}), make(chan error)

	go func() {
		for _, c := range a.cmps {
			// launch server
			if err := c.Service.Start(ctx); err != nil {
				log.Error(err)
				errCh <- err

				return
			}

			log.Infof("%v started", c.Name)
		}
		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return err
	case err := <-errCh:
		return err
	case <-okCh:
		log.Info("Application started!")

		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	a.log.Info("shutting down service...")

	okCh, errCh := make(chan struct{}), make(chan error)

	go func() {
		for i := len(a.cmps) - 1; i > 0; i-- {
			c := a.cmps[i]
			a.log.Infof("stopping %q...", c.Name)

			if err := c.Service.Stop(ctx); err != nil {
				a.log.Error(err)
				errCh <- err

				return
			}
		}

		okCh <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return nil
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info("Application stopped!")
		return nil
	}
}

func (a *App) GetStartTimeout() time.Duration { return 64 }
func (a *App) GetStopTimeout() time.Duration  { return 64 }
