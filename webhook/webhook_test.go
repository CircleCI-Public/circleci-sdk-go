package webhook

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

func TestFullWebhook(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
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
	webhookCreated, err := webhookService.Create(ctx, newWebhook)
	assert.Assert(t, err)

	idNewWebhook := webhookCreated.Id
	webhookToUpdate := Webhook{
		Name:      "Webhook updated",
		Events:    append(webhookCreated.Events, "workflow-completed"),
		VerifyTls: common.Bool(true),
	}
	webhookUpdated, err := webhookService.Update(ctx, webhookToUpdate, idNewWebhook)
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(webhookUpdated.Name, "Webhook updated"))

	webhookFetched, err := webhookService.Get(ctx, idNewWebhook)
	assert.Assert(t, err)
	t.Log(webhookFetched)

	err = webhookService.Delete(ctx, idNewWebhook)
	assert.Assert(t, err)

	webhookFetched, err = webhookService.Get(ctx, idNewWebhook)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(webhookFetched))
}
