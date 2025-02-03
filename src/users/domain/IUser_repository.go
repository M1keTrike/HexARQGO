package domain

import (
	"demo/src/users/domain/entities"
)

type IUser interface {
	Save(user *entities.User) error
	GetAll() ([]entities.User, error)
	DeleteById(id int) error
	EditById(id int, updatedUser *entities.User) error
}
