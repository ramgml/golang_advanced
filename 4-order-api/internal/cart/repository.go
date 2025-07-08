package cart

import (
	"purple/4-order-api/pkg/db"
)

type CartRepository struct {
	Database *db.Db
}

func NewCartRespository(db *db.Db) *CartRepository {
	return &CartRepository{
		Database: db,
	}
}

func (repo *CartRepository) Create(cart *Cart) (*Cart, error) {
	result := repo.Database.DB.Create(cart)
	if result.Error != nil {
		return nil, result.Error
	}
	return cart, nil
}

func (repo *CartRepository) GetByUser(userID uint) (*Cart, error) {
	var cart Cart
	result := repo.Database.DB.Preload("Items").Find(&cart, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cart, nil
}

func (repo *CartRepository) CreateItem(item *CartItem) (*CartItem, error) {
	result := repo.Database.DB.Create(item)
	if result.Error != nil {
		return nil, result.Error
	}
	return item, nil
}

func (repo *CartRepository) UpdateItem(item *CartItem) error {
	return repo.Database.DB.Save(item).Error
}

func (repo *CartRepository) FindItem(cartId uint, productId uint) (*CartItem, error) {
	var cartItem CartItem
	err := repo.Database.DB.Where("cart_id = ? AND product_id = ?", cartId, productId).First(&cartItem).Error
	if err != nil {
		return nil, err
	}
	return &cartItem, nil
}
