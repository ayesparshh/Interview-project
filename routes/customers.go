package routes

import (
	"net/http"

	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware checks if the user has the required role
func RoleMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole := c.Query("role")
		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			c.Abort()
			return
		}
		c.Next()
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
