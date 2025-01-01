package handler

import (
	"github.com/emmonbear/wallet-exchanger/internal/handler/auth"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	auth.AuthHandler
}

func NewHandler() *Handler {
	return &Handler{
		AuthHandler: auth.NewHandler(),
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/api/v1")
	{
		auth.POST("/register", h.SignUp)
		auth.POST("/login", h.SignIn)
	}

	// TODO Добавить инициализацию
	api := router.Group("/api/v1")
	{
		api.GET("/balance")
		api.POST("/deposit")
		api.POST("/withdraw")
		api.POST("/exchange")
		api.GET("/exchange/rates")
	}

	return router
}
