package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/trigger"
)

func TestListTrigger(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	triggerService := trigger.NewTriggerService(client)

	pipeline_id := "f51dd4e5-11fe-4069-adad-0df0a7493d53"
	project_id := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	trs, err := triggerService.List(project_id, pipeline_id)
	if err != nil {
		t.Log(err)
		t.Error("Error getting triggers")
		t.FailNow()
	}
	t.Log(trs)
}

func TestFullTrigger(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	triggerService := trigger.NewTriggerService(client)

	pipeline_id := "f51dd4e5-11fe-4069-adad-0df0a7493d53"
	project_id := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	new_trigger := trigger.Trigger{
		Name: "test-trigger",
		Description: "Test trigger from SDK",
		EventSource: common.EventSource{
			Provider: "github_app",
			Repo: common.Repo{
				ExternalId: "952038793",
			},
		},
		EventPreset: "all-pushes",
		ConfigRef: "main",
		CheckoutRef: "main",
	}
	trigger_created, err := triggerService.Create(new_trigger, project_id, pipeline_id)
	if err != nil {
		t.Log(err)
		t.Error("Error creating trigger")
		t.FailNow()
	}
	if err != nil {
		t.Log(err)
		t.Error("Error creating trigger")
		t.FailNow()
	}
	id_new_trigger := trigger_created.ID
	trigger_to_update := trigger.Trigger{
		Name: "trigger-updated",
		Description: "Updated description",
	}
	trigger_updated, err := triggerService.Update(trigger_to_update, project_id, id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error updating trigger")
		t.FailNow()
	}
	if trigger_updated.Description != "Updated description" {
		t.Error("Trigger was not updated")
		t.FailNow()
	}
	trigger_fetched, err := triggerService.Get(project_id, id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error getting trigger")
		t.FailNow()
	}
	t.Log(trigger_fetched)
	err = triggerService.Delete(project_id,id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting trigger")
		t.FailNow()
	}
	trigger_fetched, err = triggerService.Get(project_id,id_new_trigger)
	t.Log(err)
	if trigger_fetched != nil {
		t.Error("Trigger was not deleted")
		t.FailNow()
	}
}
