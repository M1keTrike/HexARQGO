package application

import (
	"demo/src/users/domain/entities"
	repositories "demo/src/users/infraestructure/persistence"
)

type CreateUser struct {
	db repositories.UserRepository
}

func NewCreateUser(db repositories.UserRepository) *CreateUser {
	return &CreateUser{db: db}
}

func (cu *CreateUser) Execute(newUser *entities.User)  error {
	err := cu.db.Save(newUser)
	if err != nil  {
		return err
	}
	return nil
}