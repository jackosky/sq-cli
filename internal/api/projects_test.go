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
	const (
		testName = "test-name"
		testKey  = "test-key"
		testOrg  = "test-org"
	)

	// Mock Server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/projects/create", r.URL.Path)
		assert.Equal(t, "POST", r.Method)
		
		err := r.ParseForm()
		assert.NoError(t, err)
		assert.Equal(t, testName, r.Form.Get("name"))
		assert.Equal(t, testKey, r.Form.Get("project"))
		assert.Equal(t, testOrg, r.Form.Get("organization"))
		assert.Equal(t, "public", r.Form.Get("visibility"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(CreateProjectResponse{
			Project: struct {
				Key  string `json:"key"`
				Name string `json:"name"`
			}{Key: testKey, Name: testName},
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token", Organization: testOrg}
	client := NewClient(cfg)

	resp, err := client.CreateProject(CreateProjectParams{
		Name:       testName,
		ProjectKey: testKey,
		Visibility: "public",
	})

	assert.NoError(t, err)
	assert.Equal(t, testKey, resp.Project.Key)
	assert.Equal(t, testName, resp.Project.Name)
}

func TestSearchProjects(t *testing.T) {
	const testOrg = "test-org"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/projects/search", r.URL.Path)
		assert.Equal(t, testOrg, r.URL.Query().Get("organization"))
		assert.Equal(t, "filter-term", r.URL.Query().Get("q"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SearchProjectsResponse{
			Components: []Project{
				{Key: "p1", Name: "Project 1", Visibility: "public"},
			},
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "token", Organization: testOrg}
	client := NewClient(cfg)

	resp, err := client.SearchProjects(SearchProjectsParams{Filter: "filter-term"})
	assert.NoError(t, err)
	assert.Len(t, resp.Components, 1)
	assert.Equal(t, "p1", resp.Components[0].Key)
}
