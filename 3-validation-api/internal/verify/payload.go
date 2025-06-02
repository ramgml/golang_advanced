package verify

type SendEmailResponse struct {
	Message string `json:"message"`
}

type VerifyResponse struct {
	Verified bool `json:"verified"`
	Message string `json:"message"`
}

type SendEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}
