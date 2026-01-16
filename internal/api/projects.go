package api

type CreateProjectParams struct {
	Name       string
	ProjectKey string
	Visibility string // public or private
}

type CreateProjectResponse struct {
	Project struct {
		Key  string `json:"key"`
		Name string `json:"name"`
	} `json:"project"`
}

func (c *Client) CreateProject(params CreateProjectParams) (*CreateProjectResponse, error) {
	formData := map[string]string{
		"name":         params.Name,
		"project":      params.ProjectKey,
		"organization": c.Config.Organization,
		"visibility":   params.Visibility,
	}

	var result CreateProjectResponse
	err := c.Post("/api/projects/create", formData, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type Project struct {
	Key              string `json:"key"`
	Name             string `json:"name"`
	Visibility       string `json:"visibility"`
	LastAnalysisDate string `json:"lastAnalysisDate"`
	Revision         string `json:"revision"`
}

type SearchProjectsResponse struct {
	Paging struct {
		PageIndex int `json:"pageIndex"`
		PageSize  int `json:"pageSize"`
		Total     int `json:"total"`
	} `json:"paging"`
	Components []Project `json:"components"`
}

type SearchProjectsParams struct {
	Filter string // Optional filter string
}

func (c *Client) SearchProjects(params SearchProjectsParams) (*SearchProjectsResponse, error) {
	queryParams := map[string]string{
		"organization": c.Config.Organization,
		"ps": "50", // Page size
	}
	
	if params.Filter != "" {
		queryParams["q"] = params.Filter
	}

	var result SearchProjectsResponse
	err := c.Get("/api/projects/search", queryParams, &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
