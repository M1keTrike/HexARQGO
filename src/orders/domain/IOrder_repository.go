package domain

import (
	"demo/src/orders/domain/entities"
)

type IOrder interface {
	Save(order *entities.Order) error
	GetAll() ([]entities.Order, error)
	DeleteById(id int) error
	EditById(id int, updatedOrder *entities.Order) error
}
