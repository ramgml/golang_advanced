package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"purple/4-order-api/configs"
	"purple/4-order-api/internal/auth"
	"purple/4-order-api/internal/product"
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
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
		AuthService: auth.NewAuthService(auth.NewSessionRepository(dbConn)),
	})
	product.NewProductHandler(router, product.ProductHandlerDeps{
		ProductRepository: product.NewProductRepository(dbConn),
	})
	chain := middleware.Chain(
		middleware.Logging,
	)
	server := http.Server{
		Addr:    ":8081",
		Handler: chain(router),
	}
	log.WithFields(log.Fields{
		"status": "work",
		"event": "start",
	}).Info("Server is running")
	server.ListenAndServe()
}
