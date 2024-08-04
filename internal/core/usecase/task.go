package usecase

import (
	"context"
	"task_management/internal/core/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskUsecase struct {
	taskRepository models.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository models.TaskRepository, timeout time.Duration) models.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}
func (tu *taskUsecase) GetTasks(c context.Context) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetTasks(ctx)
}
func  (tu *taskUsecase)  GetTask(c context.Context, taskId primitive.ObjectID) (models.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.GetTask(ctx, taskId)
}
func (tu *taskUsecase) AddTask(c context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.AddTask(ctx, task)

}
func (tu *taskUsecase) UpdateTask(c context.Context, taskId primitive.ObjectID, updatedTask models.Task) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.UpdateTask(ctx, taskId, updatedTask)
}

func (tu *taskUsecase) DeleteTask(c context.Context, taskId primitive.ObjectID) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.DeleteTask(ctx, taskId)
}
