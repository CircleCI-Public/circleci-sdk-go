package fakecircle

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func msg(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}

func mapBadRequest(c *gin.Context, message string, err error) bool {
	if err == nil {
		return false
	}
	slog.Warn(err.Error())
	msg(c, http.StatusBadRequest, message)
	return true
}

type listResponse[T any] struct {
	NextPageToken *string `json:"next_page_token"`
	Items         []T     `json:"items"`
}

func newListResponse[T any](items []T) listResponse[T] {
	return listResponse[T]{Items: items}
}
