package circleci_tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/CircleCI-Public/circleci-sdk-go/circleci"
)

func TestUser(t *testing.T) {
	circleCiToken := os.Getenv("CIRCLECI_TOKEN")
	circleCiContext, _ := context.WithTimeout(context.Background(), 3*time.Second)

	config := circleci.NewConfiguration()
	config.AddDefaultHeader("Circle-Token", circleCiToken)

	apiClient := circleci.NewAPIClient(config)

	user, response, err := apiClient.UserApi.GetCurrentUser(circleCiContext)
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	if response.StatusCode != 200 {
		t.Logf("%d", response.StatusCode)
		t.FailNow()
	}
	if user.Id != "c06418c1-5f9b-4f54-a300-3fd24096aec4" {
		t.Logf("%+v", user)
		t.FailNow()
	}

}
