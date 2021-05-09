package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-serv-arduino/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h* Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	news := router.Group("/news")
	{
		news.POST("", h.getNews)
	}

	return router
}
