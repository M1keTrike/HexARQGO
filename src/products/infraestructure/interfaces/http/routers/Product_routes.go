package routers

import (
	infrastructure "demo/src/products/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func AttachProductRoutes(
	r *gin.Engine,
	createProductController *infrastructure.CreateProductController,
	deleteProductByIDController *infrastructure.DeleteProductByIDController,
	getAllProductsController *infrastructure.GetAllProductController,
	updateProductByIdController *infrastructure.UpdateProductByIDController,
) {
	productsGroup := r.Group("/products")
	{
		productsGroup.POST("", createProductController.Execute)
		productsGroup.DELETE("", deleteProductByIDController.Execute)
		productsGroup.GET("", getAllProductsController.Execute)
		productsGroup.PUT("", updateProductByIdController.Execute)
	}
}
