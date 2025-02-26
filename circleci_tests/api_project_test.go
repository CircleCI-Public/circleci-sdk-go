package circleci_tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CircleCI-Public/circleci-sdk-go/circleci"
)

func TestProjectListEnvVars(t *testing.T) {
	circleCiToken := os.Getenv("CIRCLECI_TOKEN")
	circleCiContext, _ := context.WithTimeout(context.Background(), 3*time.Second)

	config := circleci.NewConfiguration()
	config.AddDefaultHeader("Circle-Token", circleCiToken)

	apiClient := circleci.NewAPIClient(config)

	projectSlug := "github/CircleCI-Public/github-cli-orb"
	_, _, err := apiClient.ProjectApi.ListEnvVars(circleCiContext, projectSlug)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
}
