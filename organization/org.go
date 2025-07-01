package organization

import (
	"encoding/json"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

type Organization struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	VcsType string `json:"vcs_type,omitempty"`
	Slug    string `json:"slug,omitempty"`
}

type OrganizationService struct {
	client *client.Client
}

func NewOrganizationService(c *client.Client) *OrganizationService {
	return &OrganizationService{client: c}
}

func (s *OrganizationService) Create(name, vcs_type string) (*Organization, error) {
	new_org := Organization{
		Name:    name,
		VcsType: vcs_type,
	}
	res, err := s.client.RequestHelper(http.MethodPost, "/organization", new_org)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var org Organization
	if err := json.NewDecoder(res.Body).Decode(&org); err != nil {
		return nil, err
	}
	return &org, nil
}

func (s *OrganizationService) Delete(org_id string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, "/organization/"+org_id, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}