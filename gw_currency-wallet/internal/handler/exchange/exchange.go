package exchange

import "github.com/gin-gonic/gin"

type ExchangeHandler interface {
	GetRates(ctx *gin.Context)
	Exchange(ctx *gin.Context)
}

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) GetRates(ctx *gin.Context) {}

func (h *handler) Exchange(ctx *gin.Context) {}
