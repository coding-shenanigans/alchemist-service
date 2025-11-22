package handler

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/database"
	"github.com/coding-shenanigans/alchemist-service/internal/repository"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

// Registers all the endpoints handled by the service.
func RegisterEndpoints(router *gin.Engine) {
	// connect to database
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// create repositories
	userRepository := repository.NewUserRepository(db)
	sessionRepository := repository.NewSessionRepository(db)

	// create services
	authService := service.NewAuthService(userRepository, sessionRepository)

	// create handlers
	opsHandler := newOpsHandler()
	authHandler := newAuthHandler(authService)

	// register endpoints
	router.GET("/health", opsHandler.health)

	router.POST("/auth/signup", authHandler.signup)
	router.POST("/auth/signin", authHandler.signin)
}
