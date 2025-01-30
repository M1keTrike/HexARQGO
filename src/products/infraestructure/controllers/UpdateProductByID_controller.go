package infraestructure

import (
	"demo/src/products/application"
	"demo/src/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateProductByIDController struct {
	up application.EditProduct
}

type UpdateProductRequest struct {
	Id  int32  `json:"id" binding:"required"` 
	Name  string  `json:"name" binding:"required"` 
	Price float32 `json:"price" binding:"required"` 
}

func NewUpdateProductByIDController(up application.EditProduct) *UpdateProductByIDController {
	return &UpdateProductByIDController{up:up}
}

func (up_c *UpdateProductByIDController) Execute(c *gin.Context) {
	var req UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updateProduct := entities.Product {
		Id: req.Id,
		Name: req.Name,
		Price: req.Price,
	}

	err := up_c.up.Execute(&updateProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product", "id": updateProduct.Id})
		return
	}

	
	c.JSON(http.StatusOK, gin.H{"succesful": "Updated product", "id": req.Id})
}
