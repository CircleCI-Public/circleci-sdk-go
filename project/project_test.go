package project

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/common"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
	"gotest.tools/v3/skip"
)

func TestGetProject(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	projectService := NewProjectService(c)

	slug := "circleci/8e4z1Akd74woxagxnvLT5q/V29Cenkg8EaiSZARmWm8Lz"
	p, err := projectService.Get(slug)
	assert.Assert(t, err)

	t.Log(p)
	assert.Check(t, cmp.Equal(p.Slug, slug))
}

func TestFullProject(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	projectService := NewProjectService(c)

	name := "test-api-client-repo"
	organizationID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"
	p, err := projectService.Create(name, organizationID)
	if err != nil {
		t.Log(err)
		t.Error("Error creating project")
		t.FailNow()
	}
	t.Log(p)

	projectSettings, err := projectService.GetSettings("circleci", p.OrganizationId, p.Id)
	assert.Assert(t, err)
	t.Log(projectSettings)
	newSettings := ProjectSettings{
		Advanced: AdvanceSettings{
			AutocancelBuilds: common.Bool(true),
			DisableSSH:       common.Bool(false),
		},
	}
	projectSettings, err = projectService.UpdateSettings(newSettings, "circleci", p.OrganizationId, p.Id)
	assert.Assert(t, err)
	t.Log(projectSettings)

	err = projectService.Delete(p.Slug)
	assert.Assert(t, err)
}

func TestClassicProject(t *testing.T) {
	t.Skip("This test is for manual usage only")

	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	projectService := NewProjectService(c)

	name := "test-api-client-repo"
	organizationID := "30361eb7-0812-447f-bca5-a299d96576c7"
	_, err := projectService.Create(name, organizationID)
	assert.Assert(t, err)
}

func TestDeleteProject(t *testing.T) {
	t.Skip("This test is for manual usage only")

	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	projectService := NewProjectService(c)

	err := projectService.Delete("github/marboledacci/test-api-client-repo")
	assert.Assert(t, err)
}
