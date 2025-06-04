package auth

import (
	"fmt"
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/pkg/request"
	"purple/4-order-api/pkg/response"
)

type AuthHandlerDeps struct{
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth", handler.AuthByPhone())
	router.HandleFunc("POST /auth/verify", handler.VerifyCode())
}

func (ah *AuthHandler) AuthByPhone() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[AuthByPhoneRequest](&w, r)
		if err != nil {
			return
		}
		// Отправка смс
		fmt.Println(*body)
		data := &AuthByPhoneResponse{
			Message: "Wait message to your phone number",
		}
		response.Json(w, data, 200)
	}
}

func (ah *AuthHandler) VerifyCode() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[VerifyCodeRequest](&w, r)
		if err != nil {
			return
		}
		// Проверка кода
		fmt.Println(*body)
		data := &VerifyCodeResponse{
			IsVerified: true,
			Message: "Phone number was verified",
		}
		response.Json(w, data, 200)
	}
}

