package fakecircle

import (
	"net/http"
)

func (s *Service) auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		circleToken := r.Header.Get("Circle-Token")
		switch circleToken {
		case "":
			msg(w, http.StatusUnauthorized, "You must log in first.")
			return
		case s.tok:
			next.ServeHTTP(w, r)
		default:
			msg(w, http.StatusUnauthorized, "Invalid token provided.")
			return
		}
	})
}
