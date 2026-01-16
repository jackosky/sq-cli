package cmd

import (
	"fmt"
	"os/exec"
	"runtime"

	"sq-cli/internal/api"
	"sq-cli/internal/config"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to SonarQube Cloud",
	RunE: func(cmd *cobra.Command, args []string) error {
		const tokenURL = "https://sonarcloud.io/account/security/"
		
		fmt.Printf("Opening browser to generate token: %s\n", tokenURL)
		openBrowser(tokenURL)

		var token, org string
		
		err := huh.NewInput().
			Title("Paste your User Token").
			Password(true).
			Value(&token).
			Validate(func(s string) error {
				if len(s) == 0 {
					return fmt.Errorf("token cannot be empty")
				}
				return nil
			}).Run()
		if err != nil {
			return err
		}

		err = huh.NewInput().
			Title("Enter your Organization Key").
			Description("You can find this in your organization settings").
			Value(&org).
			Validate(func(s string) error {
				if len(s) == 0 {
					return fmt.Errorf("organization cannot be empty")
				}
				return nil
			}).Run()
		if err != nil {
			return err
		}

		// Validate token
		fmt.Println("Validating token...")
		tempCfg := &config.Config{
			URL: "https://sonarcloud.io",
			Token: token,
			Organization: org,
		}
		
		client := api.NewClient(tempCfg)
		user, err := client.ValidateToken()
		if err != nil {
			return fmt.Errorf("authentication failed: %w", err)
		}

		fmt.Printf("Successfully authenticated as: %s (%s)\n", user.Name, user.Login)

		// Save to config
		err = config.SaveConfig(tempCfg)
		if err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}

		fmt.Println("Configuration saved to ~/.sq-cli.yaml")
		return nil
	},
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Printf("Unable to open browser: %v. Please open %s manually.\n", err, url)
	}
}

func init() {
	authCmd.AddCommand(authLoginCmd)
}
