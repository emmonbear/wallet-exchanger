package wallet

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/emmonbear/wallet-exchanger/internal/handler/middleware"
	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/service"
	"github.com/gin-gonic/gin"
)

type WalletHandler interface {
	Deposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
}

type handler struct {
	logger   *slog.Logger
	services *service.Service
}

func NewHandler(logger *slog.Logger, services *service.Service) *handler {
	return &handler{
		logger:   logger,
		services: services,
	}
}

func (h *handler) Deposit(ctx *gin.Context) {

	const op = "user is not found"
	id, ok := ctx.Get(middleware.UserCtx)
	// TODO Исправить логирование ошибок
	if !ok {
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, op, h.logger, fmt.Errorf(op))
		return
	}

	deposit, err := h.services.BalanceService.Deposit()
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) Withdraw(ctx *gin.Context) {}
