package product

import (
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/pkg/middleware"
	"purple/4-order-api/pkg/request"
	"purple/4-order-api/pkg/response"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
	Config *configs.Config
}

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps *ProductHandlerDeps) {
	handler := &ProductHandler{
		ProductRepository: deps.ProductRepository,
	}
	router.HandleFunc("GET /products/{id}", handler.GetProduct())
	router.Handle("POST /products", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("PATCH /products/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.Handle("DELETE /products/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))
}

func (ph *ProductHandler) GetProduct() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := ph.ProductRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}

func (ph *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductCreateRequest](&w, r)
		if err != nil {
			return
		}
		product := NewProduct(
			body.Name,
			body.Description,
			body.Images,
		)
		product, err = ph.ProductRepository.Create(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, product, http.StatusCreated)
	}
}

func (ph *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		product, err := ph.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
			Images:      body.Images,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		response.Json(w, product, http.StatusOK)
	}
}

func (ph *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = ph.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		response.Json(w, nil, http.StatusNoContent)
	}
}
