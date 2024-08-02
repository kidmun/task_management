package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	Pending    Status = "Pending"
	InProgress Status = "InProgress"
	Done       Status = "Done"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	DueDate     time.Time          `json:"due_date" binding:"required"`
	Status      Status             `json:"status"`
}
