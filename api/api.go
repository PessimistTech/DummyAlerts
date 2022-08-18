package api

import (
	"DummyAlerts/api/handlers"
	"github.com/gin-gonic/gin"
)

func NewApi() *gin.Engine {
	api := gin.Default()
	api.Use(handlers.ErrorHandler())

	v1 := api.Group("/api/v1")
	v1.POST("webhook/:interpreter", handlers.HandleWebhook)

	return api
}
