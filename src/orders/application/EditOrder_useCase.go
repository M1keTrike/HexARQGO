package application

import (
	"demo/src/orders/domain/entities"
	repositories "demo/src/orders/infraestructure/persistence"
)

type EditOrder struct {
	db repositories.OrderRepository
}

func NewEditOrder(db repositories.OrderRepository) *EditOrder {
	return &EditOrder{db: db}
}

func (eo *EditOrder) Execute(newData *entities.Order)  error {
	err := eo.db.EditById(int(newData.Id), newData)
	if err != nil  {
		return err
	}
	return nil
}