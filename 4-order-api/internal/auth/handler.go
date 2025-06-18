package auth

import (
	"net/http"
	"purple/4-order-api/configs"
	"purple/4-order-api/pkg/jwt"
	"purple/4-order-api/pkg/request"
	"purple/4-order-api/pkg/response"
	"purple/4-order-api/pkg/sms"
)

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

type AuthHandler struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
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
		session, err := ah.AuthService.Auth(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		sms.SendSms(session.Code)
		data := &AuthByPhoneResponse{
			SessionUid: session.Uid,
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
		session, err := ah.SessionRepository.GetByUid(body.SessionUid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		if session.Code != body.Code {
			sms.SendSms(session.Code)
			http.Error(w, "wrong code", http.StatusUnauthorized)
			return
		}
		token, err := jwt.NewJWT(ah.Config.Auth.Secret).Create(body.SessionUid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := &VerifyCodeResponse{
			Token: token,
		}
		response.Json(w, data, 200)
	}
}
