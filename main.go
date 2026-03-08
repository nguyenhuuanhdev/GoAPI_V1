package main

import (
	"go-api/database"
	"go-api/middleware"
	"go-api/models"
	"go-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(middleware.CORSMiddleware()) 

	database.ConnectDB()

	database.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run(":8080")
}