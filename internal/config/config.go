package config

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
)

const (
	portKey = "PORT"

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

	allowedOriginsKey = "ALLOWED_ORIGINS"
)

// TODO: Change this to a struct instead of global variables.
var (
	Port string

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string

	AccessTokenSecret        string
	AccessTokenDurationSecs  int
	RefreshTokenSecret       string
	RefreshTokenDurationSecs int
	SessionCookieName        string
	SessionCookieMaxAgeSecs  int
	SessionCookiePath        string
	SessionCookieDomain      string
	SessionCookieSecure      bool
	SessionCookieHttpOnly    bool

	allowedOrigins []string
)

// Reads all environment variables.
func Load() {
	Port = GetEnvStrOrDefault(portKey, "9000")

	DbHost = GetEnvStr(dbHostKey)
	DbPort = GetEnvInt(dbPortKey)
	DbUser = GetEnvStr(dbUserKey)
	DbPassword = GetEnvStr(dbPasswordKey)
	DbName = GetEnvStr(dbNameKey)
	DbSslMode = GetEnvStr(dbSslModeKey)

	AccessTokenSecret = GetEnvStr(accessTokenSecretKey)
	AccessTokenDurationSecs = GetEnvInt(accessTokenDurationSecsKey)
	RefreshTokenSecret = GetEnvStr(refreshTokenSecretKey)
	RefreshTokenDurationSecs = GetEnvInt(refreshTokenDurationSecsKey)
	SessionCookieName = GetEnvStr(sessionCookieNameKey)
	SessionCookieMaxAgeSecs = GetEnvInt(sessionCookieMaxAgeSecsKey)
	SessionCookiePath = GetEnvStr(sessionCookiePathKey)
	// Allow the cookie domain to be empty for host-only cookies
	SessionCookieDomain = GetEnvStrOrDefault(sessionCookieDomainKey, "")
	SessionCookieSecure = GetEnvBool(sessionCookieSecureKey)
	SessionCookieHttpOnly = GetEnvBool(sessionCookieHttpOnlyKey)

	for origin := range strings.SplitSeq(GetEnvStr(allowedOriginsKey), ",") {
		if origin = strings.TrimSpace(origin); origin != "" {
			allowedOrigins = append(allowedOrigins, origin)
		}
	}
}

// Gets a required environment variable.
func GetEnvStr(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("invalid config value %q for key %q", value, key)
	}

	return value
}

// Gets an optional environment variable or falls back to a default value.
func GetEnvStrOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
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

// Creates a custom CORS config.
func GetCorsConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}
