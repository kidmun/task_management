package services

import (
	"context"
	"errors"
	"task_management/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

func InitTaskCollection(client *mongo.Client) {
	taskCollection = client.Database("task_management_db").Collection("tasks")
}

func GetTasks() ([]models.Task, error) {
	tasks := []models.Task{}
	cursor, err := taskCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func GetTask(taskId primitive.ObjectID) (models.Task, error) {
	var task models.Task
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	err := taskCollection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return task, errors.New("task not found")
		}
		return task, err
	}
	return task, nil
}

func AddTask(task models.Task) (*mongo.InsertOneResult, error) {

	return taskCollection.InsertOne(context.Background(), task)

}

func UpdateTask(taskId primitive.ObjectID, updatedTask models.Task) (*mongo.UpdateResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	update := bson.D{primitive.E{Key: "$set", Value: updatedTask}}
	result, err := taskCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 {
		return nil, errors.New("task not found")
	}

	return result, nil
}

func DeleteTask(taskId primitive.ObjectID) (*mongo.DeleteResult, error) {
	filter := bson.D{primitive.E{Key: "_id", Value: taskId}}
	result, err := taskCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	if result.DeletedCount == 0 {
		return nil, errors.New("task not found")
	}

	return result, nil
}
