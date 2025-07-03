package client_test

import (
	"encoding/json"
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
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		})
		assert.Assert(t, err)
		body := decodeBody(t, res)
		assert.Check(t, cmp.DeepEqual(body, map[string]any{
			"message": "Hello World!",
		}))
	})

	t.Run("unauthed", func(t *testing.T) {
		c := client.NewClient(srv.URL, "")
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		})
		assert.Check(t, cmp.ErrorContains(err, "401 Unauthorized"))
		assert.Check(t, cmp.ErrorContains(err, "You must log in first."))
		assert.Check(t, cmp.Nil(res))
	})

	t.Run("bad_token", func(t *testing.T) {
		c := client.NewClient(srv.URL, "not-valid")
		res, err := c.RequestHelper(http.MethodGet, "/api/test", map[string]any{
			"foo": "bar",
		})
		assert.Check(t, cmp.ErrorContains(err, "401 Unauthorized"))
		assert.Check(t, cmp.ErrorContains(err, "Invalid token provided."))
		assert.Check(t, cmp.Nil(res))
	})
}

func decodeBody(t testing.TB, r *http.Response) map[string]any {
	t.Helper()
	defer func() {
		assert.Check(t, r.Body.Close())
	}()

	m := make(map[string]any)
	err := json.NewDecoder(r.Body).Decode(&m)
	assert.Assert(t, err)
	return m
}
