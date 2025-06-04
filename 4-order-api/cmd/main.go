package main

import (
	"fmt"
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/internal/auth"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()
}