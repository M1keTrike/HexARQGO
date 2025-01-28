package application

import (
	"demo/src/domain/entities"
	"demo/src/infraestructure/repositories"
)

type CreateProduct struct {
	db repositories.ProductRepository
}

func NewCreateProduct(db repositories.ProductRepository) *CreateProduct {
	return &CreateProduct{db: db}
}

func (cp *CreateProduct) Execute(newProduct *entities.Product)  error {
	err := cp.db.Save(newProduct)
	if err != nil  {
		return err
	}
	return nil
}