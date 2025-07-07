package fakecircle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) auth(c *gin.Context) {
	circleToken := c.GetHeader("Circle-Token")
	switch circleToken {
	case "":
		msg(c, http.StatusUnauthorized, "You must log in first.")
		c.Abort()
		return
	case s.tok:
		c.Next()
	default:
		msg(c, http.StatusUnauthorized, "Invalid token provided.")
		c.Abort()
		return
	}
}
