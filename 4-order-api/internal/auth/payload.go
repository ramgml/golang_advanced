package auth

type AuthByPhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type AuthByPhoneResponse struct {
	Message string `json:"message"`
}

type VerifyCodeRequest struct {
	Phone string `json:"phone" validate:"required"`
	Code  string `json:"code" validate:"required"`
}

type VerifyCodeResponse struct {
	IsVerified bool   `json:"isVerified"`
	Message    string `json:"message"`
}
