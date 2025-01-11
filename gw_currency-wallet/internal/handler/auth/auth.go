package auth

import (
	"log/slog"
	"net/http"

	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/model"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
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

func (h *handler) SignUp(ctx *gin.Context) {
	const (
		op                           = "handler.auth.SignUp"
		userCreateMessage            = "User registered successfully"
		usernameOrEmailExistsMessage = "Username or email already exists"
	)

	sl.LogRequest(h.logger, ctx, op)

	var input model.User

	if err := ctx.BindJSON(&input); err != nil {
		sl.LogError(h.logger, op, err)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sl.LogInfo(
		h.logger, op, "Attempting to create user",
		slog.String("username", input.Username),
		slog.String("email", input.Email),
	)

	err := h.services.AuthService.CreateUser(input)
	if err != nil {
		sl.LogError(
			h.logger, op, err,
			slog.String("username", input.Username),
			slog.String("email", input.Email),
		)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, usernameOrEmailExistsMessage)
		return
	}

	sl.LogInfo(
		h.logger, op, "User created successfully",
		slog.String("username", input.Username),
		slog.String("email", input.Email),
	)

	sl.NewSuccessResponse(ctx, http.StatusCreated, gin.H{
		"message": userCreateMessage,
	})
}

func (h *handler) SignIn(ctx *gin.Context) {
	const (
		op                      = "handler.auth.SignIn"
		userUnauthorizedMessage = "Invalid username or password"
	)

	sl.LogRequest(h.logger, ctx, op)

	var input model.AuthRequest

	if err := ctx.BindJSON(&input); err != nil {
		sl.LogError(h.logger, op, err)
		sl.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sl.LogInfo(
		h.logger, op, "Attempting to generate token",
		slog.String("username", input.Username),
	)

	token, err := h.services.AuthService.GenerateToken(input.Username, input.Password)
	if err != nil {
		sl.LogError(
			h.logger, op, err,
			slog.String("username", input.Username),
		)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, userUnauthorizedMessage)
		return
	}

	sl.LogInfo(h.logger, op, "Token successfully generated")
	sl.NewSuccessResponse(ctx, http.StatusOK, gin.H{
		"token": token,
	})
}
