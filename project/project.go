package project

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Project struct {
	// nolint:revive // introduced before linter
	Id               string `json:"id"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	OrganizationName string `json:"organization_name"`
	OrganizationSlug string `json:"organization_slug"`
	// nolint:revive // introduced before linter
	OrganizationId string         `json:"organization_id"`
	VcsInfo        common.VcsInfo `json:"vcs_info"`
}

type AdvanceSettings struct {
	AutocancelBuilds           *bool    `json:"autocancel_builds,omitempty"`
	BuildForkPrs               *bool    `json:"build_fork_prs,omitempty"`
	DisableSSH                 *bool    `json:"disable_ssh,omitempty"`
	ForksReceiveSecretEnvVars  *bool    `json:"forks_receive_secret_env_vars,omitempty"`
	OSS                        *bool    `json:"oss,omitempty"`
	SetGithubStatus            *bool    `json:"set_github_status,omitempty"`
	SetupWorkflows             *bool    `json:"setup_workflows,omitempty"`
	WriteSettingsRequiresAdmin *bool    `json:"write_settings_requires_admin,omitempty"`
	PROnlyBranchOverrides      []string `json:"pr_only_branch_overrides,omitempty"`
}

// nolint:revive // introduced before linter
type ProjectSettings struct {
	Advanced AdvanceSettings `json:"advanced"`
}

// nolint:revive // introduced before linter
type ProjectService struct {
	client *client.Client
}

func NewProjectService(c *client.Client) *ProjectService {
	return &ProjectService{client: c}
}

func (s *ProjectService) Get(ctx context.Context, slug string) (_ *Project, err error) {
	var project Project
	_, err = s.client.RequestHelper(ctx, http.MethodGet, "/project/"+slug, nil, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (s *ProjectService) Create(ctx context.Context, projectName, organizationID string) (_ *Project, err error) {
	payload := map[string]string{
		"name": projectName,
	}
	var project Project
	_, err = s.client.RequestHelper(ctx, http.MethodPost, fmt.Sprintf("/organization/%s/project", organizationID), payload, &project)
	if err != nil {
		return nil, err
	}

	slug := strings.Split(project.Slug, "/")
	if len(slug) == 3 && slug[1] == project.OrganizationName {
		// TODO: The URL here probably need to be used in a different way depending on how on premise works
		var user common.User
		_, err = s.client.RequestHelperAbsolute(ctx, http.MethodGet, "https://circleci.com/api/v1.1/me", nil, &user)
		if err != nil {
			return nil, err
		}
		// TODO: The URL here probably need to be used in a different way depending on how on premise works
		url := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s/follow", strings.ToLower(project.VcsInfo.Provider), user.Login, project.Name)
		_, err = s.client.RequestHelperAbsolute(ctx, http.MethodPost, url, nil, nil)
		if err != nil {
			return nil, err
		}
	}
	return &project, nil
}

// Delete - Only standalone projects can be deleted
func (s *ProjectService) Delete(ctx context.Context, slug string) (err error) {
	_, err = s.client.RequestHelper(ctx, http.MethodDelete, fmt.Sprintf("/project/%s", slug), nil, nil)
	return err
}

// GetSettings - Settings are only available for standalone projects
func (s *ProjectService) GetSettings(ctx context.Context, provider, organization, project string) (_ *ProjectSettings, err error) {
	var settings ProjectSettings
	_, err = s.client.RequestHelper(ctx, http.MethodGet, fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project), nil, &settings)
	if err != nil {
		return nil, err
	}
	return &settings, nil
}

// UpdateSettings - Settings are only available for standalone projects
func (s *ProjectService) UpdateSettings(ctx context.Context, newSettings ProjectSettings, provider, organization, project string) (_ *ProjectSettings, err error) {
	var settings ProjectSettings
	_, err = s.client.RequestHelper(ctx, http.MethodPatch, fmt.Sprintf("/project/%s/%s/%s/settings", provider, organization, project), newSettings, &settings)
	if err != nil {
		return nil, err
	}

	return &settings, nil
}
