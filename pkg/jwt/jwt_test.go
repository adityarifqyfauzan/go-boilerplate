package jwt

import (
	"os"
	"testing"
	"time"
)

func TestNewJWTService(t *testing.T) {
	// Test with default values
	service := NewJWTService()
	if service == nil {
		t.Fatal("Expected JWTService to be created")
	}

	// Test with environment variables
	os.Setenv("JWT_SECRET", "test-secret")
	os.Setenv("JWT_EXPIRY", "1h")

	service = NewJWTService()
	if service == nil {
		t.Fatal("Expected JWTService to be created with env vars")
	}
}

func TestGenerateToken(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("Expected token to be generated")
	}
}

func TestValidateToken(t *testing.T) {
	service := NewJWTService()

	// Generate a token
	token, err := service.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Validate the token
	claims, err := service.ValidateToken(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", claims.UserID)
	}

	if claims.Email != "test@example.com" {
		t.Errorf("Expected Email test@example.com, got %s", claims.Email)
	}

	if claims.Username != "testuser" {
		t.Errorf("Expected Username testuser, got %s", claims.Username)
	}

	if claims.Role != "user" {
		t.Errorf("Expected Role user, got %s", claims.Role)
	}
}

func TestGenerateRefreshToken(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateRefreshToken(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if token == "" {
		t.Fatal("Expected refresh token to be generated")
	}
}

func TestGenerateTokenPair(t *testing.T) {
	service := NewJWTService()

	accessToken, refreshToken, err := service.GenerateTokenPair(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if accessToken == "" {
		t.Fatal("Expected access token to be generated")
	}

	if refreshToken == "" {
		t.Fatal("Expected refresh token to be generated")
	}
}

func TestRefreshToken(t *testing.T) {
	service := NewJWTService()

	// Generate token pair
	accessToken, refreshToken, err := service.GenerateTokenPair(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Refresh the access token
	newAccessToken, err := service.RefreshToken(refreshToken)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if newAccessToken == "" {
		t.Fatal("Expected new access token to be generated")
	}

	if newAccessToken == accessToken {
		t.Fatal("Expected new access token to be different from original")
	}
}

func TestExtractUserID(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateToken(123, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	userID, err := service.ExtractUserID(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if userID != 123 {
		t.Errorf("Expected UserID 123, got %d", userID)
	}
}

func TestIsTokenExpired(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expired, err := service.IsTokenExpired(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if expired {
		t.Fatal("Expected token to not be expired")
	}
}

func TestGetTokenExpiry(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expiry, err := service.GetTokenExpiry(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if expiry == nil {
		t.Fatal("Expected expiry time to be returned")
	}

	// Check that expiry is in the future
	if time.Now().After(*expiry) {
		t.Fatal("Expected expiry time to be in the future")
	}
}

func TestValidateTokenWithoutExpiry(t *testing.T) {
	service := NewJWTService()

	token, err := service.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	claims, err := service.ValidateTokenWithoutExpiry(token)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if claims.UserID != 1 {
		t.Errorf("Expected UserID 1, got %d", claims.UserID)
	}
}

func TestInvalidToken(t *testing.T) {
	service := NewJWTService()

	// Test with invalid token
	_, err := service.ValidateToken("invalid-token")
	if err == nil {
		t.Fatal("Expected error for invalid token")
	}

	// Test with empty token
	_, err = service.ValidateToken("")
	if err == nil {
		t.Fatal("Expected error for empty token")
	}
}
