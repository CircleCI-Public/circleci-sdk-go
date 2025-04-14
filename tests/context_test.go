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

	organization_slug := "github/CircleCITestOrg"
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

	ctx, err := contextService.Get("70f5c82b-a7e6-464a-af0a-ba857f9d4714")
	if err != nil {
		t.Log(err)
		t.Error("Error fetching context")
		t.FailNow()
	}
	t.Log(ctx)
}

func TestFullContext(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	contextService := context.NewContextService(client)

	organization_id := "ec6887ec-7d44-4b31-b468-7e552408ee32"
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

	restrictions, err := contextService.GetRestrictions("70f5c82b-a7e6-464a-af0a-ba857f9d4714")

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

	context_id := "70f5c82b-a7e6-464a-af0a-ba857f9d4714"
	restriction, err := contextService.CreateRestriction(context_id, "eb0da417-4dfc-4d21-8265-490cd658ae40", "project")
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
