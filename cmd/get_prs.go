package cmd

import (
	"fmt"
	
	"sq-cli/internal/ui"
	"sq-cli/internal/api"

	"github.com/spf13/cobra"
)

var getPRsCmd = &cobra.Command{
	Use:   "prs [project-key]",
	Short: "Get list of pull requests for a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectKey := args[0]
		client := api.NewClient(cfg)

		fmt.Printf("Fetching Pull Requests for project: %s...\n", projectKey)
		resp, err := client.ListPullRequests(projectKey)
		if err != nil {
			return fmt.Errorf("failed to fetch PRs: %w", err)
		}

		if len(resp.PullRequests) == 0 {
			fmt.Println("No Pull Requests found.")
			return nil
		}

		title := fmt.Sprintf("Found %d Pull Requests", len(resp.PullRequests))
		fmt.Println(ui.TitleStyle.Render(title))
		for _, pr := range resp.PullRequests {
			key := ui.KeywordStyle.Render(pr.Key)
			status := ui.Status(pr.Status.QualityGateStatus)
			target := ui.FaintStyle.Render(fmt.Sprintf("Target: %s", pr.Base))

			fmt.Printf("[%s] %s (%s) - Quality Gate: %s\n", key, pr.Title, target, status)
		}

		return nil
	},
}

func init() {
	getCmd.AddCommand(getPRsCmd)
}
