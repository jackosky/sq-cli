package cmd

import (
	"sq-cli/internal/config"

	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
	Long:  `Manage your SonarQube Cloud authentication.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		// We still want to load config to get defaults or existing values, 
		// but we don't enforce constraints
		cfg, err = config.LoadConfig()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
