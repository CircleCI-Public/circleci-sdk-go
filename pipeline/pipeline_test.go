package pipeline

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

func TestFullPipeline(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	pipelineService := NewPipelineService(c)

	projectID := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	newRepo := common.Repo{
		ExternalId: "952038793",
	}
	newPipeline := Pipeline{
		Name:        "test-pipeline",
		Description: "Test pipeline from SDK",
		ConfigSource: common.ConfigSource{
			Provider: "github_app",
			Repo:     newRepo,
			FilePath: ".circleci/config.yml",
		},
		CheckoutSource: common.CheckoutSource{
			Provider: "github_app",
			Repo:     newRepo,
		},
	}
	pipelineCreated, err := pipelineService.Create(newPipeline, projectID)
	assert.Assert(t, err)

	idNewPipeline := pipelineCreated.ID
	pipelineToUpdate := Pipeline{
		Name:        "pipeline-updated",
		Description: "Updated description",
		ConfigSource: common.ConfigSource{
			FilePath: ".circleci/config2.yml",
		},
	}
	pielineUpdated, err := pipelineService.Update(pipelineToUpdate, projectID, idNewPipeline)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(pielineUpdated.Description, "Updated description"))

	pipelineFetched, err := pipelineService.Get(projectID, idNewPipeline)
	assert.Assert(t, err)
	assert.Check(t, pipelineFetched != nil)

	err = pipelineService.Delete(projectID, idNewPipeline)
	assert.Assert(t, err)

	pipelineFetched, err = pipelineService.Get(projectID, idNewPipeline)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(pipelineFetched))
}

func TestListPipeline(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	pipelineService := NewPipelineService(c)

	projectID := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	ps, err := pipelineService.List(projectID)
	assert.Assert(t, err)
	assert.Assert(t, cmp.Len(ps, 1))
}
