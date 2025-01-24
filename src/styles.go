package main

import "github.com/charmbracelet/lipgloss"

// Colors
const (
	LightGray  = lipgloss.Color("#f5efed")
	LightGray2 = lipgloss.Color("#b8b8b8")
	Green      = lipgloss.Color("#79a160")
	Red        = lipgloss.Color("#cf290c")
	Black      = lipgloss.Color("#000000")
	White      = lipgloss.Color("#ffffff")
	LightBlue  = lipgloss.Color("#98d0f5")
	Yellow     = lipgloss.Color("#f0d95b")
)

// Styles
var (
	DefaultFileNameStyle = lipgloss.NewStyle()
	LastUpdatedStyle     = lipgloss.NewStyle().Foreground(LightGray2)
	FileIsFolderStyle    = lipgloss.NewStyle().Foreground(Green)
	ErrorMessageStyle    = lipgloss.NewStyle().Padding(2).Background(LightGray).Foreground(Red)
	ClientNameStyle      = lipgloss.NewStyle().Background(lipgloss.Color(LightGray)).Padding(1, 2).Foreground(lipgloss.Color(Black))
	FeedNameStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color(Yellow))
	HighlightedRowStyle  = lipgloss.NewStyle().Background(Red)
	TableBorderStyle     = lipgloss.NormalBorder()
	TableBorderColor     = lipgloss.NewStyle().Foreground(lipgloss.Color(LightBlue))
	TableHeaderColor     = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(LightBlue))
)