package fakecircle

import (
	"errors"
	"fmt"
	"maps"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

//nolint:revive
const (
	TypeGitHub    = "github"
	TypeBitbucket = "bitbucket"
	TypeCircleCI  = "circleci"
)

func (s *Service) orgBySlug(slug string) *org {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for id, o := range s.orgs {
		if fmtOrgSlug(o.typ, id, o.name) == slug {
			return &org{
				id:       o.id,
				typ:      o.typ,
				name:     o.name,
				contexts: maps.Clone(o.contexts),
				projects: maps.Clone(o.projects),
			}
		}
	}
	return nil
}

type org struct {
	id       uuid.UUID
	typ      string
	name     string
	contexts map[uuid.UUID]*context
	projects map[uuid.UUID]*project
}

func (o *org) addProject(np NewProject) (*project, error) {
	for _, p := range o.projects {
		if p.Name == np.Name {
			return nil, errDuplicate
		}
	}

	p := &project{
		ID:   uuid.New(),
		Org:  o,
		Name: np.Name,
	}

	o.projects[p.ID] = p
	return p, nil
}

func (o *org) deleteProject(id uuid.UUID) {
	delete(o.projects, id)
}

type NewOrg struct {
	Type string
	Name string
}

type Org struct {
	ID   uuid.UUID
	Type string
	Name string
	Slug string
}

func fmtOrgSlug(typ string, id uuid.UUID, name string) string {
	second := name

	if typ == TypeCircleCI {
		second = base58.Encode(id[:])
	}

	return fmt.Sprintf("%s/%s", typ, second)
}

func fmtProjectSlugSuffix(typ string, id uuid.UUID, name string) string {
	if typ == TypeCircleCI {
		return base58.Encode(id[:])
	}

	return name
}

func (s *Service) AddOrg(newOrg NewOrg) (Org, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, o := range s.orgs {
		if o.name == newOrg.Name {
			return Org{}, errDuplicate
		}
	}

	o := &org{
		id:       uuid.New(),
		typ:      newOrg.Type,
		name:     newOrg.Name,
		contexts: make(map[uuid.UUID]*context),
		projects: make(map[uuid.UUID]*project),
	}
	s.orgs[o.id] = o
	return Org{
		ID:   o.id,
		Type: o.typ,
		Name: o.name,
		Slug: fmtOrgSlug(o.typ, o.id, o.name),
	}, nil
}

func (s *Service) deleteOrg(id uuid.UUID) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	o, ok := s.orgs[id]
	if !ok {
		return errNotFound
	}

	for contextID := range o.contexts {
		err := s.deleteEnvContext(contextID)
		if err != nil {
			return err
		}
	}

	delete(s.orgs, id)
	return nil
}

// handlers below here

func (s *Service) postOrganization(c *gin.Context) {
	type response struct {
		ID      uuid.UUID `json:"id"`
		Name    string    `json:"name"`
		VcsType string    `json:"vcs_type"`
		Slug    string    `json:"slug"`
	}

	var body struct {
		Type string `json:"vcs_type" binding:"required"`
		Name string `json:"name" binding:"required"`
	}
	err := c.BindJSON(&body)
	if mapBadRequest(c, "bad request", err) {
		return
	}

	o, err := s.AddOrg(NewOrg(body))
	switch {
	case errors.Is(err, errDuplicate):
		msg(c, http.StatusBadRequest, "duplicate org")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		ID:      o.ID,
		Name:    o.Name,
		VcsType: o.Type,
		Slug:    fmtOrgSlug(o.Type, o.ID, o.Name),
	})
}

func (s *Service) deleteOrganization(c *gin.Context) {
	orgID, err := uuid.Parse(c.Param("org-id"))
	if mapBadRequest(c, "bad org ID", err) {
		return
	}

	err = s.deleteOrg(orgID)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusNotFound, "org not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg(c, http.StatusOK, "ok")
}
