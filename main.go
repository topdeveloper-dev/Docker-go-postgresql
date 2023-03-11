package main

import (
	"fmt"
	"log"
	"test/controllers"
	"test/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin, access-control-allow-headers"},
	}))

	publicRoutes := router.Group("/api")

	publicRoutes.GET("/getuserlist", controllers.GetListUser)
	publicRoutes.POST("/adduser", controllers.AddUser)

	fmt.Println("Server running on port 8000")
	router.Run(":8000")
}
