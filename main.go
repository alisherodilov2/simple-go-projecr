package main

import (
	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/alisherodilov2/go-first/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.Connect()
	database.DB.AutoMigrate(&models.Book{})

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
