package domain

import (
	"demo/src/orders/domain/entities"
)

type IOrder interface {
	Save(product *entities.Order) error
	GetAll() ([]entities.Order, error)
	DeleteById(id int) error
	EditById(id int, updatedProduct *entities.Order) error
}
