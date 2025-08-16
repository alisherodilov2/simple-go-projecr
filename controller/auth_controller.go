package controllers

import (
	"net/http"

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
		c.JSON(http.StatusNoContent, gin.H{"error": loginErr.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user": resource.UserMake(*user)})
}
