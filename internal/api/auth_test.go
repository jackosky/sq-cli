package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sq-cli/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/users/current", r.URL.Path)
		assert.Equal(t, "Bearer valid-token", r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(UserCurrentResponse{
			Login: "user-login",
			Name:  "User Name",
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "valid-token"}
	client := NewClient(cfg)

	user, err := client.ValidateToken()
	assert.NoError(t, err)
	assert.Equal(t, "user-login", user.Login)
}

func TestValidateTokenError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{
			Errors: []struct {
				Msg string `json:"msg"`
			}{{Msg: "Authentication failed"}},
		})
	}))
	defer server.Close()

	cfg := &config.Config{URL: server.URL, Token: "bad-token"}
	client := NewClient(cfg)

	_, err := client.ValidateToken()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Authentication failed")
}
