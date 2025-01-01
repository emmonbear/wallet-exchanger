package auth

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) SignUp(ctx *gin.Context) {

}

func (h *handler) SignIn(ctx *gin.Context) {

}
