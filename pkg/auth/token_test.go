package auth

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/mchmarny/restme/pkg/fileutil"
	"github.com/stretchr/testify/assert"
)

const (
	testAuthTknPath = "../../test/test.token"
	testKeyPath     = "../../test/test.key"
	testKeyContent  = "test-key-content"
)

func createTestKeyFile() error {
	return fileutil.WriteFile(testKeyPath, testKeyContent)
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

	if err := fileutil.WriteFile(testAuthTknPath, tokenStr); err != nil {
		t.Fatalf("error writing JWT token to %s: %v", testAuthTknPath, err)
	}
}
