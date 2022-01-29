package main

import (
	b64 "encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const (
	headerAPIInvoker   = "x-api-invoker"
	expectedTokenParts = 2
)

// APITokenRequired is a authentication midleware
func APITokenRequired(tokens map[string]string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token not provided",
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}

		owner, tokenVal, err := getTokenPath(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}

		expectedTokenVal, ok := tokens[owner]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token Not Found",
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}

		if tokenVal != expectedTokenVal {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid Token",
				"status":  "Unauthorized",
			})
			c.Abort()
			return
		}

		// valid token
		c.Set(headerAPIInvoker, owner)

		c.Next()
	}
}

func getTokenPath(token string) (owner, userToken string, err error) {
	if token == "" {
		return "", "", errors.New("token is empty")
	}

	tokenStr, err := b64.StdEncoding.DecodeString(token)
	if err != nil {
		return "", "", errors.Wrap(err, "error decoding token")
	}

	tokenParts := strings.Split(string(tokenStr), ":")
	if len(tokenParts) != expectedTokenParts {
		return "", "", errors.New("token is in an invalid format")
	}

	return tokenParts[0], tokenParts[1], nil
}
