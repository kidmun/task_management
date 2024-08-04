package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "InProgress"
	Done       Status = "Done"
)
const (
	CollectionTask = "tasks"
) 
type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	DueDate     time.Time          `json:"due_date" binding:"required"`
	Status      Status             `json:"status"`
}
type TaskRepository interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTask(c context.Context, taskId primitive.ObjectID) (Task, error)
	AddTask(c context.Context, task Task) (*mongo.InsertOneResult, error)
	UpdateTask(c context.Context, taskId primitive.ObjectID, updatedTask Task) (*mongo.UpdateResult, error) 
	DeleteTask(c context.Context, taskId primitive.ObjectID) (*mongo.DeleteResult, error)
}

type TaskUsecase interface {
	GetTasks(c context.Context) ([]Task, error)
	GetTask(c context.Context, taskId primitive.ObjectID) (Task, error)
	AddTask(c context.Context, task Task) (*mongo.InsertOneResult, error)
	UpdateTask(c context.Context, taskId primitive.ObjectID, updatedTask Task) (*mongo.UpdateResult, error) 
	DeleteTask(c context.Context, taskId primitive.ObjectID) (*mongo.DeleteResult, error)
	
}