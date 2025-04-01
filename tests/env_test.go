package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/env"
)

func TestListEnvs(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(client)

	envs, err := envService.List("e51158a2-f59c-4740-9eb4-d20609baa07e")
	if err != nil {
		t.Log(err)
		t.Error("Error getting env variables")
		t.FailNow()
	}
	t.Log(envs)
}


func TestFullEnv(t *testing.T) {
	token := os.Getenv("CIRCLECI_CLI_TOKEN")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(client)

 	env_created, err := envService.Create("e51158a2-f59c-4740-9eb4-d20609baa07e", "VALUE", "test_sdk")
	if err != nil {
		t.Log(err)
		t.Error("Error creating environment variable")
		t.FailNow()
	}
	t.Log(env_created)
	err = envService.Delete("e51158a2-f59c-4740-9eb4-d20609baa07e", "test_sdk")
	if err != nil {
		t.Log(err)
		t.Error("Error deleting environment variable")
		t.FailNow()
	}
}
