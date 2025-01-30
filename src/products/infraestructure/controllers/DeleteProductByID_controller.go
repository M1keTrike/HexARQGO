package infraestructure

import (
	"demo/src/products/application"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteProductByIDRequest struct {
	Id int
}

type DeleteProductByIDController struct{
	dp application.DeleteProductByID
}

func NewDeleteProductByIDController(dp application.DeleteProductByID) *DeleteProductByIDController {
	return &DeleteProductByIDController{dp:dp}
}

func (dp_c_c *DeleteProductByIDController) Execute(c *gin.Context) {

	var req DeleteProductByIDRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	idToDelete := req.Id
	fmt.Println(idToDelete)

	err := dp_c_c.dp.Execute(idToDelete)
	if err != nil {
		
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product", "id": idToDelete})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sucessful": "Deleted product" ,"id": idToDelete})
	
}
