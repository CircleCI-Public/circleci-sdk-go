package context

import (
	"os"
	"testing"

	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

func TestListContexts(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := NewContextService(c)

	organizationSlug := "circleci/8e4z1Akd74woxagxnvLT5q"
	ctxs, err := contextService.List(organizationSlug)
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
		t.Skip("Token not found")
	}
	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := NewContextService(c)

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
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := NewContextService(c)

	organizationID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"
	t.Log("Creating...")
	ctxCreated, err := contextService.Create(organizationID, "Test ctx")
	if err != nil {
		t.Log(err)
		t.Error("Error creating context")
		t.FailNow()
	}
	idNewCtx := ctxCreated.ID
	t.Log("Deleting...")
	err = contextService.Delete(idNewCtx)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting context")
		t.FailNow()
	}
	t.Log("Validating...")
	ctxFetched, err := contextService.Get(idNewCtx)
	t.Log(err)
	if ctxFetched != nil {
		t.Log(err)
		t.Error("Context was not deleted")
		t.FailNow()
	}
}

func TestListRestrictions(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := NewContextService(c)

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
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := NewContextService(c)

	contextID := "e51158a2-f59c-4740-9eb4-d20609baa07e"
	restriction, err := contextService.CreateRestriction(contextID, "e2e8ae23-57dc-4e95-bc67-633fdeb4ac33", "project")
	if err != nil {
		t.Log(err)
		t.Error("Error creating context restriction")
		t.FailNow()
	}
	idNewRestriction := restriction.ID
	t.Log(restriction)
	err = contextService.DeleteRestriction(contextID, idNewRestriction)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting restriction")
		t.FailNow()
	}
}
