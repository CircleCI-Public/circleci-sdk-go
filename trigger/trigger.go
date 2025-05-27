package trigger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Trigger struct {
	ID          string             `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
	CreatedAt   string             `json:"created_at,omitempty"`
	CheckoutRef string             `json:"checkout_ref,omitempty"`
	ConfigRef   string             `json:"config_ref,omitempty"`
	EventSource common.EventSource `json:"event_source,omitzero"`
	EventName   string             `json:"event_name,omitempty"`
	EventPreset string             `json:"event_preset,omitempty"`
}

type TriggerItems struct {
	Items []Trigger `json:"items"`
}

type TriggerService struct {
	client *client.Client
}

func NewTriggerService(c *client.Client) *TriggerService {
	return &TriggerService{client: c}
}

func (s *TriggerService) Get(project_id, trigger_id string) (*Trigger, error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/triggers/%s", project_id, trigger_id), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var trigger Trigger
	if err := json.NewDecoder(res.Body).Decode(&trigger); err != nil {
		return nil, err
	}
	return &trigger, nil
}

func (s *TriggerService) List(project_id, pipeline_id string) ([]Trigger, error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", project_id, pipeline_id), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var triggerItems TriggerItems
	if err := json.NewDecoder(res.Body).Decode(&triggerItems); err != nil {
		return nil, err
	}
	return triggerItems.Items, nil
}

func (s *TriggerService) Create(new_trigger Trigger, project_id, pipeline_id string) (*Trigger, error) {
	res, err := s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/projects/%s/pipeline-definitions/%s/triggers", project_id, pipeline_id), new_trigger)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var trigger Trigger
	if err := json.NewDecoder(res.Body).Decode(&trigger); err != nil {
		return nil, err
	}
	return &trigger, nil
}

func (s *TriggerService) Delete(project_id, trigger_id string) error {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/projects/%s/triggers/%s", project_id, trigger_id), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// The new trigger param can only have the eseential values:
// name, description, event_preset, checkout_ref, config_ref
// This are the only values that can be updated with this method
func (s *TriggerService) Update(new_trigger Trigger, project_id, trigger_id string) (*Trigger, error) {
	res, err := s.client.RequestHelper(http.MethodPatch, fmt.Sprintf("/projects/%s/triggers/%s", project_id, trigger_id), new_trigger)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var trigger Trigger
	if err := json.NewDecoder(res.Body).Decode(&trigger); err != nil {
		return nil, err
	}
	return &trigger, nil
}
