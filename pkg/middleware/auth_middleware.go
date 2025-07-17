package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/adityarifqyfauzan/go-boilerplate/pkg/jwt"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT tokens and sets user information in context
func AuthMiddleware() gin.HandlerFunc {
	jwtService := jwt.NewJWTService()

	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid authorization header format. Expected 'Bearer <token>'",
			})
			c.Abort()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Token is required",
			})
			c.Abort()
			return
		}

		// Validate the token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// OptionalAuthMiddleware validates JWT tokens but doesn't abort if token is missing
func OptionalAuthMiddleware() gin.HandlerFunc {
	jwtService := jwt.NewJWTService()

	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// Check if the header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.Next()
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			c.Next()
			return
		}

		// Validate the token
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Set("user_claims", claims)

		c.Next()
	}
}

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "User role not found",
			})
			c.Abort()
			return
		}

		roles, ok := userRole.([]string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "Invalid user role format",
			})
			c.Abort()
			return
		}

		count := 0
		for _, rr := range requiredRoles {
			found := slices.Contains(roles, rr)
			if found {
				count++
			}
		}

		if count == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"error":   "Unauthorized",
				"message": "User does not have the required role",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminMiddleware checks if the user is an admin
func AdminMiddleware() gin.HandlerFunc {
	return RoleMiddleware("admin")
}

// UserMiddleware checks if the user is a regular user
func UserMiddleware() gin.HandlerFunc {
	return RoleMiddleware("user")
}

// GetUserID retrieves user ID from context
func GetUserID(c *gin.Context) (int, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}
	return userID.(int), true
}

// GetUserEmail retrieves user email from context
func GetUserEmail(c *gin.Context) (string, bool) {
	userEmail, exists := c.Get("user_email")
	if !exists {
		return "", false
	}
	return userEmail.(string), true
}

// GetUserRole retrieves the first user role from context
func GetUserRole(c *gin.Context) (string, bool) {
	userRoles, exists := c.Get("roles")
	if !exists {
		return "", false
	}
	roles, ok := userRoles.([]string)
	if !ok || len(roles) == 0 {
		return "", false
	}
	return roles[0], true
}

// GetUserRoles retrieves user roles from context
func GetUserRoles(c *gin.Context) ([]string, bool) {
	userRoles, exists := c.Get("roles")
	if !exists {
		return nil, false
	}
	roles, ok := userRoles.([]string)
	if !ok {
		return nil, false
	}
	return roles, true
}

// GetUserClaims retrieves user claims from context
func GetUserClaims(c *gin.Context) (*jwt.Claims, bool) {
	userClaims, exists := c.Get("user_claims")
	if !exists {
		return nil, false
	}
	return userClaims.(*jwt.Claims), true
}
