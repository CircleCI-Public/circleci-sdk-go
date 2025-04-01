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
	trs, err := triggerService.List(pipeline_id)
	if err != nil {
		t.Log(err)
		t.Error("Error getting triggers")
		t.FailNow()
	}
	t.Log(trs)
	// if len(trs) != 1 {
	// 	t.Errorf("Size of triggers is not correct")
	// }
}

func TestFullTrigger(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	triggerService := trigger.NewTriggerService(client)

	pipeline_id := "f51dd4e5-11fe-4069-adad-0df0a7493d53"
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
	trigger_created, err := triggerService.Create(new_trigger, pipeline_id)
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
	trigger_updated, err := triggerService.Update(trigger_to_update, id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error updating trigger")
		t.FailNow()
	}
	if trigger_updated.Description != "Updated description" {
		t.Error("Trigger was not updated")
		t.FailNow()
	}
	trigger_fetched, err := triggerService.Get(id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error getting trigger")
		t.FailNow()
	}
	t.Log(trigger_fetched)
	err = triggerService.Delete(id_new_trigger)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting trigger")
		t.FailNow()
	}
	trigger_fetched, err = triggerService.Get(id_new_trigger)
	t.Log(err)
	if trigger_fetched != nil {
		t.Error("Trigger was not deleted")
		t.FailNow()
	}
}
