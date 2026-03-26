package jwtTokens_test

import (
	"testing"

	"github.com/ViitoJooj/clown-crm/pkg/jwtTokens"
	"github.com/golang-jwt/jwt/v4"
)

func TestGenerateAndValidateToken(t *testing.T) {
	// Generate a token
	tokenString, err := jwtTokens.GenerateToken("test-uuid")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	// Validate the token
	token, err := jwtTokens.ValidateToken(tokenString)
	if err != nil {
		t.Fatalf("Failed to validate token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Fatalf("Invalid token")
	}
	userID, ok := claims["user_id"].(string)
	if !ok {
		t.Fatalf("Invalid user ID in token claims")
	}
	if userID != "test-uuid" {
		t.Fatalf("Expected user ID 'test-uuid', got '%s'", userID)
	}
}
