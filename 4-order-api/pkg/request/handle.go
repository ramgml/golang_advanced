package request

import (
	"net/http"
	"purple/4-order-api/pkg/response"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	var errorResponse ErrorResponse
	body, err := Decode[T](r.Body)
	if err != nil {
		errorResponse.Error = err.Error()
		response.Json(*w, &errorResponse, 402)
		return nil, err
	}
	err = IsValid(body)
	if err != nil {
		errorResponse.Error = err.Error()
		response.Json(*w, &errorResponse, 402)
		return nil, err
	}
	return &body, nil
}
