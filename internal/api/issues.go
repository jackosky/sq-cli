package api



type Issue struct {
	Key      string `json:"key"`
	Rule     string `json:"rule"`
	Severity string `json:"severity"`
	Status   string `json:"status"`
	Message  string `json:"message"`
	Type     string `json:"type"`
}

type SearchIssuesResponse struct {
	Total  int     `json:"total"`
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	Issues []Issue `json:"issues"`
}

type SearchIssuesParams struct {
	ProjectKey string
	Branch     string
	Type       string // optional
	Severities string // optional
}

func (c *Client) SearchIssues(params SearchIssuesParams) (*SearchIssuesResponse, error) {
	queryParams := map[string]string{
		"componentKeys": params.ProjectKey,
		"organization":  c.Config.Organization,
	}

	if params.Branch != "" {
		queryParams["branch"] = params.Branch
	} else {
		// If branch is not supported or not passed (e.g. community edition doesn't support branch), 
		// but for SonarCloud it's relevant. Though some APIs might behave differently.
		// For now we assume branch is optional.
	}
	
	if params.Type != "" {
		queryParams["types"] = params.Type
	}
	if params.Severities != "" {
		queryParams["severities"] = params.Severities
	}

	// Limit to reasonable default or first page
	queryParams["ps"] = "50" 

	var result SearchIssuesResponse
	err := c.Get("/api/issues/search", queryParams, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
