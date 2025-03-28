package env

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type EnvVariable struct {
	Variable  string `json:"variable,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	ContextId string `json:"context_id,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type EnvService struct {
	client *client.Client
}

func NewEnvService(c *client.Client) *EnvService {
	return &EnvService{client: c}
}

func (s *EnvService) List(context_id string) ([]EnvVariable, error) {
	var next_page_token string
	var context_list []EnvVariable
	for {
		res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/context/%s/environment-variable?page-token=%s", context_id, next_page_token), nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var response common.PaginatedResponse[EnvVariable]
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

func (s *EnvService) Create(context_id, value, name string) (*EnvVariable, error) {
	payload := map[string]string{
		"value": value,
	}
	res, err := s.client.RequestHelper(http.MethodPut, fmt.Sprintf("/context/%s/environment-variable/%s", context_id, name), payload)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var envVariable EnvVariable
	if err := json.NewDecoder(res.Body).Decode(&envVariable); err != nil {
		return nil, err
	}
	return &envVariable, nil
}

func (s *EnvService) Delete(context_id, name string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/context/%s/environment-variable/%s", context_id, name), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
