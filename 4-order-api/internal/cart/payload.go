package cart

type AddToCartRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  uint `json:"quantity" validate:"required"`
}

type CartResponse struct {
	ID     uint `json:"id"`
	Items  []CartItem
}
