package routers

import (
	infraestructure "demo/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)



func CreateProductRoutes(createProductController *infraestructure.CreateProductController, deleteProductByIDController *infraestructure.DeleteProductByIDController, getAllProductsController *infraestructure.GetAllProductController, updateProductByIdController infraestructure.UpdateProductByIDController) *gin.Engine {
	r := gin.Default()
	

	r.POST("/products", createProductController.Execute)
	r.DELETE("/products", deleteProductByIDController.Execute)
	r.GET("/products", getAllProductsController.Execute)
	r.PUT("/products", updateProductByIdController.Execute)

	return r
}

