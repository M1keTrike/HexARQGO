package application

import (
	"demo/src/users/domain/entities"
	repositories "demo/src/users/infraestructure/persistence"
)

type GetAllUsers struct {
	db repositories.UserRepository
}

func NewGetAllUsers(db repositories.UserRepository) *GetAllUsers {
	return &GetAllUsers{db: db}
}

func (ga *GetAllUsers) Execute() ([]entities.User, error) {
	res, err := ga.db.GetAll()
	if err != nil {
		return res, err
	}
	return res, nil
}
