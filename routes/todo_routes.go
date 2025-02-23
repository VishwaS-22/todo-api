package routes

import (
	"todo-api/controllers"
	"todo-api/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine) {
	// Public Routes
	router.GET("/todos", controllers.GetTodos)
	router.POST("/todos", controllers.CreateTodo)

	// Protected Routes (Require Authentication)
	protected := router.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.PUT("/todos/:id", controllers.UpdateTodo)
		protected.DELETE("/todos/:id", controllers.DeleteTodo)
	}
}
