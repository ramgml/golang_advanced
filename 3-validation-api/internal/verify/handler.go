package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/files"
	"3-validation-api/pkg/keygen"
	"3-validation-api/pkg/mail"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/response"
	"3-validation-api/pkg/vault"
	"encoding/json"
	"fmt"
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
	router.HandleFunc("POST /send", handler.SendEmail())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (h *EmailHandler) SendEmail() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Декодирование тела запроса
		var requestBody SendEmailRequest
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}
		fmt.Println("Запрос пришел")
		// Валидация запроса
		err = request.IsValid(requestBody)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}
		fileDb := vault.NewVault(files.NewJsonDb(h.Config.Urn))	
		account, err := fileDb.GetAccountByEmail(requestBody.Email)
		if err != nil {
			key := keygen.GetUserKey(requestBody.Email)
			account = *vault.NewAccount(requestBody.Email, key)
			fileDb.AddAccount(account)
		}
		err = mail.SendMail(h.Address, h.Email, h.Password, &account)
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}
		data := &SendEmailResponse{
			Message: "Check your mailbox",
		}
		response.Json(w, data, 203)
	}
}

func (h *EmailHandler) Verify() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.PathValue("hash")
		payload := &VerifyResponse{
			Verified: true,
			Message: "Success",
		}
		status := 200
		fileDb := vault.NewVault(files.NewJsonDb(h.Config.Urn))
		acc, err := fileDb.GetAccountByKey(key)
		if err != nil {
			payload = &VerifyResponse{
				Verified: false,
				Message: err.Error(),
			}
			status = 404
		} else {
			fileDb.DeleteAccount(acc.Email)
		}
		response.Json(w, payload, status)
	}
}
