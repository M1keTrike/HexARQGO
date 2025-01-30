package dependencies

import (
	"database/sql"
	"demo/src/orders/application"
	orderInfra "demo/src/orders/infraestructure/controllers"
	orderRouters "demo/src/orders/infraestructure/interfaces/http/routers" // Asegúrate de que este import sea correcto
	orderRepositories "demo/src/orders/infraestructure/persistence"

	"github.com/gin-gonic/gin"
)

type OrdersDependencies struct {
	DB *sql.DB
}

func NewOrdersDependencies(db *sql.DB) *OrdersDependencies {
	return &OrdersDependencies{DB: db}
}

func (d *OrdersDependencies) Execute(r *gin.Engine) {
	orderRepository := orderRepositories.NewOrderRepository(d.DB)

	createOrderUseCase := application.NewCreateOrder(*orderRepository)
	createOrderController := orderInfra.NewCreateOrderController(*createOrderUseCase)

	deleteOrderUseCase := application.NewDeleteOrderByID(*orderRepository)
	deleteOrderByIDController := orderInfra.NewDeleteOrderByIDController(*deleteOrderUseCase)

	getAllOrdersUseCase := application.NewGetAllOrders(*orderRepository)
	getAllOrdersController := orderInfra.NewGetAllOrdersController(*getAllOrdersUseCase)

	updateOrderByIDUseCase := application.NewEditOrder(*orderRepository)
	updateOrderByIDController := orderInfra.NewUpdateOrderByIDController(*updateOrderByIDUseCase)

	orderRouters.AttachOrderRoutes(
		r,
		createOrderController,
		deleteOrderByIDController,
		getAllOrdersController,
		updateOrderByIDController,
	)
}
