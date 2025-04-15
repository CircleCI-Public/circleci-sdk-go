package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/project"
)

func TestGetProject(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	slug := "circleci/8e4z1Akd74woxagxnvLT5q/V29Cenkg8EaiSZARmWm8Lz"
	p, err := projectService.Get(slug)
	if err != nil {
		t.Log(err)
		t.Error("Error getting project")
		t.FailNow()
	}
	t.Log(p)
	if p.Slug != slug {
		t.Errorf("Slug is not correct")
	}
}

func TestFullProject(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	name := "test-api-client-repo"
	organization_id := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"
	p, err := projectService.Create(name, organization_id)
	if err != nil {
		t.Log(err)
		t.Error("Error creating project")
		t.FailNow()
	}
	t.Log(p)
	
	project_settings, err := projectService.GetSettings("circleci", p.OrganizationId, p.Id)
	if err != nil {
		t.Log(err)
		t.Error("Error getting settings from project")
		t.FailNow()
	}
	t.Log(project_settings)
	new_settings := project.ProjectSettings{
		Advanced: project.AdvanceSettings{
			AutocancelBuilds: true,
			DisableSSH: false,
		},
	}
	project_settings, err = projectService.UpdateSettings(new_settings, "circleci", p.OrganizationId, p.Id)
	if err != nil {
		t.Log(err)
		t.Error("Error updating project settings")
		t.FailNow()
	}
	t.Log(project_settings)

	err = projectService.Delete(p.Slug)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting project")
		t.FailNow()
	}
}

// This test is for manual usage only
func _TestClassicProject(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	name := "test-api-client-repo"
	organization_id := "30361eb7-0812-447f-bca5-a299d96576c7"
	_, err := projectService.Create(name, organization_id)
	if err != nil {
		t.Log(err)
		t.Error("Error creating project")
		t.FailNow()
	}
}

// This test is for manual usage only
func _TestDeleteProject(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	err := projectService.Delete("github/marboledacci/test-api-client-repo")
	if err != nil {
		t.Log(err)
		t.Error("Error deleting project")
		t.FailNow()
	}
}
