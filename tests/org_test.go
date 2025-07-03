package tests

import (
	"os"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/organization"
)

func TestCreateOrg(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	if token == "" {
		t.Error("Error: Token not found")
	}
	client := client.NewClient("https://circleci.com/api/v2", token)
	orgService := organization.NewOrganizationService(client)

	org, err := orgService.Create("SDK_ORG_TEST", "circleci")
	if err != nil {
		t.Log(err)
		t.Error("Error creating organization")
		t.FailNow()
	}
	t.Log(org)

	err = orgService.Delete(org.Id)
	if err != nil {
		t.Log(err)
		t.Error("Error deleting organization")
		t.FailNow()
	}
}
