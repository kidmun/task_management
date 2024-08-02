package midddlewares

import (
	
	"net/http"
	"strings"
	"task_management/config"
	"task_management/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			context.Abort()
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		claims := &services.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetEnv("Jwt_Secret")), nil
		})
	
		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}
		context.Set("username", claims.Username)
		context.Set("role", claims.Role)
		context.Next()
	}
}
func AdminOnly() gin.HandlerFunc {
	return func(context *gin.Context) {
		role, exists := context.Get("role")
		if !exists || role != "Admin" {
			context.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			context.Abort()
			return
		}
		context.Next()
	}
}
