package response

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, res any, statusCode int) {
	w.Header().Set("Content-Type", "applictaion/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(res)
}