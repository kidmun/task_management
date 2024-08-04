package controllers

import (
	"net/http"
	"task_management/internal/core/models"

	"github.com/gin-gonic/gin"
)
type UserController struct {
	UserUsecase models.UserUsecase
}

func (uc *UserController) RegisterUserHandler(ctx *gin.Context) {
	var userInput models.UserInput
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse the given data"})
		return
	}
	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}
	createdUser, err := uc.UserUsecase.RegisterUser(ctx, user)
	if err != nil {
		if err.Error() == "username already taken" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, createdUser)

}
func (uc *UserController) RegisterAdminHandler(ctx *gin.Context) {
	var userInput models.UserInput
	err := ctx.ShouldBindJSON(&userInput)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse the given data"})
		return
	}
	user := models.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}
	createdUser, err := uc.UserUsecase.RegisterAdmin(ctx, user)
	if err != nil {
		if err.Error() == "username already taken" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, createdUser)

}

func (uc *UserController) LoginUserHandler(ctx *gin.Context) {
	var user models.UserInput
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse the given data"})
		return
	}
	tokenString, err := uc.UserUsecase.LoginUser(ctx, user)
	if err != nil {
		if err.Error() == "wrong Credentials" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
