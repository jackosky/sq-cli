package cmd

import (
	"fmt"
	"strings"

	"sq-cli/internal/api"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var createProjectCmd = &cobra.Command{
	Use:   "project",
	Short: "Create a new project",
	RunE: func(cmd *cobra.Command, args []string) error {
		var name, key, visibility string
		visibility = "public"

		// Prompt for Name
		err := huh.NewInput().
			Title("Enter Project Name").
			Value(&name).
			Validate(func(s string) error {
				if len(s) == 0 {
					return fmt.Errorf("project name cannot be empty")
				}
				return nil
			}).Run()
		if err != nil {
			return err
		}

		// Auto-generate key from name if possible
		defaultKey := strings.ReplaceAll(strings.ToLower(name), " ", "-")
		
		err = huh.NewInput().
			Title("Enter Project Key").
			Value(&key).
			Placeholder(defaultKey).
			Validate(func(s string) error {
				if len(s) == 0 {
					return fmt.Errorf("project key cannot be empty")
				}
				return nil
			}).Run()
		if err != nil {
			return err
		}

		err = huh.NewSelect[string]().
			Title("Visibility").
			Options(
				huh.NewOption("Public", "public"),
				huh.NewOption("Private", "private"),
			).
			Value(&visibility).
			Run()
		if err != nil {
			return err
		}

		client := api.NewClient(cfg)
		fmt.Println("Creating project...")
		resp, err := client.CreateProject(api.CreateProjectParams{
			Name:       name,
			ProjectKey: key,
			Visibility: visibility,
		})
		if err != nil {
			return fmt.Errorf("failed to create project: %w", err)
		}

		fmt.Printf("Project '%s' created successfully! Key: %s\n", resp.Project.Name, resp.Project.Key)
		return nil
	},
}

func init() {
	createCmd.AddCommand(createProjectCmd)
}
