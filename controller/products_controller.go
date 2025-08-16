package controllers

import (
	"net/http"

	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/alisherodilov2/go-first/resource"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Products
	database.DB.Preload("Comments").Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"data": resource.ProductCollection(products),
	})
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Products

	if err := database.DB.Preload("Comments").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"data": resource.ProductMake(product),
	})
}
