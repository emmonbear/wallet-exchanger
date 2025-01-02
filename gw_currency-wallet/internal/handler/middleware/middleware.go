package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/emmonbear/wallet-exchanger/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "userId"
)

type Middleware interface {
	UserIdentity(ctx *gin.Context)
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

func (h *handler) UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header", h.logger, fmt.Errorf("empty auth header"))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header", h.logger, fmt.Errorf("empty auth header"))
		return
	}

	userId, err := h.services.AuthService.ParseToken(headerParts[1])
	if err != nil {
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error(), h.logger, err)
		return
	}

	ctx.Set(UserCtx, userId)
}
