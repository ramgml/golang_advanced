package cart

import (
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/pkg/middleware"
	"purple/4-order-api/pkg/request"
	"purple/4-order-api/pkg/response"
)

type CartHandler struct {
	CartService *CartService
	Config      *configs.Config
}

type CartHandlerDeps struct {
	CartService *CartService
	Config      *configs.Config
}

func NewCartHandler(router *http.ServeMux, deps *CartHandlerDeps) {
	handler := &CartHandler{
		Config:      deps.Config,
		CartService: deps.CartService,
	}
	router.Handle("GET /cart", middleware.IsAuthed(handler.GetCart(), deps.Config))
	router.Handle("POST /cart/item", middleware.IsAuthed(handler.AddCartItem(), deps.Config))
}

func (h *CartHandler) GetCart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		cart, err := h.CartService.GetCartByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}	
		cartResponse := &CartResponse{
			ID: cart.ID,
			Items: cart.Items,
		}
		response.Json(w, cartResponse, http.StatusOK)
	}
}

func (h *CartHandler) AddCartItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		body, err := request.HandleBody[AddToCartRequest](&w, r)
		if err != nil {
			return 
		}
		cart, err := h.CartService.GetCartByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = h.CartService.AddCartItem(cart, body.ProductID, body.Quantity)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cart, err = h.CartService.GetCartByPhone(phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cartResponse := &CartResponse{
			ID: cart.ID,
			Items: cart.Items,
		}
		response.Json(w, cartResponse, http.StatusOK)
	}
}
