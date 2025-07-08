package integrationtest

import (
	"os"
	"testing"

	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

func Client(t testing.TB) *client.Client {
	t.Helper()

	token := os.Getenv("CIRCLE_TOKEN")
	skip.If(t, token == "", "Token not found")

	return client.NewClient("https://circleci.com/api/v2", token)
}
