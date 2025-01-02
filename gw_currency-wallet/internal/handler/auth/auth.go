package auth

import (
	"log/slog"
	"net/http"

	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/model"
	"github.com/emmonbear/wallet-exchanger/internal/service"
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
	var input model.User
	if err := ctx.BindJSON(&input); err != nil {
		sl.NewErrorResponse(ctx, http.StatusBadRequest, err.Error(), h.logger, err)
		return
	}

	id, err := h.services.AuthService.CreateUser(input)
	if err != nil {
		sl.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error(), h.logger, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) SignIn(ctx *gin.Context) {
	var input model.AuthRequest

	if err := ctx.BindJSON(&input); err != nil {
		sl.NewErrorResponse(ctx, http.StatusBadRequest, err.Error(), h.logger, err)
		return
	}

	token, err := h.services.AuthService.GenerateToken(input.Username, input.Password)
	if err != nil {
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error(), h.logger, err)
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
