package order

type CreatOrderRequest struct {
	Products []uint `json:"products" validate:"required"`
}