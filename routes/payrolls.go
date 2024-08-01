package routes

import (
	"github.com/ayesparshh/db"
	"github.com/gin-gonic/gin"
)

func PayrollRoutes(router *gin.Engine, db db.DB) {

	// hr and accountant have access to payrolls
	salesGroup := router.Group("/payrolls")
	salesGroup.Use(RoleMiddleware("hr", "accountant"))
	{
		salesGroup.GET("", db.GetPayrolls)
	}

	// hr have post access to payrolls
	salesGroup.Use(RoleMiddleware("hr"))
	{
		salesGroup.POST("", db.GetPayrolls)
	}
}
