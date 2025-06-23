package auth

type AuthByPhoneRequest struct {
	Phone string `json:"phone" validate:"required"`
}

type AuthByPhoneResponse struct {
	SessionUid string `json:"sessionUid"`
}

type VerifyCodeRequest struct {
	SessionUid string `json:"sessionUid" validate:"required"`
	Code       string `json:"code" validate:"required"`
}

type VerifyCodeResponse struct {
	Token string `json:"token"`
}
