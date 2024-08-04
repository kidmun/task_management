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

func SetupUserRoutes(router *gin.Engine, db *mongo.Database, timeout time.Duration) {
	ur := repository.NewUserRepository(db, models.CollectionUser)
	uc := &controllers.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, timeout),
	}
	
	userGroup := router.Group("/auth")
	{
		userGroup.POST("/register", uc.RegisterUserHandler)
		userGroup.POST("/login", uc.LoginUserHandler)
		
		userGroup.POST("/admin_register", midddlewares.AuthMiddleware(), midddlewares.AdminOnly(), uc.RegisterAdminHandler)
	
}
}
