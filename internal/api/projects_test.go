package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sq-cli/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/projects/create", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, "test-name", r.Form.Get("name"))
		assert.Equal(t, "test-key", r.Form.Get("project"))
		assert.Equal(t, "test-org", r.Form.Get("organization"))
		assert.Equal(t, "public", r.Form.Get("visibility"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateProjectResponse{
			Project: struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			}{Key: "test-key", Name: "test-name"},
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token", Organization: "test-org"}
	client := NewClient(cfg)

	resp, err := client.CreateProject(CreateProjectParams{
		Name:       "test-name",
		ProjectKey: "test-key",
		Visibility: "public",
	})

	assert.NoError(t, err)
	assert.Equal(t, "test-key", resp.Project.Key)
	assert.Equal(t, "test-name", resp.Project.Name)
}

func TestSearchProjects(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/projects/search", r.URL.Path)
		assert.Equal(t, "test-org", r.URL.Query().Get("organization"))
		assert.Equal(t, "filter-term", r.URL.Query().Get("q"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SearchProjectsResponse{
			Components: []Project{
				{Key: "p1", Name: "Project 1", Visibility: "public"},
			},
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token", Organization: "test-org"}
	client := NewClient(cfg)

	resp, err := client.SearchProjects(SearchProjectsParams{Filter: "filter-term"})
	assert.NoError(t, err)
	assert.Len(t, resp.Components, 1)
	assert.Equal(t, "p1", resp.Components[0].Key)
}
