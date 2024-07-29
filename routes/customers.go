package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRole := claims["role"].(string)
			for _, role := range roles {
				if userRole == role {
					c.Next()
					return
				}
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}

// RegisterCustomersRoutes registers the customers routes
func CustomersRoutes(router *gin.Engine, db db.DB) {
	salesGroup := router.Group("/customers")
	salesGroup.Use(RoleMiddleware("sales"))
	{
		salesGroup.GET("", db.GetCustomers)
		salesGroup.POST("", db.CreateCustomer)
	}
}
