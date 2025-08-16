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
