package main

import (
	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/handler"
)

func main() {
  router := gin.Default()

	handler.RegisterEndpoints(router)
  
	router.Run(":9000")
}
