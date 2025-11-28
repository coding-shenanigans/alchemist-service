package main

import (
	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/handler"
	"github.com/coding-shenanigans/alchemist-service/internal/middleware"
)

func main() {
	router := gin.Default()

	publicRouterGroup := router.Group("/")
	protectedRouterGroup := router.Group("/", middleware.AuthMiddleware())

	handler.RegisterEndpoints(publicRouterGroup, protectedRouterGroup)

	router.Run(":9000")
}
