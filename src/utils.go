package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func PrintStyled(s string, style lipgloss.Style) {
	fmt.Println(style.Render(s))
}

func FormattedDiff(diff time.Duration) string {
	// Format the duration into a human-readable string
	var formattedDiff string
	hours := int(diff.Hours())          // Get whole hours
	minutes := int(diff.Minutes()) % 60 // Get remainder of minutes

	// Build the string based on hours and minutes
	if hours > 0 {
		formattedDiff = fmt.Sprintf("%dh", hours)
	}
	if minutes > 0 {
		if formattedDiff != "" {
			formattedDiff += " " // Add space between hours and minutes
		}
		formattedDiff += fmt.Sprintf("%dm", minutes)
	}
	return formattedDiff
}
