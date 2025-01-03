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
	UserCtx             = "userID"
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
		errMsg := "empty auth header"
		h.logger.Error("Authorization header is missing", slog.String("error", errMsg))
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errMsg, h.logger, fmt.Errorf(errMsg))
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		errMsg := "invalid auth header"
		h.logger.Error("Authorization header format is incorrect", slog.String("error", errMsg), slog.String("header", header))
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errMsg, h.logger, fmt.Errorf(errMsg))
		return
	}

	h.logger.Info("Parsing token", slog.String("token", headerParts[1]))

	userID, err := h.services.AuthService.ParseToken(headerParts[1])
	if err != nil {
		h.logger.Error("Error parsing token", slog.String("error", err.Error()), slog.String("token", headerParts[1]))
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, err.Error(), h.logger, err)
		return
	}

	h.logger.Info("Token parsed successfully", slog.Int("userID", userID))
	ctx.Set(UserCtx, userID)
}
