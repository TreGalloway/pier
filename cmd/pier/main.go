package main

import (
	"fmt"
	"os"

	"github.com/TreGalloway/pier/pkg/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// Create a "CSS class"
myStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))

// Apply it to text
styledText := myStyle.Render("Hello")

// Print it
fmt.Println(styledText)  // Prints "Hello" in red
