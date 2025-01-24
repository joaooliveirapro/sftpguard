package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/lipgloss"
)

type ClientManager struct {
	Clients []Client
}

type Client struct {
	Name           string     `json:"client_name"`
	Treshold_hours int        `json:"treshold_hours"`
	SFTPFeeds      []SFTPFeed `json:"feeds"`
}

func (cm *ClientManager) ReadClients() error {
	jsonFile, err := os.Open("clients.json")
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	content, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &cm.Clients)
	if err != nil {
		return err
	}
	return nil
}

func (cm *ClientManager) Start() {
	for _, client := range cm.Clients { // Loop through each client
		PrintStyled(client.Name, ClientNameStyle) // Print client header
		for _, feed := range client.SFTPFeeds {   // Loop through each feed
			PrintStyled(fmt.Sprintf("🌐 %s\n⏰ Treshold %d hours", feed.Host, client.Treshold_hours), FeedNameStyle) // Print host name
			// Create and print table
			table, err := CreateFeedTable(&feed, &client)
			if err != nil {
				PrintStyled(err.Error(), ErrorMessageStyle)
				continue // ignore as the error isn't too severe
			}
			fmt.Println(table)
			PrintStyled("\n", lipgloss.NewStyle()) // Some separation between next table
			// Write data to file
			WriteToFile(&table, &feed, &client)
		}
		PrintStyled("\n", lipgloss.NewStyle())
	}
}
