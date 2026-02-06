package runner

import (
	"context"
	"testing"

	"gotest.tools/v3/assert"

	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/integrationtest"
)

func TestFullRunner(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	service := NewService(c)

	// Create a resource class
	createReq := CreateResourceClassRequest{
		ResourceClass: "test-org/test-resource-class",
		Description:   "Test resource class from SDK",
	}
	resourceClass, err := service.CreateResourceClass(ctx, createReq)
	assert.Assert(t, err)
	assert.Check(t, resourceClass != nil)
	assert.Check(t, resourceClass.Id != "")
	assert.Check(t, resourceClass.ResourceClass == createReq.ResourceClass)

	resourceClassID := resourceClass.Id

	// List resource classes
	resourceClasses, err := service.ListResourceClasses(ctx, "test-org", "")
	assert.Assert(t, err)
	assert.Check(t, len(resourceClasses) > 0)

	// Create a token
	createTokenReq := CreateTokenRequest{
		ResourceClass: createReq.ResourceClass,
		Nickname:      "test-token",
	}
	token, err := service.CreateToken(ctx, createTokenReq)
	assert.Assert(t, err)
	assert.Check(t, token != nil)
	assert.Check(t, token.Id != "")
	assert.Check(t, token.Token != "") // Token value only returned on create
	assert.Check(t, token.Nickname == createTokenReq.Nickname)

	tokenID := token.Id

	// List tokens
	tokens, err := service.ListTokens(ctx, createReq.ResourceClass)
	assert.Assert(t, err)
	assert.Check(t, len(tokens) > 0)

	// Delete token
	err = service.DeleteToken(ctx, tokenID)
	assert.Assert(t, err)

	// Try to delete resource class without force (should work now that token is deleted)
	err = service.DeleteResourceClass(ctx, resourceClassID, false)
	assert.Assert(t, err)
}

func TestDeleteResourceClassForce(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	service := NewService(c)

	// Create a resource class
	createReq := CreateResourceClassRequest{
		ResourceClass: "test-org/test-force-delete",
		Description:   "Test force delete",
	}
	resourceClass, err := service.CreateResourceClass(ctx, createReq)
	assert.Assert(t, err)
	assert.Check(t, resourceClass != nil)

	resourceClassID := resourceClass.Id

	// Create a token
	createTokenReq := CreateTokenRequest{
		ResourceClass: createReq.ResourceClass,
		Nickname:      "test-token-force",
	}
	token, err := service.CreateToken(ctx, createTokenReq)
	assert.Assert(t, err)
	assert.Check(t, token != nil)

	// Delete resource class with force (should succeed even with token)
	err = service.DeleteResourceClass(ctx, resourceClassID, true)
	assert.Assert(t, err)
}

func TestListRunners(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	service := NewService(c)

	// Create a resource class first
	createReq := CreateResourceClassRequest{
		ResourceClass: "test-org/test-runners-list",
		Description:   "Test runners list",
	}
	resourceClass, err := service.CreateResourceClass(ctx, createReq)
	assert.Assert(t, err)

	// List runners with resource class filter
	params := ListRunnersParams{
		ResourceClass: createReq.ResourceClass,
	}
	runners, err := service.ListRunners(ctx, params)
	assert.Assert(t, err)
	assert.Check(t, runners != nil) // May be empty if no runners connected

	// List runners with namespace filter
	params = ListRunnersParams{
		Namespace: "test-org",
	}
	runners, err = service.ListRunners(ctx, params)
	assert.Assert(t, err)
	assert.Check(t, runners != nil)

	// Clean up
	err = service.DeleteResourceClass(ctx, resourceClass.Id, true)
	assert.Assert(t, err)
}

func TestTaskCounts(t *testing.T) {
	ctx := context.TODO()
	c := integrationtest.Client(t)
	service := NewService(c)

	// Create a resource class
	createReq := CreateResourceClassRequest{
		ResourceClass: "test-org/test-task-counts",
		Description:   "Test task counts",
	}
	resourceClass, err := service.CreateResourceClass(ctx, createReq)
	assert.Assert(t, err)

	// Get unclaimed task count
	unclaimedCount, err := service.GetUnclaimedTaskCount(ctx, createReq.ResourceClass)
	assert.Assert(t, err)
	assert.Check(t, unclaimedCount != nil)
	assert.Check(t, unclaimedCount.UnclaimedTaskCount >= 0)

	// Get running task count
	runningCount, err := service.GetRunningTaskCount(ctx, createReq.ResourceClass)
	assert.Assert(t, err)
	assert.Check(t, runningCount != nil)
	assert.Check(t, runningCount.RunningRunnerTasks >= 0)

	// Clean up
	err = service.DeleteResourceClass(ctx, resourceClass.Id, true)
	assert.Assert(t, err)
}
