package verify

type SendResponse struct {
	Message string `json:"message"`
}

type VerifyResponse struct {
	Verified bool `json:"verified"`
	Message string `json:"message"`
}