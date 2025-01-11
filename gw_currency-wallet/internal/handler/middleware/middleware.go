package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/lib/logger/sl"
	"github.com/emmonbear/wallet-exchanger/gw_currency-wallet/internal/service"
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
	const (
		op                      = "handler.middleware.UserIdentity"
		errEmptyAuthHeader      = "empty auth header"
		errInvalidAuthHeader    = "invalid auth header"
		errParsingToken         = "error parsing token"
		successParsingToken     = "Token parsed successfully"
		successParsingUserToken = "Parsing token"
	)

	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		err := errors.New(errEmptyAuthHeader)
		sl.LogError(
			h.logger, op, err, slog.String("error", errEmptyAuthHeader),
		)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errEmptyAuthHeader)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		err := errors.New(errInvalidAuthHeader)
		sl.LogError(h.logger, op, err,
			slog.String("error", errInvalidAuthHeader),
			slog.String("header", header),
		)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errInvalidAuthHeader)
		return
	}

	sl.LogInfo(
		h.logger, op, successParsingUserToken,
		slog.String("token", headerParts[1]),
	)

	userID, err := h.services.AuthService.ParseToken(headerParts[1])
	if err != nil {
		sl.LogError(h.logger, op, err,
			slog.String(
				"error", err.Error()), slog.String("token", headerParts[1]),
		)
		sl.NewErrorResponse(ctx, http.StatusUnauthorized, errParsingToken)
		return
	}

	sl.LogInfo(h.logger, op, successParsingToken, slog.Int("userID", userID))
	ctx.Set(UserCtx, userID)
}
