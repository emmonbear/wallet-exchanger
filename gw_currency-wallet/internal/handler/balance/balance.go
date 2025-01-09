package balance

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/emmonbear/wallet-exchanger/internal/handler/middleware"
	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/service"
	"github.com/gin-gonic/gin"
)

type BalanceHandler interface {
	GetBalance(ctx *gin.Context)
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

func (h *handler) GetBalance(ctx *gin.Context) {
	const (
		op                  = "handler.balance.GetBalance"
		userNotFoundMessage = "user id not found"
	)

	sl.LogRequest(h.logger, ctx, op)

	userID, ok := ctx.Get(middleware.UserCtx)
	if !ok {
		err := errors.New(userNotFoundMessage)
		sl.LogError(h.logger, op, err)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, userNotFoundMessage)
		return
	}

	sl.LogInfo(h.logger, op, "Attempting to retrieve user balance")
	balance, _ := h.services.BalanceService.GetBalance(userID.(int))
	sl.LogInfo(h.logger, op, "Balance successfully received")

	sl.NewSuccessResponse(ctx, http.StatusOK, gin.H{
		"balance": gin.H{
			"USD": balance.USD,
			"RUB": balance.RUB,
			"EUR": balance.EUR,
		},
	})
}
