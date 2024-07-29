package routes

import (
	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

func BillingRoutes(router *gin.Engine, db db.DB) {

	// sales and accountant have access to billings
	salesGroup := router.Group("/billings")
	salesGroup.Use(RoleMiddleware("sales", "accountant"))
	{
		salesGroup.GET("", db.GetBillings)
	}

	// sales have post access to billings
	salesGroup.Use(RoleMiddleware("sales"))
	{
		salesGroup.POST("", db.GetBillings)
	}
}
