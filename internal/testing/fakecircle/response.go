package fakecircle

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

func msg(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	jsonBody(w, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func decode(r io.ReadCloser, v any) (err error) {
	defer closer.ErrorHandler(r, &err)
	defer func() {
		_, _ = io.Copy(io.Discard, r)
	}()

	return json.NewDecoder(r).Decode(v)
}

func jsonBody(w http.ResponseWriter, v any) {
	_ = json.NewEncoder(w).Encode(v)
}
