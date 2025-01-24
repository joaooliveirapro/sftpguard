package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss/table"
)

func CreateFeedTable(feed *SFTPFeed, client *Client) (string, error) {
	// Table data - one table per feed
	headers := []string{"Path", "File name", "Modified date"}
	rows := [][]string{}

	t := table.New()                // Create a new Table
	t.Border(TableBorderStyle)      // Style table border
	t.BorderStyle(TableBorderColor) // Style table color
	t.Headers(headers...)

	// Access feed SFTP and index its directories ("paths": [])
	directories, err := ReadSFTP(feed)
	if err != nil {
		return "", err
	}

	// Loop through each directory for the feed ("path")
	pathDisplayed := false
	for _, dir := range directories {

		for _, file := range *dir.Files {
			fileNameStyle := DefaultFileNameStyle

			modTime := file.ModTime()
			diff := time.Since(modTime.UTC())

			// Highlight if not updated since Treshold
			if diff.Hours() > float64(client.Treshold_hours) {
				fileNameStyle = HighlightedRowStyle
			}

			fileIcon := "📜"
			if file.IsDir() {
				fileIcon = "📁"
				fileNameStyle = FileIsFolderStyle
			}

			// Apply styles to FileName and LastUpdated
			fileName := fileNameStyle.Render(fmt.Sprintf("%s /%s", fileIcon, file.Name()))
			lastUpdated := LastUpdatedStyle.Render(fmt.Sprintf("%s (%s ago)", modTime.Format("02-01-2006 15:04:05"), FormattedDiff(diff)))

			if !pathDisplayed { // Only show dir.Path once, not duplicated each row
				rows = append(rows, []string{dir.Path, fileName, lastUpdated})
				pathDisplayed = true
			} else {
				rows = append(rows, []string{"", fileName, lastUpdated})
			}
		}

		// Reset pathDisplay
		pathDisplayed = false
	}

	t.Rows(rows...) // Add rows
	return t.String(), nil
}
