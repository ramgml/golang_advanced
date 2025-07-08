package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"purple/4-order-api/configs"
	"purple/4-order-api/internal/auth"
	"purple/4-order-api/internal/cart"
	"purple/4-order-api/internal/order"
	"purple/4-order-api/internal/product"
	"purple/4-order-api/internal/user"
	"purple/4-order-api/pkg/db"
	"purple/4-order-api/pkg/middleware"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	dbConn := db.NewDb(conf)
	// repository
	sessionRepo := auth.NewSessionRepository(dbConn)
	userRepo := user.NewUserRepositry(dbConn)
	productRepo := product.NewProductRepository(dbConn)
	cartRepo := cart.NewCartRespository(dbConn)
	orderRepo := order.NewOrderRepository(dbConn)
	// service
	cartService := cart.NewCartService(&cart.CartServiceDeps{
		UserRepository: userRepo,
		CartRepository: cartRepo,
	})
	// handler
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: conf,
		AuthService: auth.NewAuthService(
			sessionRepo,
			userRepo,
		),
	})
	product.NewProductHandler(router, &product.ProductHandlerDeps{
		ProductRepository: productRepo,
		Config:            conf,
	})
	cart.NewCartHandler(router, &cart.CartHandlerDeps{
		CartService: cartService,
		Config:      conf,
	})
	order.NewOrderHandler(router, &order.OrderHandlerDeps{
		OrderRepository: orderRepo,
		UserRepository:  userRepo,
		ProductRepository: productRepo,
		Config:          conf,
	})
	// middleware
	chain := middleware.Chain(
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":8081",
		Handler: chain(router),
	}
	log.WithFields(log.Fields{
		"status": "work",
		"event":  "start",
	}).Info("Server is running")
	server.ListenAndServe()
}
