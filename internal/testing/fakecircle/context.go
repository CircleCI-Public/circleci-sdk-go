package fakecircle

import (
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type context struct {
	ID        uuid.UUID
	Name      string
	Org       *org
	EnvVars   []EnvVar
	CreatedAt time.Time
}

func (c *context) addEnv(ev NewEnvVar) (EnvVar, error) {
	if slices.ContainsFunc(c.EnvVars, func(e EnvVar) bool {
		return e.Variable == ev.Variable
	}) {
		return EnvVar{}, errDuplicate
	}

	now := time.Now()
	e := EnvVar{
		Variable:  ev.Variable,
		Value:     ev.Value,
		UpdatedAt: now,
		CreatedAt: now,
	}
	c.EnvVars = append(c.EnvVars, e)
	return e, nil
}

func (c *context) deleteEnv(ev string) {
	c.EnvVars = slices.DeleteFunc(c.EnvVars, func(e EnvVar) bool {
		return e.Variable == ev
	})
}

type NewContext struct {
	OrgID uuid.UUID
	Name  string
}

type Context struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}

func (s *Service) AddContext(c NewContext) (Context, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	o, ok := s.orgs[c.OrgID]
	if !ok {
		return Context{}, errNotFound
	}

	orgCtx := &context{
		Name:      c.Name,
		Org:       o,
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}

	o.contexts[orgCtx.ID] = orgCtx
	s.contexts[orgCtx.ID] = orgCtx
	return Context{
		ID:        orgCtx.ID,
		Name:      orgCtx.Name,
		CreatedAt: orgCtx.CreatedAt,
	}, nil
}

func (s *Service) getContext(id uuid.UUID) (context, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	envCtx, ok := s.contexts[id]
	if !ok {
		return context{}, false
	}

	return context{
		ID:        envCtx.ID,
		Name:      envCtx.Name,
		EnvVars:   slices.Clone(envCtx.EnvVars),
		CreatedAt: envCtx.CreatedAt,
	}, true
}

func (s *Service) deleteEnvContext(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	foundCtx, ok := s.contexts[id]
	if !ok {
		return errNotFound
	}

	_, ok = foundCtx.Org.contexts[id]
	if !ok {
		return errNotFound
	}

	delete(foundCtx.Org.contexts, id)
	delete(s.contexts, id)

	return nil
}

type NewEnvVar struct {
	Variable string
	Value    string
}

type EnvVar struct {
	Variable  string
	Value     string
	UpdatedAt time.Time
	CreatedAt time.Time
}

func (s *Service) AddContextEnv(contextID uuid.UUID, ev NewEnvVar) (EnvVar, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	envCtx, ok := s.contexts[contextID]
	if !ok {
		return EnvVar{}, errNotFound
	}

	return envCtx.addEnv(ev)
}

func (s *Service) deleteContextEnvVar(id uuid.UUID, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	envCtx, ok := s.contexts[id]
	if !ok {
		return errNotFound
	}

	envCtx.deleteEnv(name)
	return nil
}

// Handlers below here

func (s *Service) postContext(c *gin.Context) {
	type response struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	type owner struct {
		ID   uuid.UUID `json:"id"`
		Type string    `json:"type"`
	}
	var body struct {
		Name  string `json:"name"`
		Owner owner  `json:"owner"`
	}
	err := c.BindJSON(&body)
	if mapBadRequest(c, "bad request", err) {
		return
	}

	envCtx, err := s.AddContext(NewContext{
		OrgID: body.Owner.ID,
		Name:  body.Name,
	})
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "context not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response(envCtx))
}

func (s *Service) deleteContext(c *gin.Context) {
	id, err := uuid.Parse(c.Param("context-id"))
	if mapBadRequest(c, "bad context ID", err) {
		return
	}

	err = s.deleteEnvContext(id)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "context not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg(c, http.StatusOK, "ok")
}

func (s *Service) getContextByID(c *gin.Context) {
	type response struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	id, err := uuid.Parse(c.Param("context-id"))
	if mapBadRequest(c, "bad context ID", err) {
		return
	}

	envCtx, ok := s.getContext(id)
	if !ok {
		msg(c, http.StatusNotFound, "context not found")
		return
	}

	c.JSON(http.StatusOK, response{
		ID:        envCtx.ID,
		Name:      envCtx.Name,
		CreatedAt: envCtx.CreatedAt,
	})
}

func (s *Service) getContextBySlug(c *gin.Context) {
	type response struct {
		ID        uuid.UUID `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	ownerSlug := c.Query("owner-slug")
	if ownerSlug == "" {
		msg(c, http.StatusBadRequest, "missing slug")
		return
	}

	o := s.orgBySlug(ownerSlug)
	if o == nil {
		msg(c, http.StatusNotFound, "not found")
		return
	}

	res := make([]response, 0, len(o.contexts))
	for _, orgCtx := range o.contexts {
		res = append(res, response{
			ID:        orgCtx.ID,
			Name:      orgCtx.Name,
			CreatedAt: orgCtx.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, newListResponse(res))
}

func (s *Service) getContextEnv(c *gin.Context) {
	type response struct {
		ContextID uuid.UUID `json:"context_id"`
		Variable  string    `json:"variable"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	id, err := uuid.Parse(c.Param("context-id"))
	if mapBadRequest(c, "bad context ID", err) {
		return
	}

	envCtx, ok := s.getContext(id)
	if !ok {
		msg(c, http.StatusNotFound, "context not found")
		return
	}

	res := make([]response, 0, len(envCtx.EnvVars))
	for _, ev := range envCtx.EnvVars {
		res = append(res, response{
			Variable:  ev.Variable,
			UpdatedAt: ev.UpdatedAt,
			CreatedAt: ev.CreatedAt,
			ContextID: id,
		})
	}
	c.JSON(http.StatusOK, newListResponse(res))
}

func (s *Service) putContextEnv(c *gin.Context) {
	type response struct {
		ContextID uuid.UUID `json:"context_id"`
		Variable  string    `json:"variable"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	}

	contextID, err := uuid.Parse(c.Param("context-id"))
	if mapBadRequest(c, "bad request", err) {
		return
	}

	envVarName := c.Param("env-var")
	if envVarName == "" {
		msg(c, http.StatusBadRequest, "bad request")
		return
	}

	var body struct {
		Value string `json:"value"`
	}

	err = c.BindJSON(&body)
	if mapBadRequest(c, "bad request", err) {
		return
	}

	ev, err := s.AddContextEnv(contextID, NewEnvVar{
		Variable: envVarName,
		Value:    body.Value,
	})
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "context not found")
		return
	case errors.Is(err, errDuplicate):
		msg(c, http.StatusBadRequest, "env var already exists")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		ContextID: contextID,
		Variable:  ev.Variable,
		UpdatedAt: ev.UpdatedAt,
		CreatedAt: ev.CreatedAt,
	})
}

func (s *Service) deleteContextEnv(c *gin.Context) {
	contextID, err := uuid.Parse(c.Param("context-id"))
	if mapBadRequest(c, "bad request", err) {
		return
	}

	envVarName := c.Param("env-var")
	if envVarName == "" {
		msg(c, http.StatusBadRequest, "bad request")
		return
	}

	err = s.deleteContextEnvVar(contextID, envVarName)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "context not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg(c, http.StatusOK, "ok")
}
