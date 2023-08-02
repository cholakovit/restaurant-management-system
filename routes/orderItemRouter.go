package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("orderItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.POST("/orderItems/:orderItem_id", controllers.UpdateOrderItem())

	incomingRoutes.POST("/orderItems-order/:order_id", controllers.GetOrderItemsByOrder())
}