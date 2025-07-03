package organization

import (
	"encoding/json"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
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
	res, err := s.client.RequestHelper(http.MethodPost, "/organization", Organization{
		Name:    name,
		VcsType: vcsType,
	})
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	org = &Organization{}
	if err := json.NewDecoder(res.Body).Decode(org); err != nil {
		return nil, err
	}
	return org, nil
}

func (s *OrganizationService) Delete(orgID string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, "/organization/"+orgID, nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}
