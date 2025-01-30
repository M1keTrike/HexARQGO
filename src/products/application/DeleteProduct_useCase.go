package application

import repositories "demo/src/products/infraestructure/persistence"

type DeleteProductByID struct {
	db repositories.ProductRepository
}

func NewDeleteProductByID(db repositories.ProductRepository) *DeleteProductByID {
	return &DeleteProductByID{db: db}
}

func (cp *DeleteProductByID) Execute(toDeleteProductId int) error {
	 err := cp.db.DeleteById(toDeleteProductId)
	 if err != nil {
		return err
	 } 
	 return nil
}