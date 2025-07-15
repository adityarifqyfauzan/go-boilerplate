package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adityarifqyfauzan/go-boilerplate/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

func TestAuthMiddleware_ValidToken(t *testing.T) {
	router := setupTestRouter()
	jwtService := jwt.NewJWTService()

	// Generate a valid token
	token, err := jwtService.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		userID, exists := GetUserID(c)
		if !exists {
			t.Error("Expected user_id to be set in context")
			return
		}
		if userID != 1 {
			t.Errorf("Expected user_id 1, got %d", userID)
		}

		userEmail, exists := GetUserEmail(c)
		if !exists {
			t.Error("Expected user_email to be set in context")
			return
		}
		if userEmail != "test@example.com" {
			t.Errorf("Expected user_email test@example.com, got %s", userEmail)
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestAuthMiddleware_NoToken(t *testing.T) {
	router := setupTestRouter()

	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidToken(t *testing.T) {
	router := setupTestRouter()

	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer invalid-token")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestAuthMiddleware_InvalidHeaderFormat(t *testing.T) {
	router := setupTestRouter()

	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "InvalidFormat token")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}

func TestOptionalAuthMiddleware_WithToken(t *testing.T) {
	router := setupTestRouter()
	jwtService := jwt.NewJWTService()

	// Generate a valid token
	token, err := jwtService.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	router.Use(OptionalAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		userID, exists := GetUserID(c)
		if !exists {
			t.Error("Expected user_id to be set in context")
			return
		}
		if userID != 1 {
			t.Errorf("Expected user_id 1, got %d", userID)
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestOptionalAuthMiddleware_WithoutToken(t *testing.T) {
	router := setupTestRouter()

	router.Use(OptionalAuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		// Should not have user information in context
		userID, exists := GetUserID(c)
		if exists {
			t.Errorf("Expected user_id to not be set in context, got %d", userID)
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestRoleMiddleware_ValidRole(t *testing.T) {
	router := setupTestRouter()
	jwtService := jwt.NewJWTService()

	// Generate a token with admin role
	token, err := jwtService.GenerateToken(1, "admin@example.com", "admin", "admin")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	router.Use(AuthMiddleware(), AdminMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestRoleMiddleware_InvalidRole(t *testing.T) {
	router := setupTestRouter()
	jwtService := jwt.NewJWTService()

	// Generate a token with user role
	token, err := jwtService.GenerateToken(1, "user@example.com", "user", "user")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	router.Use(AuthMiddleware(), AdminMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Errorf("Expected status 403, got %d", w.Code)
	}
}

func TestGetUserClaims(t *testing.T) {
	router := setupTestRouter()
	jwtService := jwt.NewJWTService()

	// Generate a valid token
	token, err := jwtService.GenerateToken(1, "test@example.com", "testuser", "user")
	if err != nil {
		t.Fatalf("Failed to generate token: %v", err)
	}

	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		claims, exists := GetUserClaims(c)
		if !exists {
			t.Error("Expected user_claims to be set in context")
			return
		}
		if claims.UserID != 1 {
			t.Errorf("Expected UserID 1, got %d", claims.UserID)
		}
		if claims.Email != "test@example.com" {
			t.Errorf("Expected Email test@example.com, got %s", claims.Email)
		}

		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
