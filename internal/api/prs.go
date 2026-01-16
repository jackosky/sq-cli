package api

type PullRequest struct {
	Key    string `json:"key"` // Use this for ID
	Title  string `json:"title"`
	Branch string `json:"branch"`
	Base   string `json:"base"`
	Status struct {
		QualityGateStatus string `json:"qualityGateStatus"`
	} `json:"status"`
	Url string `json:"url"` // Not always returned, might construct it
}

type ListPullRequestsResponse struct {
	PullRequests []PullRequest `json:"pullRequests"`
}

func (c *Client) ListPullRequests(projectKey string) (*ListPullRequestsResponse, error) {
	queryParams := map[string]string{
		"project": projectKey,
	}

	var result ListPullRequestsResponse
	err := c.Get("/api/project_pull_requests/list", queryParams, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
