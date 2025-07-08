package fakecircle

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type project struct {
	Org  *org
	ID   uuid.UUID
	Name string
}

func (p *project) ToProject() Project {
	o := p.Org
	orgSlug := fmtOrgSlug(o.typ, o.id, o.name)
	p2 := Project{
		ID:   p.ID,
		Name: p.Name,
		Slug: orgSlug + "/" + fmtProjectSlugSuffix(o.typ, p.ID, p.Name),
		Org: Org{
			ID:   o.id,
			Type: o.typ,
			Name: o.name,
			Slug: orgSlug,
		},
	}
	return p2
}

type NewProject struct {
	OrgID uuid.UUID
	Name  string
}

type Project struct {
	ID   uuid.UUID
	Name string
	Slug string

	Org Org
}

func (s *Service) AddProject(np NewProject) (Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	o, ok := s.orgs[np.OrgID]
	if !ok {
		return Project{}, errNotFound
	}

	p, err := o.addProject(np)
	if err != nil {
		return Project{}, err
	}

	s.projects[p.ID] = p
	return p.ToProject(), nil
}

func (s *Service) Project(id uuid.UUID) (Project, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.projects[id]
	if !ok {
		return Project{}, errNotFound
	}

	return p.ToProject(), nil
}

func (s *Service) projectBySlug(orgType, orgName, projectName string) (Project, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	o := s.orgBySlug(fmt.Sprintf("%s/%s", orgType, orgName))
	if o == nil {
		return Project{}, errNotFound
	}

	for _, p := range o.projects {
		if fmtProjectSlugSuffix(o.typ, p.ID, p.Name) == projectName {
			return p.ToProject(), nil
		}
	}

	return Project{}, errNotFound
}

func (s *Service) deleteProjectBySlug(orgType, orgName, projectName string) error {
	p, err := s.projectBySlug(orgType, orgName, projectName)
	if err != nil {
		return nil
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	o := s.orgs[p.Org.ID]
	o.deleteProject(p.ID)
	delete(s.projects, p.ID)
	return nil
}

// handlers below here

func (s *Service) postProject(c *gin.Context) {
	type response struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Slug string    `json:"slug"`

		OrganizationName string    `json:"organization_name"`
		OrganizationSlug string    `json:"organization_slug"`
		OrganizationID   uuid.UUID `json:"organization_id"`

		VcsInfo VcsInfo `json:"vcs_info"`
	}

	orgID, err := uuid.Parse(c.Param("org-id"))
	if mapBadRequest(c, "bad org ID", err) {
		return
	}

	var body struct {
		Name string `json:"name" binding:"required"`
	}
	err = c.BindJSON(&body)
	if mapBadRequest(c, "bad request", err) {
		return
	}

	prj, err := s.AddProject(NewProject{
		OrgID: orgID,
		Name:  body.Name,
	})
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusNotFound, "org not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		ID:   prj.ID,
		Name: prj.Name,
		Slug: prj.Slug,

		OrganizationName: prj.Org.Name,
		OrganizationSlug: prj.Org.Slug,
		OrganizationID:   prj.Org.ID,

		VcsInfo: VcsInfo{
			VcsURL:        "git://github.com/dummy-value",
			Provider:      prj.Org.Type,
			DefaultBranch: "main",
		},
	})
}

type VcsInfo struct {
	VcsURL        string `json:"vcs_url"`
	Provider      string `json:"provider"`
	DefaultBranch string `json:"default_branch"`
}

func (s *Service) getProject(c *gin.Context) {
	type response struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
		Slug string    `json:"slug"`

		OrganizationName string    `json:"organization_name"`
		OrganizationSlug string    `json:"organization_slug"`
		OrganizationID   uuid.UUID `json:"organization_id"`

		VcsInfo VcsInfo `json:"vcs_info"`
	}

	orgType := c.Param("org-type")
	switch orgType {
	case TypeGitHub, TypeBitbucket, TypeCircleCI:
		break
	default:
		msg(c, http.StatusBadRequest, "invalid org type")
		return
	}

	orgName := c.Param("org-name")
	projectName := c.Param("project-name")
	prj, err := s.projectBySlug(orgType, orgName, projectName)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusNotFound, "project not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		ID:   prj.ID,
		Name: prj.Name,
		Slug: prj.Slug,

		OrganizationName: prj.Org.Name,
		OrganizationSlug: prj.Org.Slug,
		OrganizationID:   prj.Org.ID,

		VcsInfo: VcsInfo{
			VcsURL:        "git://github.com/dummy-value",
			Provider:      prj.Org.Type,
			DefaultBranch: "main",
		},
	})
}

func (s *Service) deleteProject(c *gin.Context) {
	orgType := c.Param("org-type")
	switch orgType {
	case TypeGitHub, TypeBitbucket, TypeCircleCI:
		break
	default:
		msg(c, http.StatusBadRequest, "invalid org type")
		return
	}

	orgName := c.Param("org-name")
	projectName := c.Param("project-name")
	err := s.deleteProjectBySlug(orgType, orgName, projectName)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusNotFound, "project not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg(c, http.StatusOK, "ok")
}
