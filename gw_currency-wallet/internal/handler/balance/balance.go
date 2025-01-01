package balance

import "github.com/gin-gonic/gin"

type BalanceHandler interface {
	GetBalance(ctx *gin.Context)
}

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) GetBalance(ctx *gin.Context) {

}
