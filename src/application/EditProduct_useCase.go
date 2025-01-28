package application

import (
	"demo/src/domain/entities"
	"demo/src/infraestructure/repositories"
)

type EditProduct struct {
	db repositories.ProductRepository
}

func NewEditProduct(db repositories.ProductRepository) *EditProduct {
	return &EditProduct{db: db}
}

func (cp *EditProduct) Execute(newData *entities.Product)  error {
	err := cp.db.EditById(int(newData.Id),newData)
	if err != nil  {
		return err
	}
	return nil
}