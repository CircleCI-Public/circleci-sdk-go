package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CircleCI-Public/circleci-sdk-go/client"
	"github.com/CircleCI-Public/circleci-sdk-go/internal/testing/fakecircle"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

func TestClient_RequestHelper(t *testing.T) {
	const testTok = "CCIPAT_865d543e-9d33-4157-a6cc-8f4416a02df0"

	fs := fakecircle.New(testTok)
	srv := httptest.NewServer(fs)
	t.Cleanup(srv.Close)

	t.Run("authed", func(t *testing.T) {
		c := client.NewClient(srv.URL, testTok)
		body := make(map[string]any)
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		}, &body)
		assert.Assert(t, err)
		assert.Check(t, cmp.Equal(res.StatusCode, http.StatusOK))
		assert.Check(t, cmp.DeepEqual(body, map[string]any{
			"message": "Hello World!",
		}))
	})

	t.Run("unauthed", func(t *testing.T) {
		c := client.NewClient(srv.URL, "")
		body := make(map[string]any)
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		}, &body)
		assert.Check(t, cmp.ErrorContains(err, "401 Unauthorized"))
		assert.Check(t, cmp.ErrorContains(err, "You must log in first."))
		assert.Check(t, cmp.Nil(res))
	})

	t.Run("bad_token", func(t *testing.T) {
		c := client.NewClient(srv.URL, "not-valid")
		body := make(map[string]any)
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		}, &body)
		assert.Check(t, cmp.ErrorContains(err, "401 Unauthorized"))
		assert.Check(t, cmp.ErrorContains(err, "Invalid token provided."))
		assert.Check(t, cmp.Nil(res))
	})
}
