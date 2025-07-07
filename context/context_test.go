package context_test

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	sdkcontext "github.com/CircleCI-Public/circleci-sdk-go/context"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
)

const testTok = "9708df71-aced-497e-b9d0-f12837c72492"

func TestContextService_List(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	contextService := sdkcontext.NewContextService(c)

	var o fakecircle.Org
	var orgCtx fakecircle.Context
	t.Run("add_org", func(t *testing.T) {
		var err error
		o, err = fc.AddOrg(fakecircle.NewOrg{
			Type: "circleci",
			Name: "8e4z1Akd74woxagxnvLT5q",
		})
		assert.Assert(t, err)
		orgCtx, err = fc.AddContext(fakecircle.NewContext{
			OrgID: o.ID,
			Name:  "test context",
		})
		assert.Assert(t, err)
	})

	t.Run("add_other_org", func(t *testing.T) {
		o2, err := fc.AddOrg(fakecircle.NewOrg{
			Type: "circleci",
			Name: "other",
		})
		assert.Assert(t, err)
		_, err = fc.AddContext(fakecircle.NewContext{
			OrgID: o2.ID,
			Name:  "other test context",
		})
		assert.Assert(t, err)
	})

	t.Run("list", func(t *testing.T) {
		ctxs, err := contextService.List("circleci/8e4z1Akd74woxagxnvLT5q")
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(ctxs, []sdkcontext.Context{
			{
				ID:        orgCtx.ID.String(),
				Name:      "test context",
				CreatedAt: "ignored",
			},
		}, cmpopts.IgnoreFields(sdkcontext.Context{}, "CreatedAt")))
	})
}

func TestContextService_List_Integration(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := sdkcontext.NewContextService(c)

	ctxs, err := contextService.List("circleci/8e4z1Akd74woxagxnvLT5q")
	assert.Assert(t, err)
	t.Log(ctxs)
	assert.Check(t, len(ctxs) > 0)
}

func TestContextService_Get(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	contextService := sdkcontext.NewContextService(c)

	var o fakecircle.Org
	var orgCtx fakecircle.Context
	t.Run("add_org", func(t *testing.T) {
		var err error
		o, err = fc.AddOrg(fakecircle.NewOrg{
			Type: "circleci",
			Name: "8e4z1Akd74woxagxnvLT5q",
		})
		assert.Assert(t, err)
		orgCtx, err = fc.AddContext(fakecircle.NewContext{
			OrgID: o.ID,
			Name:  "test context",
		})
		assert.Assert(t, err)
	})

	t.Run("get", func(t *testing.T) {
		r, err := contextService.Get(orgCtx.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(r, &sdkcontext.Context{
			ID:        orgCtx.ID.String(),
			Name:      "test context",
			CreatedAt: "ignored",
		}, cmpopts.IgnoreFields(sdkcontext.Context{}, "CreatedAt")))
	})
}

func TestContextService_Get_Integration(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Skip("Token not found")
	}
	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := sdkcontext.NewContextService(c)

	ctx, err := contextService.Get("e51158a2-f59c-4740-9eb4-d20609baa07e")
	assert.Assert(t, err)
	assert.Check(t, cmp.Equal(ctx.Name, "Static Context"))
}

func TestContextService_Full(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	contextService := sdkcontext.NewContextService(c)

	var o fakecircle.Org
	t.Run("add_org", func(t *testing.T) {
		var err error
		o, err = fc.AddOrg(fakecircle.NewOrg{
			Type: "circleci",
			Name: "8e4z1Akd74woxagxnvLT5q",
		})
		assert.Assert(t, err)
	})

	var ctxCreated *sdkcontext.Context
	assert.Assert(t, t.Run("create", func(t *testing.T) {
		var err error
		ctxCreated, err = contextService.Create(o.ID.String(), "Test ctx")
		assert.Assert(t, err)
	}))

	t.Run("get", func(t *testing.T) {
		orgCtx, err := contextService.Get(ctxCreated.ID)
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(orgCtx, &sdkcontext.Context{
			ID:        "ignored",
			Name:      "Test ctx",
			CreatedAt: "ignored",
		}, cmpopts.IgnoreFields(sdkcontext.Context{}, "ID", "CreatedAt")))
		assert.Check(t, orgCtx.ID != "")
		assert.Check(t, orgCtx.CreatedAt != "")
	})

	t.Run("delete", func(t *testing.T) {
		err := contextService.Delete(ctxCreated.ID)
		assert.Assert(t, err)
	})

	t.Run("get", func(t *testing.T) {
		ctxFetched, err := contextService.Get(ctxCreated.ID)
		assert.Assert(t, cmp.ErrorContains(err, "context not found"))
		assert.Check(t, cmp.Nil(ctxFetched))
	})
}

func TestContextService_Full_Integration(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := sdkcontext.NewContextService(c)

	organizationID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"

	var ctxCreated *sdkcontext.Context
	t.Run("create", func(t *testing.T) {
		var err error
		ctxCreated, err = contextService.Create(organizationID, "Test ctx")
		assert.Assert(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		err := contextService.Delete(ctxCreated.ID)
		assert.Assert(t, err)
	})

	t.Run("get", func(t *testing.T) {
		ctxFetched, err := contextService.Get(ctxCreated.ID)
		assert.Assert(t, err != nil)
		assert.Assert(t, cmp.Nil(ctxFetched), "Context was not deleted")
	})
}

func TestListRestrictions(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	contextService := sdkcontext.NewContextService(c)

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
	contextService := sdkcontext.NewContextService(c)

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
