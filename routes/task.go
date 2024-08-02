package routes

import (
	"task_management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTaskRoutes(router *gin.Engine) {
	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", controllers.GetTasksHandler)
		taskGroup.GET("/:id", controllers.GetTaskHandler)
		taskGroup.POST("/", controllers.CreateTaskHandler)
		taskGroup.PUT("/:id", controllers.UpdateTaskHandler)
		taskGroup.DELETE("/:id", controllers.DeleteTaskHandler)
	}
}
