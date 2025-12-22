package envproject_test

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/envproject"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

const testTok = "0c3f30ae-66c3-40c9-9674-db6774f657fb"

func TestEnvService_List(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	envService := envproject.NewEnvService(c)

	o, err := fc.AddOrg(fakecircle.NewOrg{
		Type: fakecircle.TypeCircleCI,
		Name: "test org",
	})
	assert.Assert(t, err)
	orgPrj, err := fc.AddProject(fakecircle.NewProject{
		OrgID: o.ID,
		Name:  "test project",
	})
	assert.Assert(t, err)
	_, err = fc.AddProjectEnv(orgPrj.ID, fakecircle.NewEnvVarProject{
		Name: "FIREBASE_TOKEN",
	})
	assert.Assert(t, err)

	t.Run("list", func(t *testing.T) {
		ctx := context.TODO()
		envs, err := envService.List(ctx, orgPrj.Slug)
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envs, []envproject.EnvVariable{
			{
				Name:      "FIREBASE_TOKEN",
				CreatedAt: time.Now(),
			},
		}, cmpopts.EquateApproxTime(time.Second)))
	})
}

func TestEnvService_List_Integration(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	envService := envproject.NewEnvService(c)

	envs, err := envService.List(ctx, "circleci/8e4z1Akd74woxagxnvLT5q/CzMcAU8dvQo4FJhyj87QsA")
	assert.Assert(t, err)
	t.Log(envs)
}

func TestEnvService_Create(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	envService := envproject.NewEnvService(c)

	o, err := fc.AddOrg(fakecircle.NewOrg{
		Type: fakecircle.TypeCircleCI,
		Name: "test org",
	})
	assert.Assert(t, err)
	orgPrj, err := fc.AddProject(fakecircle.NewProject{
		OrgID: o.ID,
		Name:  "test project",
	})
	assert.Assert(t, err)

	t.Run("empty", func(t *testing.T) {
		ctx := context.TODO()
		envs, err := envService.List(ctx, orgPrj.Slug)
		assert.Assert(t, err)
		assert.Check(t, cmp.Len(envs, 0))
	})

	t.Run("create", func(t *testing.T) {
		ctx := context.TODO()
		envCreated, err := envService.Create(ctx, orgPrj.Slug, "VALUE", "test_sdk")
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envCreated, &envproject.EnvVariable{
			Name:      "test_sdk",
			CreatedAt: time.Now(),
			Value:     "VALUE",
		}, cmpopts.EquateApproxTime(time.Second)))
	})

	// TODO: GET ONE

	t.Run("list", func(t *testing.T) {
		ctx := context.TODO()
		envs, err := envService.List(ctx, orgPrj.Slug)
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(envs, []envproject.EnvVariable{
			{
				Name:      "test_sdk",
				Value:     "VALUE",
				CreatedAt: time.Now(),
			},
		}, cmpopts.EquateApproxTime(time.Second)))
	})

	t.Run("delete", func(t *testing.T) {
		ctx := context.TODO()
		err := envService.Delete(ctx, orgPrj.Slug, "test_sdk")
		assert.Assert(t, err)
	})

	t.Run("empty again", func(t *testing.T) {
		ctx := context.TODO()
		envs, err := envService.List(ctx, orgPrj.Slug)
		assert.Assert(t, err)
		assert.Check(t, cmp.Len(envs, 0))
	})
}

func TestEnvService_Create_Integration(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	envService := envproject.NewEnvService(c)

	projectSlug := "circleci/8e4z1Akd74woxagxnvLT5q/CzMcAU8dvQo4FJhyj87QsA"

	envCreated, err := envService.Create(ctx, projectSlug, "VALUE", "test_sdk")
	assert.Assert(t, err)

	t.Log(envCreated)

	err = envService.Delete(ctx, projectSlug, "test_sdk")
	assert.Assert(t, err)
}
