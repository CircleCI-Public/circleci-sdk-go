package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/context"
)

func TestListContexts(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	organization_slug := "circleci/8e4z1Akd74woxagxnvLT5q"
	ctxs, err := contextService.List(organization_slug)
	if err != nil {
		t.Log(err)
		t.Error("Error getting contexts")
		t.FailNow()
	}
	t.Log(ctxs)
}

func TestGetContext(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	ctx, err := contextService.Get("e51158a2-f59c-4740-9eb4-d20609baa07e")
	if err != nil {
		t.Log(err)
		t.Error("Error fetching context")
		t.FailNow()
	}
	t.Log(ctx)
	if ctx.Name != "Static Context" {
		t.Log(err)
		t.Error("Wrong context fetched")
		t.FailNow()
	}
}

func TestFullContext(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	organization_id := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"
	t.Log("Creating...")
	ctx_created, err := contextService.Create(organization_id, "Test ctx")
	if err != nil {
		t.Log(err)
		t.Error("Error creating context")
		t.FailNow()
	}
	id_new_ctx := ctx_created.ID
	t.Log("Deleting...")
	err = contextService.Delete(id_new_ctx)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting context")
		t.FailNow()
	}
	t.Log("Validating...")
	ctx_fetched, err := contextService.Get(id_new_ctx)
	t.Log(err)
	if ctx_fetched != nil {
		t.Log(err)
		t.Error("Context was not deleted")
		t.FailNow()
	}
}

func TestListRestrictions(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	restrictions, err := contextService.GetRestrictions("e51158a2-f59c-4740-9eb4-d20609baa07e")

	if err != nil {
		t.Log(err)
		t.Error("Error getting restrictions")
		t.FailNow()
	}
	t.Log(restrictions)
}

func TestFullRestrictions(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	context_id := "e51158a2-f59c-4740-9eb4-d20609baa07e"
	restriction, err := contextService.CreateRestriction(context_id, "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33", "project")
	if err != nil {
		t.Log(err)
		t.Error("Error creating context restriction")
		t.FailNow()
	}
	id_new_restriction := restriction.ID
	t.Log(restriction)
	err = contextService.DeleteRestriction(context_id, id_new_restriction)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting restriction")
		t.FailNow()
	}
}
