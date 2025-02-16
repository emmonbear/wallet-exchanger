package handler

import (
	"log/slog"

	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler/auth"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler/balance"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler/exchange"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler/middleware"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/handler/wallet"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	*service.Service
	auth.AuthHandler
	balance.BalanceHandler
	exchange.ExchangeHandler
	wallet.WalletHandler
	middleware.Middleware
}

func NewHandler(
	services *service.Service, logger *slog.Logger,
) *Handler {
	return &Handler{
		AuthHandler:     auth.NewHandler(logger, services),
		BalanceHandler:  balance.NewHandler(logger, services),
		ExchangeHandler: exchange.NewHandler(),
		WalletHandler:   wallet.NewHandler(logger, services),
		Middleware:      middleware.NewHandler(logger, services),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/v1")
	{
		auth.POST("/register", h.SignUp)
		auth.POST("/login", h.SignIn)
	}

	api := router.Group("/api/v1", h.UserIdentity)
	{
		api.GET("/balance", h.GetBalance)
		api.POST("/deposit", h.Deposit)
		api.POST("/withdraw", h.Withdraw)
		api.POST("/exchange", h.Exchange)
		api.GET("/exchange/rates", h.GetRates)
	}

	return router
}
