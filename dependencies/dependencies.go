package dependencies

import (
	"database/sql"
	"log"

	"demo/src/application"
	infraestructure "demo/src/infraestructure/controllers"
	"demo/src/infraestructure/database"
	"demo/src/infraestructure/interfaces/http/routers"
	"demo/src/infraestructure/repositories"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	DB *sql.DB 
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func (d *Dependencies) Execute() (*gin.Engine, *sql.DB) {
	database.LoadConfig()

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	d.DB = db

	productRepository := repositories.NewProductRepository(db)

	createProductUseCase := application.NewCreateProduct(*productRepository)
	createProductController := infraestructure.NewCreateProductController(*createProductUseCase)

	deleteProductUseCase := application.NewDeleteProductByID(*productRepository)
	deleteProductByIDController := infraestructure.NewDeleteProductByIDController(*deleteProductUseCase)

	getAllProductUseCase := application.NewGetAllProducts(*productRepository)
	getAllProductController := infraestructure.NewGetAllProductsController(*getAllProductUseCase)

	updateProductsByIDUseCase := application.NewEditProduct(*productRepository)
	updateProductsByIDController := infraestructure.NewUpdateProductByIDController(*updateProductsByIDUseCase)

	r := routers.CreateProductRoutes(
		createProductController,
		deleteProductByIDController,
		getAllProductController,
		*updateProductsByIDController,
	)

	return r, db
}
