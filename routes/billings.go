package routes

import (
	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

func BillingRoutes(router *gin.Engine, db db.DB) {
	// Group for sales role with read/write access
	salesGroup := router.Group("/billings")
	salesGroup.Use(RoleMiddleware("sales"))
	{
		salesGroup.GET("", db.GetBillings)
		salesGroup.POST("", db.CreateBilling)
	}

	// Group for accountant role with view access
	accountantGroup := router.Group("/billings")
	accountantGroup.Use(RoleMiddleware("accountant"))
	{
		accountantGroup.GET("", db.GetBillings)
	}
}
