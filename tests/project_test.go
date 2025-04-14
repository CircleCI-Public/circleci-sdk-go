package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/project"
)

func TestGetProject(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	slug := "github/CircleCITestOrg/test-project-terraform-provider"
	p, err := projectService.Get(slug)
	if err != nil {
		t.Log(err)
		t.Error("Error getting project")
		t.FailNow()
	}
	t.Log(p)
}

func TestGetProjectSettings(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	p, err := projectService.GetSettings("github", "CircleCITestOrg", "test-project-terraform-provider")
	if err != nil {
		t.Log(err)
		t.Error("Error getting project settings")
		t.FailNow()
	}
	t.Log(p)
}
