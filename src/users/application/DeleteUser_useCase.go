package application

import repositories "demo/src/users/infraestructure/persistence"

type DeleteUserByID struct {
	db repositories.UserRepository
}

func NewDeleteUserByID(db repositories.UserRepository) *DeleteUserByID {
	return &DeleteUserByID{db: db}
}

func (du *DeleteUserByID) Execute(toDeleteUserId int) error {
	err := du.db.DeleteById(toDeleteUserId)
	if err != nil {
		return err
	}
	return nil
}
