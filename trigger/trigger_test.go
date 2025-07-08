package trigger

import (
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

const (
	knownPipelineID = "bee796a0-7ec2-478c-ab87-6a5039d7a216"
	knownProjectID  = "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
)

func TestListTrigger(t *testing.T) {
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	trs, err := triggerService.List(knownProjectID, knownPipelineID)
	assert.Assert(t, err)

	t.Log(trs)
}

func TestFullTriggerNew(t *testing.T) {
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	pipelineID := knownPipelineID
	projectID := knownProjectID
	newTrigger := Trigger{
		EventName: "Event Name",
		EventSource: common.EventSource{
			Provider: "webhook",
			Webhook: common.Webhook{
				Sender: "Test",
			},
		},
		EventPreset: "all-pushes",
		ConfigRef:   "main",
		CheckoutRef: "main",
	}
	triggerCreated, err := triggerService.Create(newTrigger, projectID, pipelineID)
	assert.Assert(t, err)

	idNewTrigger := triggerCreated.ID
	triggerToUpdate := Trigger{
		Name:        "trigger-updated",
		Description: "Updated description",
	}

	triggerUpdated, err := triggerService.Update(triggerToUpdate, projectID, idNewTrigger)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(triggerUpdated.Description, "Updated description"))

	triggerFetched, err := triggerService.Get(projectID, idNewTrigger)
	assert.Assert(t, err)
	t.Log(triggerFetched)

	err = triggerService.Delete(projectID, idNewTrigger)
	assert.Assert(t, err)

	triggerFetched, err = triggerService.Get(projectID, idNewTrigger)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(triggerFetched))
}

func TestFullTrigger(t *testing.T) {
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	pipelineID := "bee796a0-7ec2-478c-ab87-6a5039d7a216"
	projectID := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	newTrigger := Trigger{
		Name:        "test-trigger",
		Description: "Test trigger from SDK",
		EventSource: common.EventSource{
			Provider: "github_app",
			Repo: common.Repo{
				ExternalId: "952038793",
			},
		},
		EventPreset: "all-pushes",
		ConfigRef:   "main",
		CheckoutRef: "main",
	}
	triggerCreated, err := triggerService.Create(newTrigger, projectID, pipelineID)
	assert.Assert(t, err)

	idNewTrigger := triggerCreated.ID
	triggerToUpdate := Trigger{
		Name:        "trigger-updated",
		Description: "Updated description",
	}

	triggerUpdated, err := triggerService.Update(triggerToUpdate, projectID, idNewTrigger)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(triggerUpdated.Description, "Updated description"))

	triggerFetched, err := triggerService.Get(projectID, idNewTrigger)
	assert.Assert(t, err)

	t.Log(triggerFetched)
	err = triggerService.Delete(projectID, idNewTrigger)
	assert.Assert(t, err)

	triggerFetched, err = triggerService.Get(projectID, idNewTrigger)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(triggerFetched))
}
