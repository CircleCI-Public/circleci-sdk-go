package pipeline

import (
	"encoding/json"
	"log"
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

type PipelineItems struct {
	Items []Pipeline `json:"items"`
}

type PipelineService struct {
	client *client.Client
}

func NewPipelineService(c *client.Client) *PipelineService {
	return &PipelineService{client: c}
}

func (s *PipelineService) Get(project_id, pipeline_id string) (*Pipeline, error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", project_id, pipeline_id), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineService) List(project_id string) ([]Pipeline, error) {
	res, err := s.client.RequestHelper(http.MethodGet, fmt.Sprintf("/projects/%s/pipeline-definitions", project_id), nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pipelineItems PipelineItems
	if err := json.NewDecoder(res.Body).Decode(&pipelineItems); err != nil {
		return nil, err
	}
	return pipelineItems.Items, nil
}

func (s *PipelineService) Create(new_pipeline Pipeline, project_id string) (*Pipeline, error) {
	res, err := s.client.RequestHelper(http.MethodPost, fmt.Sprintf("/projects/%s/pipeline-definitions", project_id), new_pipeline)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (s *PipelineService) Delete(project_id, pipeline_id string) (error) {
	res, err := s.client.RequestHelper(http.MethodDelete, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", project_id, pipeline_id), nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

// The new pipeline param can only have the eseential values:
// name, description, config_source.file_path, checkout_source.provider, checkout_source.repo.external_id
// This are the only values that can be updated with this method, and the objet passed can only have these defined
func (s *PipelineService) Update(new_pipeline Pipeline, project_id, pipeline_id string) (*Pipeline, error) {
	payload, err := json.Marshal(new_pipeline)
	if err != nil {
		return nil, err
	}
	log.Print(string(payload))
	res, err := s.client.RequestHelper(http.MethodPatch, fmt.Sprintf("/projects/%s/pipeline-definitions/%s", project_id, pipeline_id), new_pipeline)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var pipeline Pipeline
	if err := json.NewDecoder(res.Body).Decode(&pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}
