package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Project struct {
	Id               string         `json:"id"`
	Name             string         `json:"name"`
	Slug             string         `json:"slug"`
	OrganizationName string         `json:"organization_name"`
	OrganizationSlug string         `json:"organization_slug"`
	OrganizationId   string         `json:"organization_id"`
	VcsInfo          common.VcsInfo `json:"vcs_info"`
}

type AdvanceSettings struct {
	AutocancelBuilds           *bool     `json:"autocancel_builds,omitempty"`
	BuildForkPrs               *bool     `json:"build_fork_prs,omitempty"`
	DisableSSH                 *bool     `json:"disable_ssh,omitempty"`
	ForksReceiveSecretEnvVars  *bool     `json:"forks_receive_secret_env_vars,omitempty"`
	OSS                        *bool     `json:"oss,omitempty"`
	SetGithubStatus            *bool     `json:"set_github_status,omitempty"`
	SetupWorkflows             *bool     `json:"setup_workflows,omitempty"`
	WriteSettingsRequiresAdmin *bool     `json:"write_settings_requires_admin,omitempty"`
	PROnlyBranchOverrides      []string `json:"pr_only_branch_overrides,omitempty"`
}

type ProjectSettings struct {
	Advanced AdvanceSettings `json:"advanced"`
}

type ProjectService struct {
	client *client.Client
}

func NewProjectService(c *client.Client) *ProjectService {
	return &ProjectService{client: c}
}

func (s *ProjectService) Get(slug string) (*Project, error) {
	res, err := s.client.RequestHelper(http.MethodGet, "/project/"+slug, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var project Project
	if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
		return nil, err
	}
	return &project, nil
}

func (s *ProjectService) Create(project_name, organization_id string) (*Project, error) {
	payload := map[string]string{
		"name": project_name,
	}
	res, err := s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/organization/%s/project", organization_id), payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var project Project
	if err := json.NewDecoder(res.Body).Decode(&project); err != nil {
		return nil, err
	}
	slug := strings.Split(project.Slug, "/")
	if len(slug) == 3 && slug[1] == project.OrganizationName {
		// TODO: The URL here probably need to be used in a different way depending on how on premise works
		res, err := s.client.RequestHelperAbsolute(http.MethodGet, "https://circleci.com/api/v1.1/me", nil) 
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		var user common.User
		if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
			return nil, err
		}
		// TODO: The URL here probably need to be used in a different way depending on how on premise works
		url := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s/follow", strings.ToLower(project.VcsInfo.Provider), user.Login, project.Name)
		res, err = s.client.RequestHelperAbsolute(http.MethodPost, url, nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
	}
	return &project, nil
}

// Only standalone projects can be deleted
func (s *ProjectService) Delete(slug string) (error) {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/project/%s", slug), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// Settings are only available for standalone projects
func (s *ProjectService) GetSettings(provider, organization, project string) (*ProjectSettings, error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var settings ProjectSettings
	if err := json.NewDecoder(res.Body).Decode(&settings); err != nil {
		return nil, err
	}
	return &settings, nil
}

// Settings are only available for standalone projects
func (s *ProjectService) UpdateSettings(new_settings ProjectSettings, provider, organization, project string) (*ProjectSettings, error) {
	res, err := s.client.RequestHelper(http.MethodPatch, fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project), new_settings)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var settings ProjectSettings
	if err := json.NewDecoder(res.Body).Decode(&settings); err != nil {
		return nil, err
	}
	return &settings, nil
}
