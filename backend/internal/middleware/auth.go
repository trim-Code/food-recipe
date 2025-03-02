package middleware

import (
	"food-recipe-app/backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware to validate JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenStr := c.GetHeader("Authorization")
        if tokenStr == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
            c.Abort()
            return
        }

        // Remove "Bearer " prefix if present
        if len(tokenStr) > 7 && tokenStr[:7] == "Bearer " {
            tokenStr = tokenStr[7:]
        }

        username, err := utils.ValidateJWT(tokenStr)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
            c.Abort()
            return
        }

        c.Set("username", username)
        c.Next()
    }
}
