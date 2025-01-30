package repositories

import (
	"database/sql"
	"demo/src/orders/domain/entities"
	"log"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (repo *OrderRepository) Save(order *entities.Order) error {
	query := "INSERT INTO orders (actor, product, quantity) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, order.Actor, order.Product, order.Quantity)
	if err != nil {
		log.Printf("[OrderRepository.Save] Error inserting order: %v", err)
		return err
	}
	log.Println("[OrderRepository.Save] Order inserted successfully")
	return nil
}

func (repo *OrderRepository) GetAll() ([]entities.Order, error) {
	query := "SELECT id, actor, product, quantity FROM orders"
	rows, err := repo.DB.Query(query)
	if err != nil {
		log.Printf("[OrderRepository.GetAll] Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []entities.Order
	for rows.Next() {
		var order entities.Order
		if err := rows.Scan(&order.Id, &order.Actor, &order.Product, &order.Quantity); err != nil {
			log.Printf("[OrderRepository.GetAll] Error scanning row: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	if rows.Err() != nil {
		log.Printf("[OrderRepository.GetAll] Error iterating over rows: %v", rows.Err())
		return nil, rows.Err()
	}
	log.Printf("[OrderRepository.GetAll] Successfully retrieved %d orders", len(orders))
	return orders, nil
}

func (repo *OrderRepository) DeleteById(id int) error {
	query := "DELETE FROM orders WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	if err != nil {
		log.Printf("[OrderRepository.DeleteById] Error deleting order with ID %d: %v", id, err)
		return err
	}
	log.Printf("[OrderRepository.DeleteById] Order with ID %d deleted successfully", id)
	return nil
}

func (repo *OrderRepository) EditById(id int, updatedOrder *entities.Order) error {
	query := "UPDATE orders SET actor = ?, product = ?, quantity = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, updatedOrder.Actor, updatedOrder.Product, updatedOrder.Quantity, id)
	if err != nil {
		log.Printf("[OrderRepository.EditById] Error updating order with ID %d: %v", id, err)
		return err
	}
	log.Printf("[OrderRepository.EditById] Order with ID %d updated successfully", id)
	return nil
}
