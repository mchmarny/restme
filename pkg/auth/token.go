package auth

import (
	"os"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

// TokenClaims represents API invocation claim.
type TokenClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// MakeJWT creates JWT token valid until TTL for username and signs it using secret.
func MakeJWT(secret []byte, issuer, email, ttl string) (string, error) {
	if secret == nil {
		return "", errors.New("secret required")
	}
	if !IsEmailValid(email) {
		return "", errors.New("invalid email address")
	}

	tokenTTL, err := time.ParseDuration(ttl)
	if err != nil {
		return "", errors.New("invalid ttl")
	}

	claims := TokenClaims{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errors.Wrap(err, "error generating token")
	}
	return tokenString, nil
}

func ParseJWT(secret []byte, userToken string) (*TokenClaims, error) {
	if secret == nil {
		return nil, errors.New("secret required")
	}
	if userToken == "" {
		return nil, errors.New("userToken required")
	}

	token, err := jwt.ParseWithClaims(userToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "error parsing token")
	}

	return token.Claims.(*TokenClaims), nil
}

// IsEmailValid checks if the email provided passes the required structure and length.
func IsEmailValid(e string) bool {
	if len(e) < 3 || len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func getKey(path string) ([]byte, error) {
	if path == "" {
		return nil, errors.New("path required")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.Wrapf(err, "error reading key file: %s", path)
	}

	return b, nil
}
