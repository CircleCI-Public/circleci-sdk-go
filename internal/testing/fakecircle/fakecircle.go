package fakecircle

import (
	"net/http"
	"sync/atomic"
)

type Service struct {
	http.Handler
	tok string

	hit429 atomic.Bool
	hit500 atomic.Bool
}

func New(tok string) *Service {
	c := &Service{
		tok: tok,
	}

	mux := http.NewServeMux()
	c.Handler = mux
	c.Handler = c.auth(c.Handler)
	c.Handler = loggingHandler(c.Handler)

	mux.HandleFunc("GET /api/test/hello", c.getHello)
	mux.HandleFunc("POST /api/test/echo", c.postEcho)
	mux.HandleFunc("GET /api/test/429", c.get429)
	mux.HandleFunc("GET /api/test/500", c.get500)
	// TODO: More routes here

	return c
}

func (s *Service) getHello(w http.ResponseWriter, _ *http.Request) {
	msg(w, http.StatusOK, "Hello World!")
}

func (s *Service) postEcho(w http.ResponseWriter, r *http.Request) {
	var body map[string]any
	err := decode(r.Body, &body)
	if err != nil {
		msg(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonBody(w, body)
}

func (s *Service) get429(w http.ResponseWriter, _ *http.Request) {
	if !s.hit429.Swap(true) {
		w.Header().Set("Retry-After", "1")
		msg(w, http.StatusTooManyRequests, "Too many requests.")
		return
	}

	msg(w, http.StatusOK, "Successfully retried.")
}

func (s *Service) get500(w http.ResponseWriter, _ *http.Request) {
	if !s.hit500.Swap(true) {
		msg(w, http.StatusInternalServerError, "Internal server error.")
		return
	}

	msg(w, http.StatusOK, "Successfully retried.")
}

// TODO: More routes here
