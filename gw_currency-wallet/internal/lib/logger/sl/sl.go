package sl

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

func NewErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"error": message})
}

func NewSuccessResponse(ctx *gin.Context, statusCode int, message gin.H) {
	ctx.JSON(statusCode, gin.H{
		"message": message,
	})
}

func LogRequest(logger *slog.Logger, ctx *gin.Context, operation string) *slog.Logger {
	return logger.With(
		slog.String("operation", operation),
		slog.String("method", ctx.Request.Method),
		slog.String("path", ctx.Request.URL.Path),
		slog.String("remote_addr", ctx.ClientIP()),
		slog.String("user_agent", ctx.Request.UserAgent()),
	)
}

func LogError(logger *slog.Logger, operation string, err error, details ...interface{}) {
	logger.Error("Error occurred", append([]interface{}{
		slog.String("operation", operation),
		slog.String("error", err.Error()),
	}, details...)...)
}

func LogInfo(logger *slog.Logger, operation string, message string, details ...interface{}) {
	logger.Info(message, append([]interface{}{
		slog.String("operation", operation),
	}, details...)...)
}
