package infrastructure

import (
	"demo/src/orders/application"
	"demo/src/orders/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Actor    int32 `json:"actor" binding:"required"`   
	Product  int32 `json:"product" binding:"required"`
	Quantity int32 `json:"quantity" binding:"required"`
}

type CreateOrderController struct {
	cp application.CreateOrder
}

func NewCreateOrderController(cp application.CreateOrder) *CreateOrderController {
	return &CreateOrderController{cp: cp}
}

func (cp_c *CreateOrderController) Execute(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	order := entities.Order{
		Actor:    req.Actor, 
		Product:  req.Product,
		Quantity: req.Quantity,
	}

	err := cp_c.cp.Execute(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"actor":    order.Actor,
		"product":  order.Product,
		"quantity": order.Quantity,
	})
}
