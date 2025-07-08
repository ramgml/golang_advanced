package order

import (
	"purple/4-order-api/internal/product"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID   uint
	Products []product.Product `gorm:"many2many:order_products;"`
}
