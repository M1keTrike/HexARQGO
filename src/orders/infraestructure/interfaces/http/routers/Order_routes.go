package routers

import (
	infrastructure "demo/src/orders/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func AttachOrderRoutes(
	r *gin.Engine,
	createOrderController *infrastructure.CreateOrderController,
	deleteOrderByIDController *infrastructure.DeleteOrderByIDController,
	getAllOrdersController *infrastructure.GetAllOrdersController,
	updateOrderByIdController *infrastructure.UpdateOrderByIDController,
) {
	ordersGroup := r.Group("/orders")
	{
		ordersGroup.POST("", createOrderController.Execute)
		ordersGroup.DELETE("", deleteOrderByIDController.Execute)
		ordersGroup.GET("", getAllOrdersController.Execute)
		ordersGroup.PUT("", updateOrderByIdController.Execute)
	}
}
