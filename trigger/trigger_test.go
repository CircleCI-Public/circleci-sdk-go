package trigger

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

const (
	knownPipelineID         = "bee796a0-7ec2-478c-ab87-6a5039d7a216"
	knownProjectID          = "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	knownSchedulePipelineID = "FILL_IN_PIPELINE_DEFINITION_ID_THAT_SUPPORTS_SCHEDULE_TRIGGERS"
)

func TestListTrigger(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	trs, err := triggerService.List(ctx, knownProjectID, knownPipelineID)
	assert.Assert(t, err)

	t.Log(trs)
}

func TestFullTriggerNew(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	pipelineID := knownPipelineID
	projectID := knownProjectID
	newTrigger := Trigger{
		EventName: "Event Name",
		EventSource: common.EventSource{
			Provider: "webhook",
			Webhook: common.Webhook{
				Sender: "Test sender",
			},
		},
		EventPreset: "all-pushes",
		ConfigRef:   "main",
		CheckoutRef: "main",
		Disabled:    common.Bool(false),
	}
	triggerCreated, err := triggerService.Create(ctx, newTrigger, projectID, pipelineID)
	assert.Assert(t, err)

	idNewTrigger := triggerCreated.ID
	triggerToUpdate := Trigger{
		EventName: "New event name",
		Disabled:  common.Bool(true),
	}

	triggerUpdated, err := triggerService.Update(ctx, triggerToUpdate, projectID, idNewTrigger)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(triggerUpdated.EventName, "New event name"))

	triggerFetched, err := triggerService.Get(ctx, projectID, idNewTrigger)
	assert.Assert(t, err)
	t.Log(triggerFetched)

	err = triggerService.Delete(ctx, projectID, idNewTrigger)
	assert.Assert(t, err)

	triggerFetched, err = triggerService.Get(ctx, projectID, idNewTrigger)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(triggerFetched))
}

func TestFullScheduleTrigger(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	newTrigger := Trigger{
		EventName:   "Test schedule trigger",
		CheckoutRef: "main",
		ConfigRef:   "main",
		Disabled:    common.Bool(false),
		EventSource: common.EventSource{
			Provider: "schedule",
			Schedule: common.Schedule{
				CronExpression:   "0 1 * * *",
				AttributionActor: "current",
			},
		},
		Parameters: map[string]any{"env": "staging"},
	}

	created, err := triggerService.Create(ctx, newTrigger, knownProjectID, knownSchedulePipelineID)
	assert.Assert(t, err)
	assert.Check(t, created.ID != "")
	assert.Check(t, cmp.Equal(created.EventSource.Provider, "schedule"))

	fetched, err := triggerService.Get(ctx, knownProjectID, created.ID)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(fetched.EventSource.Schedule.CronExpression, "0 1 * * *"))

	_, err = triggerService.Update(ctx, Trigger{
		EventName: "Updated schedule trigger",
		EventSource: common.EventSource{
			Schedule: common.Schedule{CronExpression: "0 2 * * *"},
		},
	}, knownProjectID, created.ID)
	assert.Assert(t, err)

	err = triggerService.Delete(ctx, knownProjectID, created.ID)
	assert.Assert(t, err)

	deleted, err := triggerService.Get(ctx, knownProjectID, created.ID)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(deleted))
}

func TestFullTrigger(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	triggerService := NewTriggerService(c)

	pipelineID := "bee796a0-7ec2-478c-ab87-6a5039d7a216"
	projectID := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	newTrigger := Trigger{
		EventSource: common.EventSource{
			Provider: "github_app",
			Repo: common.Repo{
				ExternalId: "952038793",
			},
		},
		EventPreset: "all-pushes",
		ConfigRef:   "main",
		CheckoutRef: "main",
		Disabled:    common.Bool(false),
	}
	triggerCreated, err := triggerService.Create(ctx, newTrigger, projectID, pipelineID)
	assert.Assert(t, err)

	idNewTrigger := triggerCreated.ID
	triggerToUpdate := Trigger{
		Disabled: common.Bool(true),
	}

	_, err = triggerService.Update(ctx, triggerToUpdate, projectID, idNewTrigger)
	assert.Assert(t, err)

	triggerFetched, err := triggerService.Get(ctx, projectID, idNewTrigger)
	assert.Assert(t, err)

	t.Log(triggerFetched)
	err = triggerService.Delete(ctx, projectID, idNewTrigger)
	assert.Assert(t, err)

	triggerFetched, err = triggerService.Get(ctx, projectID, idNewTrigger)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(triggerFetched))
}
