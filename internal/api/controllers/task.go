package controllers

import (
	"errors"
	"net/http"
	"task_management/internal/core/models"
	"task_management/pkg/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type TaskController struct {
	TaskUsecase models.TaskUsecase
}



func (tc *TaskController) GetTasksHandler(ctx *gin.Context) {
	tasks, err := tc.TaskUsecase.GetTasks(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
func (tc *TaskController) GetTaskHandler(ctx *gin.Context) {

	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse task id."})
		return
	}
	tasks, err := tc.TaskUsecase.GetTask(ctx, id)
	if err != nil {
		if errors.Is(err, errors.New("task not found")) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) CreateTaskHandler(ctx *gin.Context) {
	var task models.Task
	err := ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if task.Status == "" {
		task.Status = models.Pending
	} else if !utils.IsValidStatus(task.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task status"})
		return
	}
	result, err := tc.TaskUsecase.AddTask(ctx, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task created successfully",
		"result":  result,
	})

}
func (tc *TaskController) UpdateTaskHandler(ctx *gin.Context) {
	var task models.Task
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse task id."})
		return
	}
	err = ctx.ShouldBindJSON(&task)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse the given data"})
		return
	}
	if task.Status == "" {
		task.Status = models.Pending
	} else if !utils.IsValidStatus(task.Status) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid task status"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "couldn't parse the given data"})
		return
	}
	result, err := tc.TaskUsecase.UpdateTask(ctx, id, task)
	if err != nil {
		if errors.Is(err, errors.New("task not found")) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"result":  result,
	})
}

func (tc *TaskController) DeleteTaskHandler(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse task id."})
		return
	}
	result, err := tc.TaskUsecase.DeleteTask(ctx, id)
	if err != nil {
		if errors.Is(err, errors.New("task not found")) {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task deleted successfully",
		"result":  result,
	})

}
