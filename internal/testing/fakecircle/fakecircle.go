package fakecircle

import (
	"net/http"
)

type Service struct {
	http.Handler
	tok string
}

func New(tok string) *Service {
	c := &Service{
		tok: tok,
	}

	mux := http.NewServeMux()
	c.Handler = c.auth(mux)

	mux.HandleFunc("GET /api/test", c.getTest)
	// TODO: More routes here

	return c
}

func (s *Service) getTest(w http.ResponseWriter, _ *http.Request) {
	msg(w, http.StatusOK, "Hello World!")
}

// TODO: More routes here
