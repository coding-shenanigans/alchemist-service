package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/handler"
	"github.com/coding-shenanigans/alchemist-service/internal/middleware"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	publicRouterGroup := router.Group("/")
	protectedRouterGroup := router.Group("/", middleware.AuthMiddleware())

	handler.RegisterEndpoints(publicRouterGroup, protectedRouterGroup)

	router.Run(":9000")
}
