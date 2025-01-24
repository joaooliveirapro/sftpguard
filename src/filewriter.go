package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func WriteToFile(table *string, feed *SFTPFeed, client *Client) {
	// Write data to file
	f, _ := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(fmt.Sprintf("💼 %s\n", client.Name))
	f.WriteString(fmt.Sprintf("🌐 %s\n", feed.Host))
	f.WriteString(fmt.Sprintf("⏰ Treshold %d hours\n", client.Treshold_hours))
	// Write regexes
	for _, rg := range feed.Regexes {
		for _, pattern := range rg.Patterns {
			f.WriteString(fmt.Sprintf("🔎 %s\n", pattern))
		}
	}
	// write table without ANSI sequences
	re := regexp.MustCompile(`\x1b\[(?:[0-9;]*m|\[[0-9;]*[a-zA-Z])`)
	plainText := re.ReplaceAllString(*table, "")
	f.WriteString(plainText)
	f.WriteString(fmt.Sprintf("\n%s\n", strings.Repeat("#", 150)))
}
