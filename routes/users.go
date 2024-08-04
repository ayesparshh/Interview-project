package routes

import (
	"github.com/ayesparshh/db"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, db db.DB) {

	// administrator have access to users
	administratorGroup := router.Group("/users")
	administratorGroup.Use(RoleMiddleware("administrator"))
	{
		administratorGroup.GET("", db.GetUsers)
		administratorGroup.POST("", db.CreateUser)
	}
}
