package application

import (
	"demo/src/users/domain/entities"
	repositories "demo/src/users/infraestructure/persistence"
)

type EditUser struct {
	db repositories.UserRepository
}

func NewEditUser(db repositories.UserRepository) *EditUser {
	return &EditUser{db: db}
}

func (eu *EditUser) Execute(newData *entities.User) error {
	err := eu.db.EditById(int(newData.Id), newData)
	if err != nil {
		return err
	}
	return nil
}