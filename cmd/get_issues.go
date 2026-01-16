package cmd

import (
	"fmt"
	
	"sq-cli/internal/ui"
	"sq-cli/internal/api"

	"github.com/spf13/cobra"
)

var (
	issuesBranch     string
	issuesType       string
	issuesSeverities string
)

var getIssuesCmd = &cobra.Command{
	Use:   "issues [project-key]",
	Short: "Get list of issues for a project",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		projectKey := args[0]
		client := api.NewClient(cfg)

		fmt.Printf("Fetching issues for project: %s (Branch: %s)...\n", projectKey, issuesBranch)
		resp, err := client.SearchIssues(api.SearchIssuesParams{
			ProjectKey: projectKey,
			Branch:     issuesBranch,
			Type:       issuesType,
			Severities: issuesSeverities,
		})
		if err != nil {
			return fmt.Errorf("failed to fetch issues: %w", err)
		}

		if len(resp.Issues) == 0 {
			fmt.Println("No issues found.")
			return nil
		}

		title := fmt.Sprintf("Found %d issues (Total: %d)", len(resp.Issues), resp.Total)
		fmt.Println(ui.TitleStyle.Render(title))
		for _, issue := range resp.Issues {
			severity := ui.Severity(issue.Severity)
			key := ui.KeywordStyle.Render(issue.Key)
			status := ui.Status(issue.Status)
			
			fmt.Printf("[%s] %s - %s (%s)\n", severity, key, issue.Message, status)
		}

		return nil
	},
}

func init() {
	getIssuesCmd.Flags().StringVar(&issuesBranch, "branch", "", "Branch to fetch issues from (default: project default)")
	getIssuesCmd.Flags().StringVar(&issuesType, "type", "", "Filter by issue type (BUG, VULNERABILITY, CODE_SMELL)")
	getIssuesCmd.Flags().StringVar(&issuesSeverities, "severity", "", "Filter by severity (INFO, MINOR, MAJOR, CRITICAL, BLOCKER)")
	
	getCmd.AddCommand(getIssuesCmd)
}
