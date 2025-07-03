package fakecircle

import (
	"encoding/json"
	"net/http"
)

func msg(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	jsonBody(w, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func jsonBody(w http.ResponseWriter, v any) {
	_ = json.NewEncoder(w).Encode(v)
}
