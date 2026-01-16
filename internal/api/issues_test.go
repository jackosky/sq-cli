package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sq-cli/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchIssues(t *testing.T) {
	const testOrg = "test-org"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/issues/search", r.URL.Path)
		
		q := r.URL.Query()
		assert.Equal(t, "proj-key", q.Get("componentKeys"))
		assert.Equal(t, testOrg, q.Get("organization")) // The fix we verified earlier
		assert.Equal(t, "main", q.Get("branch"))
		assert.Equal(t, "BUG", q.Get("types"))
		assert.Equal(t, "CRITICAL", q.Get("severities"))

		w.Header().Set("Content-Type", "application/json")
		response := SearchIssuesResponse{
			Total: 1,
			Issues: []Issue{
				{Key: "issue-1", Message: "Bad bug", Severity: "CRITICAL", Status: "OPEN"},
			},
		}
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token", Organization: testOrg}
	client := NewClient(cfg)

	resp, err := client.SearchIssues(SearchIssuesParams{
		ProjectKey: "proj-key",
		Branch:     "main",
		Type:       "BUG",
		Severities: "CRITICAL",
	})

	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, 1, resp.Total)
		if assert.NotEmpty(t, resp.Issues) {
			assert.Equal(t, "issue-1", resp.Issues[0].Key)
		}
	}
}

func TestSearchIssuesDefaults(t *testing.T) {
	const testOrg = "test-org"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		assert.Empty(t, q.Get("branch")) // Should not be sent if empty
		json.NewEncoder(w).Encode(SearchIssuesResponse{})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Organization: testOrg}
	client := NewClient(cfg)

	_, err := client.SearchIssues(SearchIssuesParams{ProjectKey: "pk"})
	assert.NoError(t, err)
}
