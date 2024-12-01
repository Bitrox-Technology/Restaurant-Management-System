package routes

import (
	"golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("order-items/", controllers.GetOrderItems())
	incomingRoutes.GET("/order-item/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/order-item-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/order-item", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order-item", controllers.UpdateOrder())
}
