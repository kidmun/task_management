package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Client, error){
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
		return nil, err
    }
    mongoUri := GetEnv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, err
	}
return client, nil
}