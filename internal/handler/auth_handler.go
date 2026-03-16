package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/config"
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
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, err.Error()))
		return
	}

	if err := req.Validate(); err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, err.Error()))
		return
	}

	userSession, apiErr := h.authService.Signup(req.Email, req.Username, req.Password)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	c.SetCookie(
		config.SessionCookieName,
		userSession.RefreshToken,
		config.SessionCookieMaxAgeSecs,
		config.SessionCookiePath,
		config.SessionCookieDomain,
		config.SessionCookieSecure,
		config.SessionCookieHttpOnly,
	)

	res := &dto.SignupResponse{UserSession: userSession}

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) signin(c *gin.Context) {
	req := new(dto.SigninRequest)

	if err := c.ShouldBindJSON(req); err != nil {
		status := http.StatusBadRequest
		c.JSON(status, dto.NewErrorResponse(status, err.Error()))
		return
	}

	userSession, apiErr := h.authService.Signin(req.Email, req.Password)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	c.SetCookie(
		config.SessionCookieName,
		userSession.RefreshToken,
		config.SessionCookieMaxAgeSecs,
		config.SessionCookiePath,
		config.SessionCookieDomain,
		config.SessionCookieSecure,
		config.SessionCookieHttpOnly,
	)

	res := &dto.SigninResponse{UserSession: userSession}

	c.JSON(http.StatusOK, res)
}

func (h *authHandler) refresh(c *gin.Context) {
	refreshToken, err := c.Cookie(config.SessionCookieName)
	if err != nil {
		status := http.StatusUnauthorized
		c.JSON(
			status, dto.NewErrorResponse(status, "a user session was not found"),
		)
		return
	}

	userSession, apiErr := h.authService.Refresh(refreshToken)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	c.SetCookie(
		config.SessionCookieName,
		userSession.RefreshToken,
		config.SessionCookieMaxAgeSecs,
		config.SessionCookiePath,
		config.SessionCookieDomain,
		config.SessionCookieSecure,
		config.SessionCookieHttpOnly,
	)

	res := &dto.RefreshResponse{UserSession: userSession}

	c.JSON(http.StatusOK, res)
}

func (h *authHandler) signout(c *gin.Context) {
	refreshToken, err := c.Cookie(config.SessionCookieName)
	if err != nil {
		status := http.StatusUnauthorized
		c.JSON(
			status, dto.NewErrorResponse(status, "a user session was not found"),
		)
		return
	}

	apiErr := h.authService.Signout(refreshToken)
	if apiErr != nil {
		c.JSON(apiErr.Status(), dto.NewErrorResponseFromApiError(apiErr))
		return
	}

	c.SetCookie(
		config.SessionCookieName,
		"",
		-1,
		config.SessionCookiePath,
		config.SessionCookieDomain,
		config.SessionCookieSecure,
		config.SessionCookieHttpOnly,
	)

	res := &dto.SignoutResponse{Message: "signed out successfully"}

	c.JSON(http.StatusOK, res)
}
