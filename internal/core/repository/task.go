package repository

import (
	"context"
	"errors"
	"task_management/internal/core/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
type taskRepository struct {
	database   *mongo.Database
	collection string
}
func NewTaskRepository(db *mongo.Database, collection string) models.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}
// var taskCollection *mongo.Collection

// func InitTaskCollection(client *mongo.Client) {
// 	taskCollection = client.Database("task_management_db").Collection("tasks")
// }

func (tr *taskRepository)GetTasks(c context.Context) ([]models.Task, error) {
	tasks := []models.Task{}
	cursor, err := tr.database.Collection(tr.collection).Find(c, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(c, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (tr *taskRepository) GetTask(c context.Context, taskId primitive.ObjectID) (models.Task, error) {
	var task models.Task
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	err := tr.database.Collection(tr.collection).FindOne(c, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, err
	}
	return task, nil
}

func (tr *taskRepository) AddTask(c context.Context, task models.Task) (*mongo.InsertOneResult, error) {

	return tr.database.Collection(tr.collection).InsertOne(c, task)

}

func (tr *taskRepository) UpdateTask(c context.Context, taskId primitive.ObjectID, updatedTask models.Task) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	update := bson.D{primitive.E{Key: "$set", Value: updatedTask}}
	result, err := tr.database.Collection(tr.collection).UpdateOne(c, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	return result, nil
}

func (tr *taskRepository) DeleteTask(c context.Context, taskId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	result, err := tr.database.Collection(tr.collection).DeleteOne(c, filter)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, errors.New("task not found")
	}

	return result, nil
}
