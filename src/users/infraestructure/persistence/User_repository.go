package repositories

import (
	"database/sql"
	"demo/src/users/domain/entities"
	"log"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) Save(user *entities.User) error {
	query := "INSERT INTO users (username, password, role) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, user.Username, user.Password, user.Role)
	if err != nil {
		log.Printf("[UserRepository.Save] Error inserting user: %v", err)
		return err
	}
	log.Println("[UserRepository.Save] User inserted successfully")
	return nil
}

func (repo *UserRepository) GetAll() ([]entities.User, error) {
	query := "SELECT id, username, password, role FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil {
		log.Printf("[UserRepository.GetAll] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Role); err != nil {
			log.Printf("[UserRepository.GetAll] Error scanning row: %v", err)
			return nil, err
		}
		users = append(users, user)
	}
	if rows.Err() != nil {
		log.Printf("[UserRepository.GetAll] Error iterating over rows: %v", rows.Err())
		return nil, rows.Err()
	}
	log.Printf("[UserRepository.GetAll] Successfully retrieved %d users", len(users))
	return users, nil
}

func (repo *UserRepository) DeleteById(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		log.Printf("[UserRepository.DeleteById] Error deleting user with ID %d: %v", id, err)
		return err
	}
	log.Printf("[UserRepository.DeleteById] User with ID %d deleted successfully", id)
	return nil
}

func (repo *UserRepository) EditById(id int, updatedUser *entities.User) error {
	query := "UPDATE users SET username = ?, password = ?, role = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, updatedUser.Username, updatedUser.Password, updatedUser.Role, id)
	if err != nil {
		log.Printf("[UserRepository.EditById] Error updating user with ID %d: %v", id, err)
		return err
	}
	log.Printf("[UserRepository.EditById] User with ID %d updated successfully", id)
	return nil
}
