package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	handler := &RandomHandler{}
	router.Handle("/dice", handler)
	fmt.Println("Server is running")
	http.ListenAndServe(":8081", router)
}

type RandomHandler struct {}

func (rh *RandomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	num := fmt.Sprint(rand.Intn(6) + 1)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(num))
}
