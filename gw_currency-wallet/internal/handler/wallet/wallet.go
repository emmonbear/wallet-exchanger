package wallet

import "github.com/gin-gonic/gin"

type WalletHandler interface {
	Deposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
}

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Deposit(ctx *gin.Context) {}

func (h *handler) Withdraw(ctx *gin.Context) {}
