package routes

import (
	"task_management/controllers"
	"task_management/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/register", controllers.RegisterUserHandler)
		userGroup.POST("/login", controllers.LoginUserHandler)
		userGroup.POST("/admin_register", midddlewares.AuthMiddleware(), midddlewares.AdminOnly(), controllers.LoginUserHandler)
	}
}
