package organization_test

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
	"github.com/CircleCI-Public/circleci-sdk-go/organization"
)

func TestOrganizationService_Create_Integration(t *testing.T) {
	c := integrationtest.Client(t)
	orgService := organization.NewOrganizationService(c)

	org, err := orgService.Create("SDK_ORG_TEST2", "circleci")
	assert.Assert(t, err)
	t.Log(org)

	err = orgService.Delete(org.Id)
	assert.Assert(t, err)
}
