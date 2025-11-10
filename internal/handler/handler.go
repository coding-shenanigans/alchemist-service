package handler

import "github.com/gin-gonic/gin"

// Registers all the endpoints handled by the service.
func RegisterEndpoints(router *gin.Engine) {
	opsHandler := newOpsHandler()

	router.GET("/health", opsHandler.health)
}
