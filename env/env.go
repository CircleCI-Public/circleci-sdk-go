package env

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

//nolint:revive
type EnvVariable struct {
	Variable  string `json:"variable,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ContextId string `json:"context_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

//nolint:revive
type EnvService struct {
	client *client.Client
}

func NewEnvService(c *client.Client) *EnvService {
	return &EnvService{client: c}
}

func (s *EnvService) List(contextID string) (_ []EnvVariable, err error) {
	var nextPageToken string
	var contextList []EnvVariable
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context/%s/environment-variable?page-token=%s", contextID, nextPageToken), nil)
		if err != nil {
			return nil, err
		}
		defer closer.ErrorHandler(res.Body, &err)

		var response common.PaginatedResponse[EnvVariable]
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

func (s *EnvService) Create(contextID, value, name string) (_ *EnvVariable, err error) {
	payload := map[string]string{
		"value": value,
	}
	res, err := s.client.RequestHelper(http.MethodPut, fmt.Sprintf("/context/%s/environment-variable/%s", contextID, name), payload)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var envVariable EnvVariable
	if err := json.NewDecoder(res.Body).Decode(&envVariable); err != nil {
		return nil, err
	}
	return &envVariable, nil
}

func (s *EnvService) Delete(contextID, name string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/context/%s/environment-variable/%s", contextID, name), nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}
