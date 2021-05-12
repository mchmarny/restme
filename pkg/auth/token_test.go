package auth

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	testAuthTknPath = "../../test/test.token"
	testKeyPath     = "../../test/test.key"
	testKeyContent  = "test-key-content"
)

func writeFile(path, content string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return errors.Wrapf(err, "error creating file: %s", path)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return err
	}
	return nil
}

func createTestKeyFile() error {
	return writeFile(testKeyPath, testKeyContent)
}

func TestToken(t *testing.T) {
	if err := createTestKeyFile(); err != nil {
		t.Fatalf("error creating test key: %v", err)
	}

	testSecret, err := ioutil.ReadFile(testKeyPath)
	if err != nil {
		t.Fatalf("error reading test key %s: %v", testKeyPath, err)
	}

	tokenStr, err := MakeJWT(testSecret, "test", "user@domain.com", "8760h")
	if err != nil {
		t.Fatalf("error making JWT: %v", err)
	}
	assert.NotEmpty(t, tokenStr)
	t.Logf("token: %s", tokenStr)

	token, err := ParseJWT(testSecret, tokenStr)
	if err != nil {
		t.Fatalf("error parsing JWT: %v", err)
	}
	assert.NotNil(t, token)
	assert.NotEmpty(t, token.Email)
	assert.Nil(t, token.Valid())
	assert.True(t, token.VerifyExpiresAt(time.Now().Unix(), true))

	if err := writeFile(testAuthTknPath, tokenStr); err != nil {
		t.Fatalf("error writing JWT token to %s: %v", testAuthTknPath, err)
	}
}
