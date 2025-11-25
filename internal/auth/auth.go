package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/coding-shenanigans/alchemist-service/internal/config"
)

const (
	accessKeyId  = "Access"
	refreshKeyId = "Refresh"
)

// Generates an access token.
func GenerateAccessToken(userId int) (string, error) {
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
func GenerateRefreshToken(userId int) (string, error) {
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
func CreateSessionCookie(userId int) (*http.Cookie, error) {
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

// Checks if the token is valid.
func ValidateToken(token string) (*jwt.Token, error) {
	keyFunc := func(parsedToken *jwt.Token) (any, error) {
		if _, ok := parsedToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(
				"unexpected signing method: %v", parsedToken.Header["alg"],
			)
		}

		kid, ok := parsedToken.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("failed to get the token's key identifier")
		}

		switch kid {
		case accessKeyId:
			return []byte(config.AccessTokenSecret), nil
		case refreshKeyId:
			return []byte(config.RefreshTokenSecret), nil
		default:
			return nil, fmt.Errorf("invalid key identifier: %v", kid)
		}
	}

	return jwt.Parse(token, keyFunc)
}

// Generates an authentication token.
func generateToken(
	keyId string, secretKey string, durationSecs int, userId int,
) (string, error) {
	duration := time.Duration(durationSecs) * time.Second
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": strconv.Itoa(userId),
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
