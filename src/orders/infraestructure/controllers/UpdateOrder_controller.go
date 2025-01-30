package infrastructure

import (
	"demo/src/orders/application"
	"demo/src/orders/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateOrderByIDController struct {
	up application.EditOrder
}

type UpdateOrderRequest struct {
	Id       int32  `json:"id" binding:"required"`
	Actor    string `json:"actor" binding:"required"`
	Product  int32  `json:"product" binding:"required"`
	Quantity int32  `json:"quantity" binding:"required"`
}

func NewUpdateOrderByIDController(up application.EditOrder) *UpdateOrderByIDController {
	return &UpdateOrderByIDController{up: up}
}

func (up_c *UpdateOrderByIDController) Execute(c *gin.Context) {
	var req UpdateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updateOrder := entities.Order{
		Id:       req.Id,
		Actor:    req.Actor,
		Product:  req.Product,
		Quantity: req.Quantity,
	}

	err := up_c.up.Execute(&updateOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order", "id": updateOrder.Id})
		return
	}

	c.JSON(http.StatusOK, gin.H{"successful": "Updated order", "id": req.Id})
}
