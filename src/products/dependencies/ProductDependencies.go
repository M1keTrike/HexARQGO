package dependencies

import (
	"database/sql"
	"demo/src/products/application"
	productInfra "demo/src/products/infraestructure/controllers"
	productRouters "demo/src/products/infraestructure/interfaces/http/routers"
	productRepositories "demo/src/products/infraestructure/persistence"

	"github.com/gin-gonic/gin"
)

type ProductsDependencies struct {
	DB *sql.DB
}

func NewProductsDependencies(db *sql.DB) *ProductsDependencies {
	return &ProductsDependencies{DB: db}
}

func (d *ProductsDependencies) Execute(r *gin.Engine) {
	productRepository := productRepositories.NewProductRepository(d.DB)

	createProductUseCase := application.NewCreateProduct(*productRepository)
	createProductController := productInfra.NewCreateProductController(*createProductUseCase)

	deleteProductUseCase := application.NewDeleteProductByID(*productRepository)
	deleteProductByIDController := productInfra.NewDeleteProductByIDController(*deleteProductUseCase)

	getAllProductUseCase := application.NewGetAllProducts(*productRepository)
	getAllProductController := productInfra.NewGetAllProductsController(*getAllProductUseCase)

	updateProductsByIDUseCase := application.NewEditProduct(*productRepository)
	updateProductsByIDController := productInfra.NewUpdateProductByIDController(*updateProductsByIDUseCase)

	productRouters.AttachProductRoutes(
		r,
		createProductController,
		deleteProductByIDController,
		getAllProductController,
		updateProductsByIDController,
	)
}
