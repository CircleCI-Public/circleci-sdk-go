package group

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

func TestFullGroup(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	Service := NewService(c)

	orgID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"

	newGroup := Group{
		Name:        "test-group",
		Description: "Test group from SDK",
	}
	groupCreated, err := Service.Create(ctx, newGroup, orgID)
	assert.Assert(t, err)

	idNewGroup := groupCreated.ID

	groupFetched, err := Service.Get(ctx, orgID, idNewGroup)
	assert.Assert(t, err)
	assert.Check(t, groupFetched != nil)

	err = Service.Delete(ctx, orgID, idNewGroup)
	assert.Assert(t, err)

	groupFetched, err = Service.Get(ctx, orgID, idNewGroup)
	assert.Assert(t, err != nil)
	assert.Check(t, cmp.Nil(groupFetched))
}

func TestListGroup(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	Service := NewService(c)

	orgID := "3ddcf1d1-7f5f-4139-8cef-71ad0921a968"
	groups, err := Service.List(ctx, orgID)
	assert.Assert(t, err)
	assert.Assert(t, cmp.Len(groups, 1))
}
