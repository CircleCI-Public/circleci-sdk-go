package organization_test

import (
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
	"github.com/CircleCI-Public/circleci-sdk-go/organization"
)

const testTok = "2d0a120d-0d44-40ae-906e-5856cc331f76"

func TestOrganizationService_Create(t *testing.T) {
	fc := fakecircle.New(testTok)
	srv := httptest.NewServer(fc)
	t.Cleanup(srv.Close)

	c := client.NewClient(srv.URL+"/api/v2", testTok)
	os := organization.NewOrganizationService(c)

	var org *organization.Organization
	t.Run("create", func(t *testing.T) {
		var err error
		org, err = os.Create("test org name", "github")
		assert.Assert(t, err)
		assert.Check(t, cmp.DeepEqual(org, &organization.Organization{
			Id:      "ignored",
			Name:    "test org name",
			VcsType: "github",
			Slug:    "github/test org name",
		}, cmpopts.IgnoreFields(organization.Organization{}, "Id")))
		assert.Check(t, org.Id != "")
	})

	t.Run("delete", func(t *testing.T) {
		err := os.Delete(org.Id)
		assert.Assert(t, err)
	})
}

func TestOrganizationService_Create_Integration(t *testing.T) {
	c := integrationtest.Client(t)
	orgService := organization.NewOrganizationService(c)

	org, err := orgService.Create("SDK_ORG_TEST2", fakecircle.TypeCircleCI)
	assert.Assert(t, err)
	t.Log(org)

	err = orgService.Delete(org.Id)
	assert.Assert(t, err)
}
