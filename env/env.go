package env

import (
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
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
		var response common.PaginatedResponse[EnvVariable]
		_, err = s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context/%s/environment-variable?page-token=%s", contextID, nextPageToken), nil, &response)
		if err != nil {
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
	var envVariable EnvVariable
	_, err = s.client.RequestHelper(http.MethodPut, fmt.Sprintf("/context/%s/environment-variable/%s", contextID, name), payload, &envVariable)
	if err != nil {
		return nil, err
	}
	return &envVariable, nil
}

func (s *EnvService) Delete(contextID, name string) (err error) {
	_, err = s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/context/%s/environment-variable/%s", contextID, name), nil, nil)
	return err
}
