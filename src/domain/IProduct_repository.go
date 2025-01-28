package domain

import "demo/src/domain/entities"

type IProduct interface {
	Save(product *entities.Product) error
	GetAll() ([]entities.Product, error)
	DeleteById(id int) error
	EditById(id int, updatedProduct *entities.Product) error
}
