package routes

import (
	"github.com/alisherodilov2/go-first/controller"
	"github.com/alisherodilov2/go-first/middleware"
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

	productsRouter := engine.Group("/products")
	productsRouter.Use(middleware.AuthMiddleware())
	{
		productsRouter.GET("/", controllers.GetProducts)
		productsRouter.GET("/:id", controllers.GetProduct)
	}

	authRouter := engine.Group("/auth")
	{
		authRouter.POST("/register", controllers.Register)
		authRouter.POST("/login", controllers.Login)
	}
}
