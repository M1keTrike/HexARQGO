package application

import repositories "demo/src/orders/infraestructure/persistence"

type DeleteOrderByID struct {
	db repositories.OrderRepository
}

func NewDeleteOrderByID(db repositories.OrderRepository) *DeleteOrderByID {
	return &DeleteOrderByID{db: db}
}

func (do *DeleteOrderByID) Execute(toDeleteOrderId int) error {
	 err := do.db.DeleteById(toDeleteOrderId)
	 if err != nil {
		return err
	 } 
	 return nil
}