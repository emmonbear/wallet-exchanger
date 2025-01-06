package wallet

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/emmonbear/wallet-exchanger/internal/handler/middleware"
	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/model"
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
	userID, ok := ctx.Get(middleware.UserCtx)
	if !ok {
		errMsg := "user id not found"
		h.logger.Error(errMsg)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errMsg, h.logger, fmt.Errorf(errMsg))
		return
	}

	input := &model.Wallet{
		UserID: userID.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		errMsg := "deposit fail"
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errMsg, h.logger, fmt.Errorf(errMsg))
		return
	}

	h.services.WalletService.Deposit(input)

}

func (h *handler) Withdraw(ctx *gin.Context) {}
