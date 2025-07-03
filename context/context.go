package context

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

type Context struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

//nolint:revive
type ContextRestriction struct {
	ID               string `json:"id,omitempty"`
	ContextId        string `json:"context_id,omitempty"`
	ProjectId        string `json:"project_id,omitempty"`
	Name             string `json:"name,omitempty"`
	RestrictionType  string `json:"restriction_type,omitempty"`
	RestrictionValue string `json:"restriction_value,omitempty"`
}

//nolint:revive
type ContextService struct {
	client *client.Client
}

func NewContextService(c *client.Client) *ContextService {
	return &ContextService{client: c}
}

func (s *ContextService) Get(contextID string) (_ *Context, err error) {
	res, err := s.client.RequestHelper(http.MethodGet, "/context/"+contextID, nil)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var context Context
	if err := json.NewDecoder(res.Body).Decode(&context); err != nil {
		return nil, err
	}
	return &context, nil
}

func (s *ContextService) List(organizationSlug string) (_ []Context, err error) {
	var nextPageToken string
	var contextList []Context
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context?owner-slug=%s&page-token=%s", organizationSlug, nextPageToken), nil)
		if err != nil {
			return nil, err
		}
		defer closer.ErrorHandler(res.Body, &err)

		var response common.PaginatedResponse[Context]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		contextList = append(contextList, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		nextPageToken = response.NextPageToken
	}
	return contextList, nil
}

func (s *ContextService) Create(organizationID, name string) (_ *Context, err error) {
	payload := map[string]any{
		"name": name,
		"owner": map[string]string{
			"id":   organizationID,
			"type": "organization",
		},
	}
	res, err := s.client.RequestHelper(http.MethodPost, "/context", payload)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var context Context
	if err := json.NewDecoder(res.Body).Decode(&context); err != nil {
		return nil, err
	}
	return &context, nil
}

func (s *ContextService) Delete(contextID string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, "/context/"+contextID, nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}

func (s *ContextService) GetRestrictions(contextID string) (_ []ContextRestriction, err error) {
	var nextPageToken string
	var contextRestrictionList []ContextRestriction
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context/%s/restrictions?page-token=%s", contextID, nextPageToken), nil)
		if err != nil {
			return nil, err
		}
		defer closer.ErrorHandler(res.Body, &err)

		var response common.PaginatedResponse[ContextRestriction]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		contextRestrictionList = append(contextRestrictionList, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		nextPageToken = response.NextPageToken
	}
	return contextRestrictionList, nil
}

func (s *ContextService) DeleteRestriction(contextID, restrictionID string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/context/%s/restrictions/%s", contextID, restrictionID), nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}

// CreateRestriction - context_id is the context this restriction applies to
// restriction_type is the type of resource this restrictions is related, either organization or project
// restriction_value is the id of the resource this restriction is related, the id of the org or project
func (s *ContextService) CreateRestriction(contextID, restrictionValue, restrictionType string) (_ *ContextRestriction, err error) {
	payload := map[string]string{
		"restriction_value": restrictionValue,
		"restriction_type":  restrictionType,
	}
	res, err := s.client.RequestHelper(http.MethodPost, "/context/"+contextID+"/restrictions", payload)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var contextRestriction ContextRestriction
	if err := json.NewDecoder(res.Body).Decode(&contextRestriction); err != nil {
		return nil, err
	}
	return &contextRestriction, nil
}
