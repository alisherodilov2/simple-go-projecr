package routes

import (
	"github.com/alisherodilov2/go-first/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	bookRouter := engine.Group("/books")
	{
		bookRouter.GET("/", controllers.GetBooks)
		bookRouter.GET("/:id", controllers.GetBook)
		bookRouter.DELETE("/:id", controllers.DeleteBook)
		bookRouter.POST("/", controllers.CreateBook)
		bookRouter.POST("/:id", controllers.UpdateBook)
	}
}
