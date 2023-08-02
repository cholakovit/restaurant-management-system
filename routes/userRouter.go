package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("users/:user_id", controllers.GetUser())
	incomingRoutes.POST("/users", controllers.CreateUser())
	incomingRoutes.POST("/users/:user_id", controllers.UpdateUser())
}