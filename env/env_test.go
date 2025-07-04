package env

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

func TestListEnvs(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	envService := NewEnvService(c)

	envs, err := envService.List("e51158a2-f59c-4740-9eb4-d20609baa07e")
	assert.Assert(t, err)
	t.Log(envs)
}

func TestFullEnv(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	envService := NewEnvService(c)

	envCreated, err := envService.Create("e51158a2-f59c-4740-9eb4-d20609baa07e", "VALUE", "test_sdk")
	assert.Assert(t, err)

	t.Log(envCreated)

	err = envService.Delete("e51158a2-f59c-4740-9eb4-d20609baa07e", "test_sdk")
	assert.Assert(t, err)
}
