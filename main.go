package main

import (
	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/alisherodilov2/go-first/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		return
	}

	database.Connect()
	err := database.DB.AutoMigrate(&models.Book{})
	products := database.DB.AutoMigrate(&models.Products{})
	commment := database.DB.AutoMigrate(&models.Comments{})

	if products != nil {
		return
	}

	if commment != nil {
		return
	}

	if err != nil {
		return
	}

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":8080")
}
