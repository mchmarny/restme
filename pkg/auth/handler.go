package auth

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mchmarny/restme/pkg/log"
	"github.com/pkg/errors"
)

const (
	authUsernameCookie = "username"
	authTokenHeader    = "Authorization"
	expectedTokenParts = 2
)

var (
	sessionCookieAge = 5 * 60 // maxSessionAge in secs
)

// NewService creates a new EchoService instance.
func NewTokenAuthenticator(path string, logger *log.Logger) (*TokenAuthenticator, error) {
	key, err := getKey(path)
	if err != nil {
		return nil, errors.Wrap(err, "error getting key")
	}

	return &TokenAuthenticator{
		logger: logger,
		key:    key,
	}, nil
}

// TokenAuthenticator provides message echo service.
type TokenAuthenticator struct {
	logger *log.Logger
	key    []byte
}

// Authenticate is a token authentication midleware.
func (a *TokenAuthenticator) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token string from header
		tokenHeader := strings.TrimSpace(c.GetHeader(authTokenHeader))
		tokenParts := strings.Split(tokenHeader, " ")
		if len(tokenParts) != expectedTokenParts {
			a.logger.Errorf("invalid token format '%s'", tokenHeader)
			makeAuthError(c, "Invalid token format")
			return
		}

		tokenVal := tokenParts[1]
		if tokenVal == "" {
			makeAuthError(c, "User not authenticated")
			return
		}

		// token from token string
		token, err := ParseJWT(a.key, tokenVal)
		if err != nil {
			a.logger.Errorf("error parsing token '%s': %v", tokenVal, err)
			makeAuthError(c, "Token parsing error")
			return
		}

		// validate token
		if (token.Valid() != nil) || (!token.VerifyExpiresAt(time.Now().Unix(), true)) {
			a.logger.Errorf("token invalid or expired '%+v'", token)
			makeAuthError(c, "Invalid token")
			return
		}

		// set session cookie
		c.SetCookie(authUsernameCookie, token.Email,
			sessionCookieAge, "/", c.Request.Host, false, true)

		c.Next()
	}
}

func makeAuthError(c *gin.Context, msg string) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": msg,
		"status":  "Unauthorized",
	})
	c.Abort()
}
