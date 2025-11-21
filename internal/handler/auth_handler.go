package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/dto"
	"github.com/coding-shenanigans/alchemist-service/internal/service"
)

type authHandler struct {
	authService *service.AuthService
}

func newAuthHandler(authService *service.AuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
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

	userSession, apiErr := h.authService.Signup(req.Email, req.Username, req.Password)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponse(apiErr.Error()))
		return
	}

	c.SetCookie(
		userSession.SessionCookie.Name,
		userSession.SessionCookie.Value,
		userSession.SessionCookie.MaxAge,
		userSession.SessionCookie.Path,
		userSession.SessionCookie.Domain,
		userSession.SessionCookie.Secure,
		userSession.SessionCookie.HttpOnly,
	)

	res := dto.SignupResponse{UserSession: userSession}

	c.JSON(http.StatusCreated, res)
}
