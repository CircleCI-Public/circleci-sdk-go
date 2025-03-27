package project

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

type Project struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type AdvanceSettings struct {
	AutocancelBuilds           bool     `json:"autocancel_builds,omitempty"`
	BuildForkPrs               bool     `json:"build_fork_prs,omitempty"`
	DisableSSH                 bool     `json:"disable_ssh,omitempty"`
	ForksReceiveSecretEnvVars  bool     `json:"forks_receive_secret_env_vars,omitempty"`
	OSS                        bool     `json:"oss,omitempty"`
	SetGithubStatus            bool     `json:"set_github_status,omitempty"`
	SetupWorkflows             bool     `json:"setup_workflows,omitempty"`
	WriteSettingsRequiresAdmin bool     `json:"write_settings_requires_admin,omitempty"`
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

func (s *ProjectService) Create(project, organization, provider string) (*ProjectSettings, error) {
	res, err := s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/project/%s/%s/%s", provider, organization, project), nil)
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

func (s *ProjectService) UpdateSettings(new_settings ProjectSettings, project, organization, provider string) (*ProjectSettings, error) {
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
