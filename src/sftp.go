package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type Regex struct {
	Directory string   `json:"directory"`
	Patterns  []string `json:"patterns"`
}

type SFTPFeed struct {
	FeedName    string   `json:"feed_name"`
	Host        string   `json:"host"`
	Port        int      `json:"port"`
	Filepaths   []string `json:"filepaths"`
	Directories []string `json:"directories"`
	Regexes     []Regex  `json:"regex"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
}

type Directory struct {
	Path  string
	Files *[]os.FileInfo
}

func ConnectToSFTP(f *SFTPFeed) (*sftp.Client, *ssh.Client, error) {
	// Create SSH client configuration
	config := &ssh.ClientConfig{
		User: f.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(f.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // **For testing only, DO NOT use in production**
	}

	// Connect to the SSH server
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", f.Host, f.Port), config)
	if err != nil {
		return nil, nil, err
	}

	// Create SFTP client
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, nil, err
	}
	return client, conn, nil
}

// Reads SFTP and returns a []os.FileInfo
func ReadSFTP(f *SFTPFeed) ([]Directory, error) {
	client, sshConn, err := ConnectToSFTP(f)
	if err != nil {
		return nil, err
	}
	defer sshConn.Close()
	defer client.Close()

	// Simple file access (better for memory management)
	// Uses "filepaths"
	directories := []Directory{}
	for _, filepath := range f.Filepaths {
		file, err := client.Stat(filepath)
		if err != nil {
			return nil, err
		}
		directories = append(directories, Directory{Path: filepath, Files: &[]os.FileInfo{file}})
	}

	// List all directorty (may be memory intensive so use sparingly)
	// Uses "paths"
	for _, path := range f.Directories {
		// Get a list of files and directories in the specified path
		files, err := client.ReadDir(path)
		if err != nil {
			return nil, err
		}
		directories = append(directories, Directory{Path: path, Files: &files})
	}

	// Regex lookup
	for _, rg := range f.Regexes {
		// Read directory to apply regex patterns
		files, err := client.ReadDir(rg.Directory)
		if err != nil {
			return nil, err
		}
		// For each regex pattern
		for _, pattern := range rg.Patterns {
			now := time.Now()
			// Replace special code ({yyyy}, {mm}, {dd})
			pattern = strings.Replace(pattern, "{yyyy}", fmt.Sprintf("%d", now.Year()), 1)
			pattern = strings.Replace(pattern, "{mm}", fmt.Sprintf("%02d", now.Month()), 1)
			pattern = strings.Replace(pattern, "{dd}", fmt.Sprintf("%d", now.Day()), 1)

			PrintStyled(fmt.Sprintf("🔎 Regex: %s", pattern), lipgloss.NewStyle().Foreground(LightBlue))

			// Build regex
			pattern := regexp.MustCompile(pattern)
			// Loop through files in directory
			matchedFiles := []os.FileInfo{}
			for _, file := range files {
				// Ignore dirs and check name matches
				if !file.IsDir() && pattern.MatchString(file.Name()) {
					matchedFiles = append(matchedFiles, file)
				}
			}
			directories = append(directories, Directory{Path: rg.Directory, Files: &matchedFiles})
		}

	}

	return directories, nil
}
