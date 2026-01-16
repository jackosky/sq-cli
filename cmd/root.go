package cmd

import (
	"fmt"
	"os"

	"sq-cli/internal/config"

	"github.com/spf13/cobra"
)

var (
	cfg *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "sq",
	Short: "A simple CLI for SonarQube Cloud",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		cfg, err = config.LoadConfig("")
		if err != nil {
			return err
		}
		if cfg.Token == "" {
			return fmt.Errorf("SONAR_TOKEN is required. Please set it in environment variable or config file")
		}
		if cfg.Organization == "" {
			return fmt.Errorf("SONAR_ORGANIZATION is required. Please set it in environment variable or config file")
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
