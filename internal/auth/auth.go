package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/coding-shenanigans/alchemist-service/internal/config"
)

const (
	accessKeyId  = "Access"
	refreshKeyId = "Refresh"
)

// Generates an access token.
func GenerateAccessToken(userId string) (string, error) {
	token, err := generateToken(
		accessKeyId,
		config.AccessTokenSecret,
		config.AccessTokenDurationSecs,
		userId,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Generates a refresh token.
func GenerateRefreshToken(userId string) (string, error) {
	token, err := generateToken(
		refreshKeyId,
		config.RefreshTokenSecret,
		config.RefreshTokenDurationSecs,
		userId,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Creates a session cookie.
func CreateSessionCookie(userId string) (*http.Cookie, error) {
	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     config.SessionCookieName,
		Value:    refreshToken,
		MaxAge:   config.SessionCookieMaxAgeSecs,
		Path:     config.SessionCookiePath,
		Domain:   config.SessionCookieDomain,
		Secure:   config.SessionCookieSecure,
		HttpOnly: config.SessionCookieHttpOnly,
	}, nil
}

// Generates an authentication token.
func generateToken(
	keyId string, secretKey string, durationSecs int, userId string,
) (string, error) {
	duration := time.Duration(durationSecs) * time.Second
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": userId,
		"iat": issuedAt.Unix(),
		"exp": expiresAt.Unix(),
	})
	token.Header["kid"] = keyId

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		// TODO: log error
		return "", fmt.Errorf("failed to sign the token")
	}

	return signedToken, nil
}
