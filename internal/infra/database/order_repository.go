package database

import (
	"database/sql"

	"github.com/GuilhermeVgl/go/internal/entities"
)

type OrderRepository struct {
	Db *sql.DB
}

// GetTotal implements entities.OrderRepositoryInterface.
func (*OrderRepository) GetTotal() (int, error) {
	panic("unimplemented")
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entities.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotalTransactions(order *entities.Order) (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
