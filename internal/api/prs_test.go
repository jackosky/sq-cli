package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sq-cli/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPullRequests(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/project_pull_requests/list", r.URL.Path)
		assert.Equal(t, "proj-key", r.URL.Query().Get("project"))

		w.Header().Set("Content-Type", "application/json")
		response := ListPullRequestsResponse{
			PullRequests: []PullRequest{
				{Key: "123", Title: "Fix bug", Branch: "feature/bug", Base: "main"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token"}
	client := NewClient(cfg)

	resp, err := client.ListPullRequests("proj-key")
	assert.NoError(t, err)
	assert.Len(t, resp.PullRequests, 1)
	assert.Equal(t, "123", resp.PullRequests[0].Key)
}
