package routes

import (
	"github.com/ayesparshh/db"
	"github.com/gin-gonic/gin"
)

func PayrollRoutes(router *gin.Engine, db db.DB) {
	// Group for hr role with read/write access
	salesGroup := router.Group("/payrolls")
	salesGroup.Use(RoleMiddleware("hr"))
	{
		salesGroup.GET("", db.GetPayrolls)
		salesGroup.POST("", db.CreatePayroll)
	}

	// Group for accountant role with view access
	accountantGroup := router.Group("/payrolls")
	accountantGroup.Use(RoleMiddleware("accountant"))
	{
		accountantGroup.GET("", db.GetPayrolls)
	}
}
