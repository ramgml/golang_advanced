package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	verify.NewEmailHandler(router, verify.EmailHandlerDeps{
		Config: conf,
	})
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("Server is running")
	server.ListenAndServe()
}