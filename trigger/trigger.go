package trigger

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Trigger struct {
	ID          string             `json:"id,omitempty"`
	CreatedAt   string             `json:"created_at,omitempty"`
	CheckoutRef string             `json:"checkout_ref,omitempty"`
	ConfigRef   string             `json:"config_ref,omitempty"`
	EventSource common.EventSource `json:"event_source,omitzero"`
	EventName   string             `json:"event_name,omitempty"`
	EventPreset string             `json:"event_preset,omitempty"`
	Disabled    *bool              `json:"disabled,omitempty"`
}

// nolint:revive // introduced before linter
type TriggerItems struct {
	Items []Trigger `json:"items"`
}

// nolint:revive // introduced before linter
type TriggerService struct {
	client *client.Client
}

func NewTriggerService(c *client.Client) *TriggerService {
	return &TriggerService{client: c}
}

func (s *TriggerService) Get(ctx context.Context, projectID, triggerID string) (_ *Trigger, err error) {
	var trigger Trigger
	_, err = s.client.RequestHelper(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/triggers/%s", projectID, triggerID), nil, &trigger)
	if err != nil {
		return nil, err
	}

	return &trigger, nil
}

func (s *TriggerService) List(ctx context.Context, projectID, pipelineID string) (_ []Trigger, err error) {
	var triggerItems TriggerItems
	_, err = s.client.RequestHelper(ctx, http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", projectID, pipelineID), nil, &triggerItems)
	if err != nil {
		return nil, err
	}

	return triggerItems.Items, nil
}

func (s *TriggerService) Create(ctx context.Context, newTrigger Trigger, projectID, pipelineID string) (_ *Trigger, err error) {
	var trigger Trigger
	_, err = s.client.RequestHelper(ctx, http.MethodPost, fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", projectID, pipelineID), newTrigger, &trigger)
	if err != nil {
		return nil, err
	}

	return &trigger, nil
}

func (s *TriggerService) Delete(ctx context.Context, projectID, triggerID string) (err error) {
	_, err = s.client.RequestHelper(ctx, http.MethodDelete, fmt.Sprintf("/projects/%s/triggers/%s", projectID, triggerID), nil, nil)
	return err
}

// Update The new trigger param can only have the esseential values:
// name, description, event_preset, checkout_ref, config_ref, disabled
// This are the only values that can be updated with this method
func (s *TriggerService) Update(ctx context.Context, newTrigger Trigger, projectID, triggerID string) (_ *Trigger, err error) {
	var trigger Trigger
	_, err = s.client.RequestHelper(ctx, http.MethodPatch, fmt.Sprintf("/projects/%s/triggers/%s", projectID, triggerID), newTrigger, &trigger)
	if err != nil {
		return nil, err
	}

	return &trigger, nil
}
