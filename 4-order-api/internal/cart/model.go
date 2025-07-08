package cart

import (
	"purple/4-order-api/internal/user"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint
	User   user.User
	Items  []CartItem `json:"items"`
}

func NewCart(user *user.User) *Cart {
	return &Cart{
		UserID: user.ID,
		User:   *user,
	}
}

type CartItem struct {
	gorm.Model
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  uint `json:"quantity"`
}
