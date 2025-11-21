package fakecircle

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type project struct {
	Org     *org
	ID      uuid.UUID
	Name    string
	EnvVars []EnvVarProject
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

func (s *Service) getProjectById(id uuid.UUID) (project, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	envPrj, ok := s.projects[id]
	if !ok {
		return project{}, false
	}

	return project{
		ID:        envPrj.ID,
		Name:      envPrj.Name,
		EnvVars:   slices.Clone(envPrj.EnvVars),
		Org:       envPrj.Org,
	}, true
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

type NewEnvVarProject struct {
	Name string
	Value    string
}

type EnvVarProject struct {
	Name      string
	Value     string
	CreatedAt time.Time
}

func (p *project) addEnv(ev NewEnvVarProject) (EnvVarProject, error) {
	if slices.ContainsFunc(p.EnvVars, func(e EnvVarProject) bool {
		return e.Name == ev.Name
	}) {
		return EnvVarProject{}, errDuplicate
	}

	now := time.Now()
	e := EnvVarProject{
		Name:      ev.Name,
		Value:     ev.Value,
		CreatedAt: now,
	}
	p.EnvVars = append(p.EnvVars, e)
	return e, nil
}

func (p *project) deleteEnv(ev string) {
	p.EnvVars = slices.DeleteFunc(p.EnvVars, func(e EnvVarProject) bool {
		return e.Name == ev
	})
}

// func (s *Service) deleteEnvProject(orgType, orgName, projectName string) error {
// 	s.mu.Lock()
// 	defer s.mu.Unlock()

// 	prj, err := s.projectBySlug(orgType, orgName, projectName)
// 	if err != nil {
// 		return errNotFound
// 	}
// 	id := prj.ID
// 	foundPrj, ok := s.projects[id]

// 	if !ok {
// 		return errNotFound
// 	}

// 	_, ok = foundPrj.Org.projects[id]
// 	if !ok {
// 		return errNotFound
// 	}

// 	delete(foundPrj.Org.projects, id)
// 	delete(s.projects, id)

// 	return nil
// }

func (s *Service) AddProjectEnv(projectID uuid.UUID, ev NewEnvVarProject) (EnvVarProject, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	envPrj, ok := s.projects[projectID]
	if !ok {
		return EnvVarProject{}, errNotFound
	}

	return envPrj.addEnv(ev)
}

func (s *Service) deleteProjectEnvVar(id uuid.UUID, name string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	envPrj, ok := s.projects[id]
	if !ok {
		return errNotFound
	}

	envPrj.deleteEnv(name)
	return nil
}

func (s *Service) getProjectEnv(c *gin.Context) {
	type response struct {
		Value     string    `json:"value"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created-at"`
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
	if err != nil {
		msg(c, http.StatusNotFound, "project not found")
		return
	}
	id := prj.ID
	envPrj, ok := s.projects[id]

	if !ok {
		msg(c, http.StatusNotFound, "project not found")
		return
	}

	res := make([]response, 0, len(envPrj.EnvVars))
	for _, ev := range envPrj.EnvVars {
		res = append(res, response{
			Name:      ev.Name,
			CreatedAt: ev.CreatedAt,
			Value:     ev.Value,
		})
	}
	c.JSON(http.StatusOK, newListResponse(res))
}

func (s *Service) postProjectEnv(c *gin.Context) {
	type response struct {
		Name      string    `json:"name"`
		Value     string	`json:"value"`
		CreatedAt time.Time `json:"created-at"`
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

	var body struct {
		Value string `json:"value" binding:"required"`
		Name string  `json:"name" binding:"required"`
	}
	err := c.BindJSON(&body)
	if mapBadRequest(c, "bad request", err) {
		return
	}
	prj, err := s.projectBySlug(orgType, orgName, projectName)
	if err != nil {
		msg(c, http.StatusNotFound, "project not found")
		return
	}
	projectID := prj.ID
	ev, err := s.AddProjectEnv(projectID, NewEnvVarProject{
		Name:  body.Name,
		Value: body.Value,
	})

	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "project not found")
		return
	case errors.Is(err, errDuplicate):
		msg(c, http.StatusBadRequest, "env var already exists")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, response{
		Name:      ev.Name,
		CreatedAt: ev.CreatedAt,
		Value:     ev.Value,
	})
}

func (s *Service) deleteProjectEnv(c *gin.Context) {
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
	if err != nil {
		msg(c, http.StatusNotFound, "project not found")
		return
	}
	projectID := prj.ID

	envVarName := c.Param("env-var")
	if envVarName == "" {
		msg(c, http.StatusBadRequest, "bad request")
		return
	}

	err = s.deleteProjectEnvVar(projectID, envVarName)
	switch {
	case errors.Is(err, errNotFound):
		msg(c, http.StatusBadRequest, "project not found")
		return
	case err != nil:
		msg(c, http.StatusInternalServerError, err.Error())
		return
	}

	msg(c, http.StatusOK, "ok")
}
