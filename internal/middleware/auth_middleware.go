package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/coding-shenanigans/alchemist-service/internal/auth"
	"github.com/coding-shenanigans/alchemist-service/internal/constant"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
)

// Extracts the authenticated user id from the Authorization header.
func extractAuthenticatedUserId(authHeader string) (int, *dto.ErrorResponse) {
	parts := strings.Fields(authHeader)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "invalid format for the Authorization header",
		)
	}

	token := parts[1]
	parsedToken, err := auth.ValidateToken(token)
	if err != nil {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "the token is not valid",
		)
	}

	kid, ok := parsedToken.Header["kid"].(string)
	if !ok || kid != constant.AccessKeyId {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "invalid token type",
		)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "failed to extract the token's claims",
		)
	}

	subClaim, ok := claims["sub"].(string)
	if !ok {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "failed to extract the sub claim",
		)
	}

	authenticatedUserId, err := strconv.Atoi(subClaim)
	if err != nil {
		return 0, dto.NewErrorResponse(
			http.StatusUnauthorized, "invalid sub claim",
		)
	}

	return authenticatedUserId, nil
}

// Verifies the access token, if provided, and sets the authenticated user for
// the request.
//
// If the access token is not provided, the request proceeds without an
// authenticated user.
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		authenticatedUserId, err := extractAuthenticatedUserId(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(err.ErrorInfo.Code, err)
			return
		}

		c.Set(constant.AuthenticatedUserId, authenticatedUserId)
		c.Next()
	}
}

// Verifies the access token and sets the authenticated user for the request.
func RequiredAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status,
				dto.NewErrorResponse(status, "the Authorization header is required"),
			)
			return
		}

		authenticatedUserId, err := extractAuthenticatedUserId(authHeader)
		if err != nil {
			c.AbortWithStatusJSON(err.ErrorInfo.Code, err)
			return
		}

		c.Set(constant.AuthenticatedUserId, authenticatedUserId)
		c.Next()
	}
}
