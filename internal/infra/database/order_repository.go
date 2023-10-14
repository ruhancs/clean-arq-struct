package repository

import (
	"clean-arq-struct/internal/domain/entity"
	"database/sql"
)


type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(database *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: database,
	}
}

func(repo *OrderRepository) Create(order *entity.Order) error{
	stmt, err := repo.DB.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

