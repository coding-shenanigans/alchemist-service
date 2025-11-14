package handler

import "github.com/gin-gonic/gin"

// Registers all the endpoints handled by the service.
func RegisterEndpoints(router *gin.Engine) {
	// create handlers
	opsHandler := newOpsHandler()
	authHandler := newAuthHandler()

	// register endpoints
	router.GET("/health", opsHandler.health)
	
	router.POST("/auth/signup", authHandler.signup)
}
