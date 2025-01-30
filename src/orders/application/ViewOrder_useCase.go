package application

import (
	"demo/src/orders/domain/entities"
	repositories "demo/src/orders/infraestructure/persistence"
)

type GetAllOrders struct {
	db repositories.OrderRepository
}

func NewGetAllOrders(db repositories.OrderRepository) *GetAllOrders {
	return &GetAllOrders{db: db}
}

func (ga *GetAllOrders) Execute()  ([]entities.Order,error) {
	res ,err  := ga.db.GetAll()
	if err != nil  {
			
		return res,err
	}
	
	return res,nil
}