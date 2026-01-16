package cmd

import (
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get resources (issues, prs, etc)",
	Long:  `Get information about resources from SonarQube Cloud.`,
}

func init() {
	rootCmd.AddCommand(getCmd)
}
