package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type opsHandler struct {}

func newOpsHandler() *opsHandler {
	return &opsHandler{}
}

func (h *opsHandler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ "status": "OK" })
}
