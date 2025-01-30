package infrastructure

import (
	"demo/src/orders/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllOrdersController struct {
	ga application.GetAllOrders
}

func NewGetAllOrdersController(ga application.GetAllOrders) *GetAllOrdersController {
	return &GetAllOrdersController{ga: ga}
}

func (ga_c *GetAllOrdersController) Execute(c *gin.Context) {
	res, err := ga_c.ga.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve all orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Orders": res})
}
