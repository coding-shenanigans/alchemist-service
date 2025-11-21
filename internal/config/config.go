package config

import (
	"log"
	"os"
	"strconv"
)

const (
	dbHostKey     = "DB_HOST"
	dbPortKey     = "DB_PORT"
	dbUserKey     = "DB_USER"
	dbPasswordKey = "DB_PASSWORD"
	dbNameKey     = "DB_NAME"
	dbSslModeKey  = "DB_SSL_MODE"

	accessTokenSecretKey        = "ACCESS_TOKEN_SECRET"
	accessTokenDurationSecsKey  = "ACCESS_TOKEN_DURATION_SECS"
	refreshTokenSecretKey       = "REFRESH_TOKEN_SECRET"
	refreshTokenDurationSecsKey = "REFRESH_TOKEN_DURATION_SECS"
	sessionCookieNameKey        = "SESSION_COOKIE_NAME"
	sessionCookieMaxAgeSecsKey  = "SESSION_COOKIE_MAX_AGE_SECS"
	sessionCookiePathKey        = "SESSION_COOKIE_PATH"
	sessionCookieDomainKey      = "SESSION_COOKIE_DOMAIN"
	sessionCookieSecureKey      = "SESSION_COOKIE_SECURE"
	sessionCookieHttpOnlyKey    = "SESSION_COOKIE_HTTP_ONLY"
)

var (
	DbHost     = GetEnvStr(dbHostKey)
	DbPort     = GetEnvInt(dbPortKey)
	DbUser     = GetEnvStr(dbUserKey)
	DbPassword = GetEnvStr(dbPasswordKey)
	DbName     = GetEnvStr(dbNameKey)
	DbSslMode  = GetEnvStr(dbSslModeKey)

	AccessTokenSecret        = GetEnvStr(accessTokenSecretKey)
	AccessTokenDurationSecs  = GetEnvInt(accessTokenDurationSecsKey)
	RefreshTokenSecret       = GetEnvStr(refreshTokenSecretKey)
	RefreshTokenDurationSecs = GetEnvInt(refreshTokenDurationSecsKey)
	SessionCookieName        = GetEnvStr(sessionCookieNameKey)
	SessionCookieMaxAgeSecs  = GetEnvInt(sessionCookieMaxAgeSecsKey)
	SessionCookiePath        = GetEnvStr(sessionCookiePathKey)
	SessionCookieDomain      = GetEnvStr(sessionCookieDomainKey)
	SessionCookieSecure      = GetEnvBool(sessionCookieSecureKey)
	SessionCookieHttpOnly    = GetEnvBool(sessionCookieHttpOnlyKey)
)

// Gets a required environment variable.
func GetEnvStr(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("invalid config value %q for key %q", value, key)
	}

	return value
}

// Gets a required environment variable and converts it to an integer.
func GetEnvInt(key string) int {
	valueStr := os.Getenv(key)

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("invalid config value %q for key %q: %v", valueStr, key, err)
	}

	return value
}

// Gets a required environment variable and converts it to a boolean.
func GetEnvBool(key string) bool {
	valueStr := os.Getenv(key)

	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		log.Fatalf("invalid config value %q for key %q: %v", valueStr, key, err)
	}

	return value
}
