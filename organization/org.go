package organization

import (
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

type Organization struct {
	// nolint:revive // introduced before linter
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	VcsType string `json:"vcs_type,omitempty"`
	Slug    string `json:"slug,omitempty"`
}

// nolint:revive // introduced before linter
type OrganizationService struct {
	client *client.Client
}

func NewOrganizationService(c *client.Client) *OrganizationService {
	return &OrganizationService{client: c}
}

func (s *OrganizationService) Create(name, vcsType string) (org *Organization, err error) {
	org = &Organization{}
	_, err = s.client.RequestHelper(http.MethodPost, "/organization", Organization{
		Name:    name,
		VcsType: vcsType,
	}, org)
	if err != nil {
		return nil, err
	}
	return org, nil
}

func (s *OrganizationService) Delete(orgID string) (err error) {
	_, err = s.client.RequestHelper(http.MethodDelete, "/organization/"+orgID, nil, nil)
	return err
}
