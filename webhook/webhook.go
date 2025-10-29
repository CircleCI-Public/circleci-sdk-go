package webhook

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
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

func (s *WebhookService) Get(ctx context.Context, webhookID string) (_ *Webhook, err error) {
	var webhook Webhook
	_, err = s.client.RequestHelper(ctx, http.MethodGet, "/webhook/"+webhookID, nil, &webhook)
	if err != nil {
		return nil, err
	}
	return &webhook, nil
}

func (s *WebhookService) List(ctx context.Context, scopeID string) (_ []Webhook, err error) {
	var nextPageToken string
	var webhookList []Webhook
	for {
		var response common.PaginatedResponse[Webhook]
		_, err = s.client.RequestHelper(ctx, http.MethodGet,
			fmt.Sprintf("/webhook?scope-id=%s&scope-type=project&page-token=%s", scopeID, nextPageToken),
			nil,
			&response,
		)
		if err != nil {
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

func (s *WebhookService) Create(ctx context.Context, newWebhook Webhook) (_ *Webhook, err error) {
	var webhook Webhook
	_, err = s.client.RequestHelper(ctx, http.MethodPost, "/webhook", newWebhook, &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

// Update - The scope cannot be updated
func (s *WebhookService) Update(ctx context.Context, newWebhook Webhook, webhookID string) (_ *Webhook, err error) {
	var webhook Webhook
	_, err = s.client.RequestHelper(ctx, http.MethodPut, "/webhook/"+webhookID, newWebhook, &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

func (s *WebhookService) Delete(ctx context.Context, webhookID string) (err error) {
	_, err = s.client.RequestHelper(ctx, http.MethodDelete, "/webhook/"+webhookID, nil, nil)
	return err
}
