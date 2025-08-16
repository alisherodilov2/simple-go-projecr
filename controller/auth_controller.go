package controllers

import (
	"net/http"

	"github.com/alisherodilov2/go-first/models"
	"github.com/alisherodilov2/go-first/resource"
	"github.com/alisherodilov2/go-first/services"
	"github.com/gin-gonic/gin"

	"github.com/alisherodilov2/go-first/request"
)

var authService = services.AuthService{}

func Register(c *gin.Context) {
	request := new(request.RegisterRequest)

	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, registerErr := authService.Register(request.Username, request.Password)

	if registerErr != nil {
		c.JSON(http.StatusNoContent, gin.H{"error": registerErr.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": resource.UserMake(*user)})
}

func Login(c *gin.Context) {
	request := new(request.RegisterRequest)
	if err := c.ShouldBind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, user, loginErr := authService.Login(request.Username, request.Password)
	if loginErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": loginErr.Error()})
		return
	}

	c.SetCookie(
		"auth_token",
		tokenString,
		3600*24*3,
		"/",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"user":    resource.UserMake(*user), // don’t expose password
	})
}
func GetUser(c *gin.Context) {

	u, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, ok := u.(models.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resource.UserMake(user),
	})
}
