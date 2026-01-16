package ui

import "github.com/charmbracelet/lipgloss"

var (
	// Colors
	ColorPrimary   = lipgloss.Color("#2563EB") // Blue
	ColorSuccess   = lipgloss.Color("#10B981") // Green
	ColorError     = lipgloss.Color("#EF4444") // Red
	ColorWarning   = lipgloss.Color("#F59E0B") // Amber
	ColorSecondary = lipgloss.Color("#6B7280") // Gray
	ColorKeyword   = lipgloss.Color("#8B5CF6") // Violet

	// Styles
	TitleStyle = lipgloss.NewStyle().
			Foreground(ColorPrimary).
			Bold(true).
			MarginBottom(1)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)
	
	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	KeywordStyle = lipgloss.NewStyle().
			Foreground(ColorKeyword)

	FaintStyle = lipgloss.NewStyle().
			Foreground(ColorSecondary)

	// Status Helpers
	StatusOpenStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#3B82F6")).Padding(0, 1).Bold(true)
	StatusClosedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#9CA3AF")).Padding(0, 1)
	StatusPassedStyle = lipgloss.NewStyle().Foreground(ColorSuccess).Padding(0, 1).Bold(true)
	StatusFailedStyle = lipgloss.NewStyle().Foreground(ColorError).Padding(0, 1).Bold(true)
)

func Status(status string) string {
	switch status {
	case "OPEN", "CONFIRMED":
		return StatusOpenStyle.Render(status)
	case "CLOSED", "RESOLVED":
		return StatusClosedStyle.Render(status)
	case "OK", "PASSED":
		return StatusPassedStyle.Render(status)
	case "ERROR", "FAILED":
		return StatusFailedStyle.Render(status)
	default:
		return lipgloss.NewStyle().Padding(0, 1).Render(status)
	}
}

func Severity(severity string) string {
	switch severity {
	case "BLOCKER", "CRITICAL":
		return lipgloss.NewStyle().Foreground(ColorError).Bold(true).Render(severity)
	case "MAJOR":
		return lipgloss.NewStyle().Foreground(ColorWarning).Render(severity)
	default:
		return lipgloss.NewStyle().Foreground(ColorPrimary).Render(severity)
	}
}
