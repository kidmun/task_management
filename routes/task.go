package routes

import (
	"task_management/controllers"
	"task_management/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", midddlewares.AuthMiddleware(), controllers.GetTasksHandler)
		taskGroup.GET("/:id", midddlewares.AuthMiddleware(), controllers.GetTaskHandler)
		taskGroup.POST("/", midddlewares.AuthMiddleware(), controllers.CreateTaskHandler)
		taskGroup.PUT("/:id", midddlewares.AuthMiddleware(), controllers.UpdateTaskHandler)
		taskGroup.DELETE("/:id", midddlewares.AuthMiddleware(), midddlewares.AdminOnly(), controllers.DeleteTaskHandler)
	}

}
