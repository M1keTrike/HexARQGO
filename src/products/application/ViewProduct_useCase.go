package application

import (
	"demo/src/products/domain/entities"
	repositories "demo/src/products/infraestructure/persistence"
)

type GetAllProducts struct {
	db repositories.ProductRepository
}

func NewGetAllProducts(db repositories.ProductRepository) *GetAllProducts {
	return &GetAllProducts{db: db}
}

func (ga *GetAllProducts) Execute()  ([]entities.Product,error) {
	res ,err  := ga.db.GetAll()
	if err != nil  {
			
		return res,err
	}
	
	return res,nil
}