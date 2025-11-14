package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/dto"
)

type authHandler struct{}

func newAuthHandler() *authHandler {
	return &authHandler{}
}

func (h *authHandler) signup(c *gin.Context) {
	req := new(dto.SignupRequest)

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
		return
	}

	// TODO: create user
	// TODO: create user session
	c.JSON(http.StatusNotImplemented, gin.H{"error": "not implemented"})
}
