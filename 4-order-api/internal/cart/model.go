package cart

import (
	"purple/4-order-api/internal/product"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
}

type CartItem struct {
	gorm.Model
	Cart Cart
	Item product.Product
	Count int
}

