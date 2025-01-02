package wallet

import (
	"net/http"

	"github.com/emmonbear/wallet-exchanger/internal/handler/middleware"
	"github.com/gin-gonic/gin"
)

type WalletHandler interface {
	Deposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
}

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Deposit(ctx *gin.Context) {
	id, _ := ctx.Get(middleware.UserCtx)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *handler) Withdraw(ctx *gin.Context) {}
