package main

import (
	"log"
	"time"

	"task_management/internal/api/routes"
	"task_management/internal/config"

	"github.com/gin-gonic/gin"
)
func main() {
	client, err := config.InitDB()
	
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	router := gin.Default()
	routes.SetupTaskRoutes(router ,client.Database("task_management_db"), time.Second * 10)
	routes.SetupUserRoutes(router ,client.Database("task_management_db"), time.Second * 10)
	router.Run(":8080")

}
