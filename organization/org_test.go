package organization

import (
	"os"
	"testing"

	"gotest.tools/v3/skip"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
)

func TestCreateOrg(t *testing.T) {
	token := os.Getenv("CCIPERSONALACCESSTOKEN_ASKSEC_310")
	skip.If(t, token == "", "Token not found")

	c := client.NewClient("https://circleci.com/api/v2", token)
	orgService := NewOrganizationService(c)

	org, err := orgService.Create("SDK_ORG_TEST2", "circleci")
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
