package env_test

import (
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/env"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
)

const testTok = "0c3f30ae-66c3-40c9-9674-db6774f657fb"

func TestEnvService_List(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	envService := env.NewEnvService(c)

	o, err := fc.AddOrg(fakecircle.NewOrg{
		Type: "circleci",
		Name: "test org",
	})
	assert.Assert(t, err)
	orgCtx, err := fc.AddContext(fakecircle.NewContext{
		OrgID: o.ID,
		Name:  "test context",
	})
	assert.Assert(t, err)
	_, err = fc.AddContextEnv(orgCtx.ID, fakecircle.NewEnvVar{
		Variable: "FIREBASE_TOKEN",
	})
	assert.Assert(t, err)

	t.Run("list", func(t *testing.T) {
		envs, err := envService.List(orgCtx.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envs, []env.EnvVariable{
			{
				ContextId: orgCtx.ID.String(),
				Variable:  "FIREBASE_TOKEN",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
		}, cmpopts.EquateApproxTime(time.Second)))
	})
}

func TestEnvService_List_Integration(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(c)

	envs, err := envService.List("e51158a2-f59c-4740-9eb4-d20609baa07e")
	assert.Assert(t, err)
	t.Log(envs)
}

func TestEnvService_Create(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	envService := env.NewEnvService(c)

	o, err := fc.AddOrg(fakecircle.NewOrg{
		Type: "circleci",
		Name: "test org",
	})
	assert.Assert(t, err)
	orgCtx, err := fc.AddContext(fakecircle.NewContext{
		OrgID: o.ID,
		Name:  "test context",
	})
	assert.Assert(t, err)

	t.Run("empty", func(t *testing.T) {
		envs, err := envService.List(orgCtx.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.Len(envs, 0))
	})

	t.Run("create", func(t *testing.T) {
		envCreated, err := envService.Create(orgCtx.ID.String(), "VALUE", "test_sdk")
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envCreated, &env.EnvVariable{
			ContextId: orgCtx.ID.String(),
			Variable:  "test_sdk",
			UpdatedAt: time.Now(),
			CreatedAt: time.Now(),
		}, cmpopts.EquateApproxTime(time.Second)))
	})

	t.Run("list", func(t *testing.T) {
		envs, err := envService.List(orgCtx.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envs, []env.EnvVariable{
			{
				ContextId: orgCtx.ID.String(),
				Variable:  "test_sdk",
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			},
		}, cmpopts.EquateApproxTime(time.Second)))
	})

	t.Run("delete", func(t *testing.T) {
		err := envService.Delete(orgCtx.ID.String(), "test_sdk")
		assert.Assert(t, err)
	})

	t.Run("empty again", func(t *testing.T) {
		envs, err := envService.List(orgCtx.ID.String())
		assert.Assert(t, err)
		assert.Check(t, cmp.Len(envs, 0))
	})
}

func TestEnvService_Create_Integration(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(c)

	contextID := "e51158a2-f59c-4740-9eb4-d20609baa07e"

	envCreated, err := envService.Create(contextID, "VALUE", "test_sdk")
	assert.Assert(t, err)

	t.Log(envCreated)

	err = envService.Delete(contextID, "test_sdk")
	assert.Assert(t, err)
}
