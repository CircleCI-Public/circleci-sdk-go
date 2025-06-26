package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Webhook struct {
	Id            string       `json:"id,omitempty"`
	Name          string       `json:"name,omitempty"`
	Url           string       `json:"url,omitempty"`
	VerifyTls     *bool        `json:"verify-tls,omitempty"`
	SigningSecret string       `json:"signing-secret,omitempty"`
	UpdatedAt     string       `json:"updated-at,omitempty"`
	CreatedAt     string       `json:"created-at,omitempty"`
	Scope         common.Scope `json:"scope,omitempty"`
	Events        []string     `json:"events,omitempty"`
}

type WebhookService struct {
	client *client.Client
}

func NewWebhookService(c *client.Client) *WebhookService {
	return &WebhookService{client: c}
}

func (s *WebhookService) Get(webhook_id string) (*Webhook, error) {
	res, err := s.client.RequestHelper(http.MethodGet, "/webhook/"+webhook_id, nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) List(scope_id string) ([]Webhook, error) {
	var next_page_token string
	var webhook_list []Webhook
	for {
		res, err := s.client.RequestHelper(http.MethodGet,
			fmt.Sprintf("/webhook?scope-id=%s&scope-type=project&page-token=%s", scope_id, next_page_token), nil)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var response common.PaginatedResponse[Webhook]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		webhook_list = append(webhook_list, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		next_page_token = response.NextPageToken
	}
	return webhook_list, nil
}

func (s *WebhookService) Create(new_webhook Webhook) (*Webhook, error) {
	res, err := s.client.RequestHelper(http.MethodPost, "/webhook", new_webhook)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

// The scope cannot be updated
func (s *WebhookService) Update(new_webhook Webhook, webhook_id string) (*Webhook, error) {
	res, err := s.client.RequestHelper(http.MethodPut, "/webhook/"+webhook_id, new_webhook)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) Delete(webhook_id string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, "/webhook/"+webhook_id, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
