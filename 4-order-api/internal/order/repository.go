package order

import "purple/4-order-api/pkg/db"

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(db *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: db,
	}
}

func (repo *OrderRepository) Create(order *Order) error {
	return repo.Database.DB.Create(order).Error
}

func (repo *OrderRepository) GetOrderForUser(id uint, userID uint) (*Order, error) {
	var order Order
	result := repo.Database.DB.Preload("Products").First(&order, "id = ? AND user_id = ?",  id, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

func (repo *OrderRepository) GetAllOrderForUser(userID uint) (*[]Order, error) {
	var orders []Order
	result := repo.Database.DB.Preload("Products").Find(&orders, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}

