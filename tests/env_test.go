package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/env"
)

func TestListEnvs(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(client)

	envs, err := envService.List("70f5c82b-a7e6-464a-af0a-ba857f9d4714")
	if err != nil {
		t.Log(err)
		t.Error("Error getting env variables")
		t.FailNow()
	}
	t.Log(envs)
}


func TestFullEnv(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	envService := env.NewEnvService(client)

 	env_created, err := envService.Create("70f5c82b-a7e6-464a-af0a-ba857f9d4714", "VALUE", "test_sdk")
	if err != nil {
		t.Log(err)
		t.Error("Error creating environment variable")
		t.FailNow()
	}
	t.Log(env_created)
	err = envService.Delete("70f5c82b-a7e6-464a-af0a-ba857f9d4714", "test_sdk")
	if err != nil {
		t.Log(err)
		t.Error("Error deleting environment variable")
		t.FailNow()
	}
}
