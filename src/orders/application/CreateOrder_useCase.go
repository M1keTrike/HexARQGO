package application

import (
	"demo/src/orders/domain/entities"
	repositories "demo/src/orders/infraestructure/persistence"
)

type CreateOrder struct {
	db repositories.OrderRepository
}

func NewCreateOrder(db repositories.OrderRepository) *CreateOrder {
	return &CreateOrder{db: db}
}

func (co *CreateOrder) Execute(NewOrder *entities.Order)  error {
	err := co.db.Save(NewOrder)
	if err != nil  {
		return err
	}
	return nil
}