package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/config"
	"github.com/coding-shenanigans/alchemist-service/internal/handler"
	"github.com/coding-shenanigans/alchemist-service/internal/middleware"
)

func main() {
	config.Load()

	router := gin.Default()
	router.Use(cors.New(config.GetCorsConfig()))

	publicRouterGroup := router.Group("/")
	protectedRouterGroup := router.Group("/", middleware.RequiredAuthMiddleware())
	hybridRouterGroup := router.Group("/", middleware.OptionalAuthMiddleware())

	handler.RegisterEndpoints(
		publicRouterGroup,
		protectedRouterGroup,
		hybridRouterGroup,
	)

	router.Run(":9000")
}
