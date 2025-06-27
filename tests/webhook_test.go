package tests
import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/webhook"
)

func TestFullWebhook(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	webhookService :=  webhook.NewWebhookService(client)
	project_id := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	new_webhook :=  webhook.Webhook{
		Name: "Test webhook",
		Url: "https://example.circleci.com/",
		Events: []string{"job-completed",},
		VerifyTls: common.Bool(false),
		SigningSecret: "secret",
		Scope: common.Scope{
			Type: "project",
			Id: project_id,
		},
	}
	webhook_created, err :=  webhookService.Create(new_webhook)
	if err != nil {
		t.Log(err)
		t.Error("Error creating  webhook")
		t.FailNow()
	}
	
	id_new_webhook :=  webhook_created.Id
	webhook_to_update := webhook.Webhook{
		Name: "Webhook updated",
		Events: append(webhook_created.Events, "workflow-completed"),
		VerifyTls: common.Bool(true),
	}
	webhook_updated, err :=  webhookService.Update(webhook_to_update, id_new_webhook)
	if err != nil {
		t.Log(err)
		t.Error("Error updating webhook")
		t.FailNow()
	}
	if webhook_updated.Name != "Webhook updated" {
		t.Error("Webhook was not updated")
		t.FailNow()
	}
	webhook_fetched, err :=  webhookService.Get(id_new_webhook)
	if err != nil {
		t.Log(err)
		t.Error("Error getting webhook")
		t.FailNow()
	}
	t.Log(webhook_fetched)
	err =  webhookService.Delete(id_new_webhook)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting webhook")
		t.FailNow()
	}
	webhook_fetched, err =  webhookService.Get(id_new_webhook)
	t.Log(err)
	if  webhook_fetched != nil {
		t.Error("Webhook was not deleted")
		t.FailNow()
	}
}
