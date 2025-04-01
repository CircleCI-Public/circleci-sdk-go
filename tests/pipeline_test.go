package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/pipeline"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
)

func TestFullPipeline(t *testing.T) {
	token := os.Getenv("LOCAL_CCI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	pipelineService := pipeline.NewPipelineService(client)

	project_id := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	new_repo := common.Repo{
		ExternalId: "952038793",
	}
	new_pipeline := pipeline.Pipeline{
		Name: "test-pipeline",
		Description: "Test pipeline from SDK",
		ConfigSource: common.ConfigSource{
			Provider: "github_app",
			Repo: new_repo,
			FilePath: ".circleci/config.yml",
		},
		CheckoutSource: common.CheckoutSource{
			Provider: "github_app",
			Repo: new_repo,
		},
	}
	pipeline_created, err := pipelineService.Create(new_pipeline, project_id)
	if err != nil {
		t.Log(err)
		t.Error("Error creating pipeline")
		t.FailNow()
	}
	id_new_pipeline := pipeline_created.ID
	pipeline_to_update := pipeline.Pipeline{
		Name: "pipeline-updated",
		Description: "Updated description",
		ConfigSource: common.ConfigSource{
			FilePath: ".circleci/config2.yml",
		},
	}
	pieline_updated, err := pipelineService.Update(pipeline_to_update, project_id, id_new_pipeline)
	if err != nil {
		t.Log(err)
		t.Error("Error updating pipeline")
		t.FailNow()
	}
	if pieline_updated.Description != "Updated description" {
		t.Error("Pipeline was not updated")
		t.FailNow()
	}
	pipeline_fetched, err := pipelineService.Get(project_id, id_new_pipeline)
	if err != nil {
		t.Log(err)
		t.Error("Error getting pipeline")
		t.FailNow()
	}
	err = pipelineService.Delete(project_id, id_new_pipeline)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting pipeline")
		t.FailNow()
	}
	pipeline_fetched, err = pipelineService.Get(project_id, id_new_pipeline)
	t.Log(err)
	if pipeline_fetched != nil {
		t.Error("Pipeline was not deleted")
		t.FailNow()
	}
}

func TestListPipeline(t *testing.T) {
	token := os.Getenv("LOCAL_CCI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	pipelineService := pipeline.NewPipelineService(client)

	project_id := "43155b9d-afdf-4616-a8c6-b32952416104"
	ps, err := pipelineService.List(project_id)
	if err != nil {
		t.Log(err)
		t.Error("Error getting pipelines")
		t.FailNow()
	}
	t.Log(ps)
	if len(ps) != 1 {
		t.Errorf("Size of pipelines is not correct")
	}
}
