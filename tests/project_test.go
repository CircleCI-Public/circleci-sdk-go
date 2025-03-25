package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/project"
)

func TestGetProject(t *testing.T) {
	token := os.Getenv("LOCAL_CCI_TOKEN")
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

func _TestGetProjectSettings(t *testing.T) {
	token := os.Getenv("LOCAL_CCI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	projectService := project.NewProjectService(client)

	p, err := projectService.GetSettings("circleci", "8e4z1Akd74woxagxnvLT5q", "V29Cenkg8EaiSZARmWm8Lz")
	if err != nil {
		t.Log(err)
		t.Error("Error getting project settings")
		t.FailNow()
	}
	t.Log(p)
}
