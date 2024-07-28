package routes

import (
	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, db db.DB) {
	// Group for sales role with read/write access
	administratorGroup := router.Group("/users")
	administratorGroup.Use(RoleMiddleware("administrator"))
	{
		administratorGroup.GET("", db.GetBillings)
		administratorGroup.POST("", db.CreateBilling)
	}
}
