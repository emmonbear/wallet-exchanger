package wallet

import (
	"errors"
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
	const (
		op                  = "handler.wallet.Deposit"
		userNotFoundMessage = "user id not found"
		invalidRequest      = "Invalid amount or currency"
	)

	sl.LogRequest(h.logger, ctx, op)
	userID, ok := ctx.Get(middleware.UserCtx)

	if !ok {
		err := errors.New(userNotFoundMessage)
		sl.LogError(h.logger, op, err)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, userNotFoundMessage)
		return
	}

	input := &model.Wallet{
		UserID: userID.(int),
	}

	if err := ctx.BindJSON(input); err != nil {
		sl.LogError(h.logger, op, err)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sl.LogInfo(h.logger, op, "Attempting to replenish the account")

	_, err := h.services.WalletService.Deposit(input)
	if err != nil {
		sl.LogError(h.logger, op, err,
			slog.Int("UserID", input.UserID),
			slog.String("Currency", input.Currency),
			slog.Float64("Amount", input.Amount),
		)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, invalidRequest)
		return
	}
	sl.LogInfo(h.logger, op, "Balance successfully replenished",
		slog.Int("UserID", input.UserID),
		slog.String("Currency", input.Currency),
		slog.Float64("Amount", input.Amount),
	)

	balance, _ := h.services.BalanceService.GetBalance(userID.(int))

	sl.NewSuccessResponse(ctx, http.StatusCreated, gin.H{
		"message": "Account topped up successfully",
		"new_balance": gin.H{
			"USD": balance.USD,
			"RUB": balance.RUB,
			"EUR": balance.EUR,
		},
	})
}

func (h *handler) Withdraw(ctx *gin.Context) {}
