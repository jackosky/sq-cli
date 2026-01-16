package cmd

import (
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create resources (project, etc)",
	Long:  `Create resources in SonarQube Cloud.`,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
