package sl

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string `json:"message"`
}

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func NewErrorResponse(
	ctx *gin.Context, statusCode int, message string, log *slog.Logger, err error,
) {
	log.Error(message, Err(err))
	ctx.AbortWithStatusJSON(statusCode, response{Message: message})
}
