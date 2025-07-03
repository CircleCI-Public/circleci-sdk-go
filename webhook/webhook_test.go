package webhook

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

func TestFullWebhook(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	webhookService := NewWebhookService(c)
	projectID := "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33"
	newWebhook := Webhook{
		Name:          "Test webhook",
		Url:           "https://example.circleci.com/",
		Events:        []string{"job-completed"},
		VerifyTls:     common.Bool(false),
		SigningSecret: "secret",
		Scope: common.Scope{
			Type: "project",
			Id:   projectID,
		},
	}
	webhookCreated, err := webhookService.Create(newWebhook)
	assert.Assert(t, err)

	idNewWebhook := webhookCreated.Id
	webhookToUpdate := Webhook{
		Name:      "Webhook updated",
		Events:    append(webhookCreated.Events, "workflow-completed"),
		VerifyTls: common.Bool(true),
	}
	webhookUpdated, err := webhookService.Update(webhookToUpdate, idNewWebhook)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(webhookUpdated.Name, "Webhook updated"))

	webhookFetched, err := webhookService.Get(idNewWebhook)
	assert.Assert(t, err)
	t.Log(webhookFetched)

	err = webhookService.Delete(idNewWebhook)
	assert.Assert(t, err)

	webhookFetched, err = webhookService.Get(idNewWebhook)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(webhookFetched))
}
