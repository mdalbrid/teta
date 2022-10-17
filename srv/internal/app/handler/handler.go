package handler

import (
	"github.com/gin-gonic/gin"
	"go-app/internal/domain/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/", h.create)
	}

	return router
}
