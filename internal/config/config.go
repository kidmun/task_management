package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Client, error) {
	if err := os.Chdir(".."); err != nil {
        log.Fatalf("Error changing directory: %v", err)
    }

	err := godotenv.Load()
	fmt.Println("++++++++++++++++++")
	fmt.Println(os.Getwd())
	fmt.Println("-------------------")
	if err != nil {
		fmt.Println(err)
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
