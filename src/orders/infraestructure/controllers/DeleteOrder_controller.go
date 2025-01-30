package infrastructure

import (
	"demo/src/orders/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteOrderByIDRequest struct {
	Id int `json:"id" binding:"required"`
}

type DeleteOrderByIDController struct {
	do application.DeleteOrderByID
}

func NewDeleteOrderByIDController(do application.DeleteOrderByID) *DeleteOrderByIDController {
	return &DeleteOrderByIDController{do: do}
}

func (do_c *DeleteOrderByIDController) Execute(c *gin.Context) {
	var req DeleteOrderByIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	idToDelete := req.Id
	fmt.Println("Deleting Order ID:", idToDelete)

	err := do_c.do.Execute(idToDelete)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order", "id": idToDelete})
		return
	}

	c.JSON(http.StatusOK, gin.H{"successful": "Deleted order", "id": idToDelete})
}
