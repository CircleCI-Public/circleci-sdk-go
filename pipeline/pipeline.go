package pipeline

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/closer"
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
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), nil)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineService) List(projectID string) (_ []Pipeline, err error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions", projectID), nil)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var pipelineItems PipelineItems
	if err := json.NewDecoder(res.Body).Decode(&pipelineItems); err != nil {
		return nil, err
	}
	return pipelineItems.Items, nil
}

func (s *PipelineService) Create(newPipeline Pipeline, projectID string) (_ *Pipeline, err error) {
	res, err := s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/projects/%s/pipeline-definitions", projectID), newPipeline)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)

	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineService) Delete(projectID, pipelineID string) (err error) {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), nil)
	if err != nil {
		return err
	}
	defer closer.ErrorHandler(res.Body, &err)

	return nil
}

// Update - The new pipeline param can only have the eseential values:
// name, description, config_source.file_path, checkout_source.provider, checkout_source.repo.external_id
// This are the only values that can be updated with this method, and the objet passed can only have these defined
func (s *PipelineService) Update(newPipeline Pipeline, projectID, pipelineID string) (_ *Pipeline, err error) {
	res, err := s.client.RequestHelper(http.MethodPatch, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", projectID, pipelineID), newPipeline)
	if err != nil {
		return nil, err
	}
	defer closer.ErrorHandler(res.Body, &err)
	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}
