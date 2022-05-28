package auth

import (
	"fibonacci-spiral-matrix-go/internal/auth"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	generatedToken, err := jwtWrapper.GenerateToken("jwt@email.com", "testUser", "USER")
	assert.NoError(t, err)
	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	encodedToken := os.Getenv("testToken")

	jwtWrapper := auth.JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "AuthService",
	}

	claims, err := jwtWrapper.ValidateToken(encodedToken)
	assert.NoError(t, err)
	assert.Equal(t, "jwt@email.com", claims.Email)
	assert.Equal(t, "testUser", claims.Username)
	assert.Equal(t, "USER", claims.Role)
	assert.Equal(t, "AuthService", claims.Issuer)
}
