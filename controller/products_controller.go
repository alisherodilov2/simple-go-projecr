package controllers

import (
	"net/http"

	"github.com/alisherodilov2/go-first/database"
	"github.com/alisherodilov2/go-first/models"
	"github.com/alisherodilov2/go-first/resource"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"data": resource.BookCollection(books),
	})
}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resource.BookMake(book),
	})
}

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(http.StatusCreated, gin.H{
		"data": resource.BookMake(book),
	})
}

func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var input models.Book
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = input.Title
	book.Author = input.Author
	database.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"data": resource.BookMake(book),
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	database.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
