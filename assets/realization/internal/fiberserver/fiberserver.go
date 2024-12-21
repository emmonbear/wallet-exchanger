package fiberserver

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"template/internal/delivery"

	//swagger "github.com/arsmn/fiber-swagger/v2"
	//_ "gitlab.axarea.ru/main/aifory/svc/gateway/docs"
	logging "template/pkg/logs"

	"github.com/gofiber/fiber/v2/middleware/cors"
	//authclient "gitlab.axarea.ru/main/aifory/proto/auth-center"
	//"gitlab.axarea.ru/main/aifory/svc/gateway/utils"
	//"gitlab.com/zhazhazha1/dimaprojects/packages/logger"
	//"gitlab.com/zhazhazha1/dimaprojects/packages/logger/log"
	//"gitlab.axarea.ru/main/aifory/svc/gateway/internal/delivery"
	//mw "gitlab.axarea.ru/main/aifory/svc/gateway/internal/middleware"
	//"gitlab.axarea.ru/main/aifory/svc/gateway/internal/usecase"
	//"gitlab.axarea.ru/main/aifory/svc/gateway/pkg/errlist"
)

type FiberServer struct {
	cfg   Config
	fiber *fiber.App
	//useCase usecase.UseCase
	//auth    *authclient.Connector
	handler *delivery.Handler
	//mw      mw.Middleware
	//logger  *logger.LogManager
	log *logging.Logger
}

func New(
	logger *logging.Logger,
	cfg Config,
	// useCase usecase.UseCase,
	// auth *authclient.Connector,
	handler *delivery.Handler,
	// mw mw.Middleware,
) *FiberServer {
	return &FiberServer{
		cfg: cfg,
		log: logger,
		//useCase: useCase,
		fiber: fiber.New(
			fiber.Config{
				DisableStartupMessage: true,
				ProxyHeader:           cfg.IpHeader,
				BodyLimit:             cfg.BodyLimit,
				StreamRequestBody:     *cfg.StreamRequestBody,
			}),
		//auth:    auth,
		handler: handler,
		//mw:      mw,
	}
}

func (f *FiberServer) Start(_ context.Context) error {
	f.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     f.cfg.AllowOrigins,
		AllowCredentials: true,
		AllowHeaders:     f.cfg.AllowHeaders,
		ExposeHeaders:    f.cfg.ExposeHeaders,
	}))
	//f.fiber.Use(f.mw.GetErrorsMiddleware(
	//	log.Logger,
	//	f.cfg.IpHeader,
	//	nil, // todo
	//	func(c *fiber.Ctx) any {
	//		type UserNotFoundErr struct {
	//			Msg string `json:"msg"`
	//		}
	//
	//		user, err := mw.GetUser(c)
	//		if err != nil {
	//			return UserNotFoundErr{
	//				Msg: err.Error(),
	//			}
	//		}
	//
	//		return user
	//	},
	//	f.cfg.SecureReqJsonPaths,
	//	f.cfg.SecureResJsonPaths,
	//	f.cfg.ShowUnknownErrorsInResponse,
	//))
	//f.fiber.Use(f.mw.DefaultMiddleware)

	f.bind()

	go func() {
		if err := f.fiber.Listen(f.cfg.Host); err != nil {
			f.log.Fatal(err.Error())
		}
	}()

	return nil
}

func (f *FiberServer) Stop(_ context.Context) error {
	okCh, errCh := make(chan struct{}), make(chan error)
	go func() {
		if err := f.fiber.Shutdown(); err != nil {
			errCh <- err
		}
		okCh <- struct{}{}
	}()

	select {
	case <-okCh:
		return nil
	case err := <-errCh:
		return err
		//case <-time.After(f.cfg.StopTimeout.Duration):
		//	return errlist.ErrStartTimeout
	}
}

func (f *FiberServer) bind() {
	//adminRoute := f.fiber.Group("/admin")
	//delivery.MapAdminV1Routes(adminRoute, f.mw, f.handler)

	clientRoute := f.fiber.Group("/lk")
	delivery.MapClientV1Routes(clientRoute, f.handler)

	//delivery.MapClientV1Routes(clientRoute, f.mw, f.handler)

	//f.fiber.Get("/health_check", func(c *fiber.Ctx) error {
	//	return utils.ReturnFiberStatusOK(c)
	//})
	//
	//f.fiber.Get("/lk/monitoring/health_check", func(c *fiber.Ctx) error {
	//	return utils.ReturnFiberStatusOK(c)
	//})
	//
	//f.fiber.Get("/swagger/*", swagger.New(swagger.Config{ // default
	//	URL: "/swagger/doc.json",
	//}))
	//
	//f.fiber.Get("lk/swagger/*", swagger.New(swagger.Config{ // default
	//	URL: "doc.json",
	//}))
}
