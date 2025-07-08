package fakecircle

import (
	"errors"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	errDuplicate = errors.New("duplicate")
	errNotFound  = errors.New("not found")
)

type Service struct {
	http.Handler
	tok string

	hit429 atomic.Bool
	hit500 atomic.Bool

	mu       sync.RWMutex
	orgs     map[uuid.UUID]*org
	projects map[uuid.UUID]*project
	contexts map[uuid.UUID]*context
}

func New(tok string) *Service {
	r := gin.New()
	c := &Service{
		tok:      tok,
		Handler:  r,
		orgs:     make(map[uuid.UUID]*org),
		projects: make(map[uuid.UUID]*project),
		contexts: make(map[uuid.UUID]*context),
	}

	r.Use(c.auth)

	r.GET("/api/test/hello", c.getHello)
	r.POST("/api/test/echo", c.postEcho)
	r.GET("/api/test/429", c.get429)
	r.GET("/api/test/500", c.get500)

	r.POST("/api/v2/organization", c.postOrganization)
	r.DELETE("/api/v2/organization/:org-id", c.deleteOrganization)
	r.POST("/api/v2/organization/:org-id/project", c.postProject)

	r.GET("/api/v2/project/:org-type/:org-name/:project-name", c.getProject)
	r.DELETE("/api/v2/project/:org-type/:org-name/:project-name", c.deleteProject)

	r.GET("/api/v2/context", c.getContextBySlug)
	r.POST("/api/v2/context", c.postContext)
	r.GET("/api/v2/context/:context-id", c.getContextByID)
	r.DELETE("/api/v2/context/:context-id", c.deleteContext)
	r.GET("/api/v2/context/:context-id/environment-variable", c.getContextEnv)
	r.PUT("/api/v2/context/:context-id/environment-variable/:env-var", c.putContextEnv)
	r.DELETE("/api/v2/context/:context-id/environment-variable/:env-var", c.deleteContextEnv)

	return c
}

func (s *Service) getHello(c *gin.Context) {
	msg(c, http.StatusOK, "Hello World!")
}

func (s *Service) postEcho(c *gin.Context) {
	var body map[string]any
	err := c.BindJSON(&body)
	if err != nil {
		msg(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, body)
}

func (s *Service) get429(c *gin.Context) {
	if !s.hit429.Swap(true) {
		c.Header("Retry-After", "1")
		msg(c, http.StatusTooManyRequests, "Too many requests.")
		return
	}

	msg(c, http.StatusOK, "Successfully retried.")
}

func (s *Service) get500(c *gin.Context) {
	if !s.hit500.Swap(true) {
		msg(c, http.StatusInternalServerError, "Internal server error.")
		return
	}

	msg(c, http.StatusOK, "Successfully retried.")
}
