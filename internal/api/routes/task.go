package routes

import (
	"task_management/internal/api/controllers"
	midddlewares "task_management/internal/api/middlewares"
	"task_management/internal/core/models"
	"task_management/internal/core/repository"
	"task_management/internal/core/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupTaskRoutes(router *gin.Engine, db *mongo.Database, timeout time.Duration) {
	tr := repository.NewTaskRepository(db, models.CollectionTask)
	tc := &controllers.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	taskGroup := router.Group("/tasks")
	{
		taskGroup.GET("/", midddlewares.AuthMiddleware(), tc.GetTasksHandler)
		taskGroup.GET("/:id", midddlewares.AuthMiddleware(), tc.GetTaskHandler)
		taskGroup.POST("/", midddlewares.AuthMiddleware(), tc.CreateTaskHandler)
		taskGroup.PUT("/:id", midddlewares.AuthMiddleware(), tc.UpdateTaskHandler)
		taskGroup.DELETE("/:id", midddlewares.AuthMiddleware(), midddlewares.AdminOnly(), tc.DeleteTaskHandler)
	}

}
