package main

import (
	"bufio"
	"os"
)

func main() {
	cm := ClientManager{}   // Create client manager object (this manages clients.json interaction)
	err := cm.ReadClients() // Load the data in clients.json
	if err != nil {
		PrintStyled(err.Error(), ErrorMessageStyle)
		os.Exit(-1) // Fatal error
	}
	cm.Start()
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Prevent window from closing
}
