package balance

import (
	"fmt"
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
	userID, ok := ctx.Get(middleware.UserCtx) // Извлекаем userID из контекста
	if !ok {
		errMsg := "user id not found"
		h.logger.Error(errMsg)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errMsg, h.logger, fmt.Errorf(errMsg))
		return
	}

	balance, err := h.services.BalanceService.GetBalance(userID.(int)) // Передаем userID в сервис
	if err != nil {
		errMsg := "error in receiving balance"
		h.logger.Error(errMsg, slog.String("error", err.Error())) // Логируем ошибку, если она есть
		sl.NewErrorResponse(ctx, http.StatusInternalServerError, errMsg, h.logger, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"balance": balance,
	})
}
