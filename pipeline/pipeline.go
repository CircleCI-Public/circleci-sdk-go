package pipeline

import (
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

type Pipeline struct {
	ID             string                `json:"id,omitempty"`
	Name           string                `json:"name,omitempty"`
	Description    string                `json:"description,omitempty"`
	CreatedAt      string                `json:"created_at,omitempty"`
	ConfigSource   common.ConfigSource   `json:"config_source,omitzero"`
	CheckoutSource common.CheckoutSource `json:"checkout_source,omitzero"`
}

// nolint:revive // introduced before linter
type PipelineItems struct {
	Items []Pipeline `json:"items"`
}

// nolint:revive // introduced before linter
type PipelineService struct {
	client *client.Client
}

func NewPipelineService(c *client.Client) *PipelineService {
	return &PipelineService{client: c}
}

func (s *PipelineService) Get(projectID, pipelineID string) (_ *Pipeline, err error) {
	var pipeline Pipeline
	_, err = s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), nil, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}

func (s *PipelineService) List(projectID string) (_ []Pipeline, err error) {
	var pipelineItems PipelineItems
	_, err = s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions", projectID), nil, &pipelineItems)
	if err != nil {
		return nil, err
	}

	return pipelineItems.Items, nil
}

func (s *PipelineService) Create(newPipeline Pipeline, projectID string) (_ *Pipeline, err error) {
	var pipeline Pipeline
	_, err = s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/projects/%s/pipeline-definitions", projectID), newPipeline, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}

func (s *PipelineService) Delete(projectID, pipelineID string) (err error) {
	_, err = s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), nil, nil)
	return err
}

// Update - The new pipeline param can only have the eseential values:
// name, description, config_source.file_path, checkout_source.provider, checkout_source.repo.external_id
// This are the only values that can be updated with this method, and the objet passed can only have these defined
func (s *PipelineService) Update(newPipeline Pipeline, projectID, pipelineID string) (_ *Pipeline, err error) {
	var pipeline Pipeline
	_, err = s.client.RequestHelper(http.MethodPatch, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), newPipeline, &pipeline)
	if err != nil {
		return nil, err
	}

	return &pipeline, nil
}
