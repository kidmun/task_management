package main

import (
	"log"
	"task_management/config"
	"task_management/routes"
	"task_management/services"
	"github.com/gin-gonic/gin"
)

func main() {
	client, err := config.InitDB()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	services.InitTaskCollection(client)
	router := gin.Default()
	routes.SetupTaskRoutes(router)
	router.Run(":8080")

}
