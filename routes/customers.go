package routes

import (
	"net/http"

	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Query("role")
		for _, role := range roles {
			if userRole == role {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
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
