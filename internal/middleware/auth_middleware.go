package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/coding-shenanigans/alchemist-service/internal/auth"
	"github.com/coding-shenanigans/alchemist-service/internal/constant"
	"github.com/coding-shenanigans/alchemist-service/internal/dto"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.NewErrorResponse("the Authorization header is required"),
			)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.NewErrorResponse("invalid format for the Authorization header"),
			)
			return
		}

		token := parts[1]
		parsedToken, err := auth.ValidateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.NewErrorResponse("the token is not valid"),
			)
			return
		}

		kid, ok := parsedToken.Header["kid"].(string)
		if !ok || kid != constant.AccessKeyId {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				dto.NewErrorResponse("invalid token type"),
			)
			return
		}

		c.Next()
	}
}
