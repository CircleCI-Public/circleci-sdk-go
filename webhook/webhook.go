package webhook

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
)

type Webhook struct {
	//nolint:revive // introduced before linter
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	// nolint:revive // introduced before linter
	Url string `json:"url,omitempty"`
	//nolint:revive // introduced before linter
	VerifyTls     *bool        `json:"verify-tls,omitempty"`
	SigningSecret string       `json:"signing-secret,omitempty"`
	UpdatedAt     string       `json:"updated-at,omitempty"`
	CreatedAt     string       `json:"created-at,omitempty"`
	Scope         common.Scope `json:"scope,omitempty"`
	Events        []string     `json:"events,omitempty"`
}

// nolint:revive // introduced before linter
type WebhookService struct {
	client *client.Client
}

func NewWebhookService(c *client.Client) *WebhookService {
	return &WebhookService{client: c}
}

func (s *WebhookService) Get(webhookID string) (_ *Webhook, err error) {
	res, err := s.client.RequestHelper(http.MethodGet, "/webhook/"+webhookID, nil)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) List(scopeID string) (_ []Webhook, err error) {
	var nextPageToken string
	var webhookList []Webhook
	for {
		res, err := s.client.RequestHelper(http.MethodGet,
			fmt.Sprintf("/webhook?scope-id=%s&scope-type=project&page-token=%s", scopeID, nextPageToken), nil)
		if err != nil {
			return nil, err
		}
		defer closer.ErrorHandler(res.Body, &err)

		var response common.PaginatedResponse[Webhook]
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			return nil, err
		}
		webhookList = append(webhookList, response.Items...)
		if response.NextPageToken == "" {
			break
		}
		nextPageToken = response.NextPageToken
	}
	return webhookList, nil
}

func (s *WebhookService) Create(newWebhook Webhook) (_ *Webhook, err error) {
	res, err := s.client.RequestHelper(http.MethodPost, "/webhook", newWebhook)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

// Update - The scope cannot be updated
func (s *WebhookService) Update(newWebhook Webhook, webhookID string) (_ *Webhook, err error) {
	res, err := s.client.RequestHelper(http.MethodPut, "/webhook/"+webhookID, newWebhook)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)
	var webhook Webhook
	if err := json.NewDecoder(res.Body).Decode(&webhook); err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) Delete(webhookID string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, "/webhook/"+webhookID, nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}
