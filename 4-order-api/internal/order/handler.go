package order

import (
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/internal/product"
	"purple/4-order-api/internal/user"
	"purple/4-order-api/pkg/middleware"
	"purple/4-order-api/pkg/request"
	"purple/4-order-api/pkg/response"
	"strconv"
)

type OrderHandler struct {
	OrderRepository   *OrderRepository
	UserRepository    *user.UserRepository
	ProductRepository *product.ProductRepository
	Config            *configs.Config
}

type OrderHandlerDeps struct {
	OrderRepository *OrderRepository
	UserRepository  *user.UserRepository
	ProductRepository *product.ProductRepository
	Config          *configs.Config
}

func NewOrderHandler(router *http.ServeMux, deps *OrderHandlerDeps) {
	handler := &OrderHandler{
		OrderRepository: deps.OrderRepository,
		UserRepository: deps.UserRepository,
		ProductRepository: deps.ProductRepository,
		Config:          deps.Config,
	}
	router.Handle("POST /order", middleware.IsAuthed(handler.CreateOrder(), deps.Config))
	router.Handle("GET /order/{id}", middleware.IsAuthed(handler.GetOrder(), deps.Config))
	router.Handle("GET /my-orders", middleware.IsAuthed(handler.GetAllOrders(), deps.Config))
}

func (h *OrderHandler) CreateOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		user, err := h.UserRepository.GetByPhone(phone)
		if err != nil {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		body, err := request.HandleBody[CreatOrderRequest](&w, r)
		if err != nil {
			return
		}
		products, err := h.ProductRepository.GetByIds(body.Products)
		if err != nil {
			http.Error(w, "can't get product", http.StatusBadRequest)
			return
		}
		order := &Order{
			UserID:   user.ID,
			Products: products,
		}
		err = h.OrderRepository.Create(order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, "A new order has been successfully created", http.StatusCreated)
	}
}

func (h *OrderHandler) GetOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orderID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		user, err := h.UserRepository.GetByPhone(phone)
		if err != nil {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		order, err := h.OrderRepository.GetOrderForUser(uint(orderID), user.ID)
		if err != nil {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}
		response.Json(w, order, http.StatusOK)
	}
}

func (h *OrderHandler) GetAllOrders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		phone, ok := r.Context().Value(middleware.ContextPhoneKey).(string)
		if !ok {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		user, err := h.UserRepository.GetByPhone(phone)
		if err != nil {
			http.Error(w, "can't get phone", http.StatusBadRequest)
			return
		}
		orders, err := h.OrderRepository.GetAllOrderForUser(user.ID)
		if err != nil {
			http.Error(w, "orders not found", http.StatusNotFound)
			return
		}
		response.Json(w, orders, http.StatusOK)
	}
}
