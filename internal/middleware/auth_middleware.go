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

func AuthMiddleware() gin.HandlerFunc {
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

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status,
				dto.NewErrorResponse(
					status, "invalid format for the Authorization header",
				),
			)
			return
		}

		token := parts[1]
		parsedToken, err := auth.ValidateToken(token)
		if err != nil {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status, dto.NewErrorResponse(status, "the token is not valid"),
			)
			return
		}

		kid, ok := parsedToken.Header["kid"].(string)
		if !ok || kid != constant.AccessKeyId {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status, dto.NewErrorResponse(status, "invalid token type"),
			)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status, dto.NewErrorResponse(
					status, "failed to extract the token's claims",
				),
			)
			return
		}

		subClaim, ok := claims["sub"].(string)
		if !ok {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status, dto.NewErrorResponse(
					status, "failed to extract the sub claim",
				),
			)
		}

		userId, err := strconv.Atoi(subClaim)
		if err != nil {
			status := http.StatusUnauthorized
			c.AbortWithStatusJSON(
				status, dto.NewErrorResponse(status, "invalid sub claim"),
			)
		}

		c.Set(constant.AuthenticatedUserId, userId)

		c.Next()
	}
}
