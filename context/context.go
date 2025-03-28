package context

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Context struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type ContextRestriction struct {
	ID               string `json:"id,omitempty"`
	ContextId        string `json:"context_id,omitempty"`
	ProjectId        string `json:"project_id,omitempty"`
	Name             string `json:"name,omitempty"`
	RestrictionType  string `json:"restriction_type,omitempty"`
	RestrictionValue string `json:"restriction_value,omitempty"`
}

type ContextService struct {
	client *client.Client
}

func NewContextService(c *client.Client) *ContextService {
	return &ContextService{client: c}
}

func (s *ContextService) Get(context_id string) (*Context, error) {
	res, err := s.client.RequestHelper(http.MethodGet, "/context/"+context_id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var context Context
	if err := json.NewDecoder(res.Body).Decode(&context); err != nil {
		return nil, err
	}
	return &context, nil
}

func (s *ContextService) List(organization_slug string) ([]Context, error) {
	var next_page_token string
	var context_list []Context
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context?owner-slug=%s&page-token=%s", organization_slug, next_page_token), nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var response common.PaginatedResponse[Context]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		context_list = append(context_list, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		next_page_token = response.NextPageToken
	}
	return context_list, nil
}

func (s *ContextService) Create(organization_id, name string) (*Context, error) {
	payload := map[string]any{
		"name": name,
		"owner": map[string]string{
			"id":   organization_id,
			"type": "organization",
		},
	}
	res, err := s.client.RequestHelper(http.MethodPost, "/context", payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var context Context
	if err := json.NewDecoder(res.Body).Decode(&context); err != nil {
		return nil, err
	}
	return &context, nil
}

func (s *ContextService) Delete(context_id string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, "/context/"+context_id, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (s *ContextService) GetRestrictions(context_id string) ([]ContextRestriction, error) {
	var next_page_token string
	var context_restriction_list []ContextRestriction
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context/%s/restrictions?page-token=%s", context_id, next_page_token), nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var response common.PaginatedResponse[ContextRestriction]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		context_restriction_list = append(context_restriction_list, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		next_page_token = response.NextPageToken
	}
	return context_restriction_list, nil
}

func (s *ContextService) DeleteRestriction(context_id, restriction_id string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/context/%s/restrictions/%s", context_id, restriction_id), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// context_id is the context this restriction applies to
// restriction_type is the type of resource this restrictions is related, either organization or project
// restriction_value is the id of the resource this restriction is related, the id of the org or project
func (s *ContextService) CreateRestriction(context_id, restriction_value, restriction_type string) (*ContextRestriction, error) {
	payload := map[string]string{
		"restriction_value": restriction_value,
		"restriction_type": restriction_type,
	}
	res, err := s.client.RequestHelper(http.MethodPost, "/context/"+context_id+"/restrictions", payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var contextRestriction ContextRestriction
	if err := json.NewDecoder(res.Body).Decode(&contextRestriction); err != nil {
		return nil, err
	}
	return &contextRestriction, nil
}
