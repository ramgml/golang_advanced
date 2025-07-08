package cart

import (
	"purple/4-order-api/internal/user"

	"gorm.io/gorm"
)

type CartService struct {
	CartRepository *CartRepository
	UserRepository *user.UserRepository
}

type CartServiceDeps struct {
	CartRepository *CartRepository
	UserRepository *user.UserRepository
}

func NewCartService(deps *CartServiceDeps) *CartService {
	return &CartService{
		CartRepository: deps.CartRepository,
		UserRepository: deps.UserRepository,
	}
}

func (service *CartService) GetCartByPhone(phone string) (*Cart, error) {
	user, err := service.UserRepository.GetByPhone(phone)
	if err != nil {
		return nil, err
	}
	cart, err := service.CartRepository.GetByUser(user.ID)
	if err != nil {
		cart, err = service.CartRepository.Create(NewCart(user))
	}
	return cart, err
}

func (service *CartService) AddCartItem(cart *Cart, productID uint, quantity uint) error {
	cartItem, err := service.CartRepository.FindItem(cart.ID, productID)
	if cartItem != nil {
		cartItem.Quantity = quantity
		return service.CartRepository.UpdateItem(cartItem)
	} else if err == gorm.ErrRecordNotFound {
		newItem := CartItem{
			CartID: cart.ID,
			ProductID: productID,
			Quantity: quantity,
		}
		_, err := service.CartRepository.CreateItem(&newItem)	
		return err
	}
	return err
}