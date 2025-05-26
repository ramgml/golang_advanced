package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/response"
	"net/http"
)

type EmailHandler struct{
	*configs.Config
}

type EmailHandlerDeps struct {
	*configs.Config
}

func NewEmailHandler(router *http.ServeMux, deps EmailHandlerDeps) {
	handler := &EmailHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *EmailHandler) Send() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := &SendResponse{
			Message: "Check your mailbox",
		}
		response.Json(w, payload, 203)
	}
}

func (h *EmailHandler) Verify() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := &VerifyResponse{
			Verified: true,
			Message: "Success",
		}
		response.Json(w, payload, 200)
	}
}
