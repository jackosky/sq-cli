package cmd

import (
	"fmt"
	
	"sq-cli/internal/ui"
	"sq-cli/internal/api"

	"github.com/spf13/cobra"
)

var getProjectsFilter string

var getProjectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Get list of projects for organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient(cfg)

		fmt.Println("Fetching projects...")
		resp, err := client.SearchProjects(api.SearchProjectsParams{
			Filter: getProjectsFilter,
		})
		if err != nil {
			return fmt.Errorf("failed to fetch projects: %w", err)
		}

		if len(resp.Components) == 0 {
			fmt.Println("No projects found.")
			return nil
		}

		title := fmt.Sprintf("Found %d projects (Total: %d)", len(resp.Components), resp.Paging.Total)
		fmt.Println(ui.TitleStyle.Render(title))
		for _, proj := range resp.Components {
			lastAnalysis := proj.LastAnalysisDate
			if lastAnalysis == "" {
				lastAnalysis = "Never"
			}
			
			key := ui.KeywordStyle.Render(proj.Key)
			name := proj.Name
			visibility := ui.FaintStyle.Render(fmt.Sprintf("Visibility: %s", proj.Visibility))
			analysis := ui.FaintStyle.Render(fmt.Sprintf("Last Analysis: %s", lastAnalysis))

			fmt.Printf("[%s] %s (%s, %s)\n", key, name, visibility, analysis)
		}

		return nil
	},
}

func init() {
	getProjectsCmd.Flags().StringVarP(&getProjectsFilter, "filter", "f", "", "Filter projects by name or key")
	getCmd.AddCommand(getProjectsCmd)
}
