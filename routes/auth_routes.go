package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}
