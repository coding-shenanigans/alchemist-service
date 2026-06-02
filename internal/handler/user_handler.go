package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

type userHandler struct {
	userService *service.UserService
}

func newUserHandler(userService *service.UserService) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) getUserProfile(c *gin.Context) {
	username := c.Param("username")

	user, apiErr := h.userService.GetUserProfile(username)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	res := &dto.GetUserProfileResponse{Username: user.Username}
	c.JSON(http.StatusOK, res)
}
