package api

import (
	"fmt"
	"sq-cli/internal/config"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty  *resty.Client
	Config *config.Config
}

func NewClient(cfg *config.Config) *Client {
	r := resty.New()
	r.SetBaseURL(cfg.URL)
	r.SetHeader("Authorization", fmt.Sprintf("Bearer %s", cfg.Token))
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded") 

	return &Client{
		resty:  r,
		Config: cfg,
	}
}

// ErrorResponse represents a generic error from SonarQube API
type ErrorResponse struct {
	Errors []struct {
		Msg string `json:"msg"`
	} `json:"errors"`
}

type UserCurrentResponse struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *Client) ValidateToken() (*UserCurrentResponse, error) {
	var result UserCurrentResponse
	err := c.Get("/api/users/current", nil, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) Post(endpoint string, params map[string]string, result interface{}) error {
	req := c.resty.R()
	
	if len(params) > 0 {
		req.SetFormData(params)
	}

	if result != nil {
		req.SetResult(result)
	}

	req.SetError(&ErrorResponse{})

	resp, err := req.Post(endpoint)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		errResp := resp.Error().(*ErrorResponse)
		if len(errResp.Errors) > 0 {
			return fmt.Errorf("api error: %s", errResp.Errors[0].Msg)
		}
		return fmt.Errorf("api error: status %d", resp.StatusCode())
	}

	return nil
}

func (c *Client) Get(endpoint string, params map[string]string, result interface{}) error {
	req := c.resty.R()

	if len(params) > 0 {
		req.SetQueryParams(params)
	}

	if result != nil {
		req.SetResult(result)
	}
	
	req.SetError(&ErrorResponse{})

	resp, err := req.Get(endpoint)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	if resp.IsError() {
		errResp := resp.Error().(*ErrorResponse)
		if len(errResp.Errors) > 0 {
			return fmt.Errorf("api error: %s", errResp.Errors[0].Msg)
		}
		return fmt.Errorf("api error: status %d", resp.StatusCode())
	}

	return nil
}
