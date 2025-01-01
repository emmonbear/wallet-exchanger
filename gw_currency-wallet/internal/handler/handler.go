package handler

import (
	"github.com/emmonbear/wallet-exchanger/internal/handler/auth"
	"github.com/emmonbear/wallet-exchanger/internal/handler/balance"
	"github.com/emmonbear/wallet-exchanger/internal/handler/exchange"
	"github.com/emmonbear/wallet-exchanger/internal/handler/wallet"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	auth.AuthHandler
	balance.BalanceHandler
	exchange.ExchangeHandler
	wallet.WalletHandler
}

func NewHandler() *Handler {
	return &Handler{
		AuthHandler:     auth.NewHandler(),
		BalanceHandler:  balance.NewHandler(),
		ExchangeHandler: exchange.NewHandler(),
		WalletHandler:   wallet.NewHandler(),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/v1")
	{
		auth.POST("/register", h.SignUp)
		auth.POST("/login", h.SignIn)
	}

	// TODO Добавить инициализацию
	api := router.Group("/api/v1")
	{
		api.GET("/balance", h.GetBalance)
		api.POST("/deposit", h.Deposit)
		api.POST("/withdraw", h.Withdraw)
		api.POST("/exchange", h.Exchange)
		api.GET("/exchange/rates", h.GetRates)
	}

	return router
}
